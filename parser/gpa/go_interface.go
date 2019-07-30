package gpaparser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	crudListener "github.com/liupangzi/gpa/listener/crud"
	"github.com/liupangzi/gpa/logger"
	crudParser "github.com/liupangzi/gpa/parser/crud"
)

type CRUD int8

const (
	Unknown CRUD = iota
	Insert
	Select
	Update
	Delete
)

type GoComment struct {
	Content string
	Line    int
}

type GoMethod struct {
	Name    string
	Params  []*ast.Field
	Results []*ast.Field
	Line    int

	DML          CRUD
	DMLStatement string
}

type GoInterface struct {
	StructName string
	Comments   []*GoComment
	Methods    []*GoMethod
}

type GoImplGenerator struct {
	FileSet *token.FileSet

	Source  string
	Package string
	Imports map[string]string // path -> alias

	IFace     *GoInterface
	IFaceFrom int
	IFaceTo   int
}

func (g *GoImplGenerator) Parse() {
	f, err := parser.ParseFile(g.FileSet, g.Source, nil, parser.ParseComments)
	if err != nil {
		logger.Log.Errorf("Parse() err: %s", err.Error())
		os.Exit(-1)
	}

	g.validate(f)
	g.parsePackage(f)
	g.parseImports(f)
	g.parseInterface(f)
	g.parseComments(f)
	g.parseMethods(f)
	g.bindAnnotations()
}

func (g *GoImplGenerator) validate(f *ast.File) {
	// 0. one interface only
	interfaceCount := 0
	for _, declaration := range f.Decls {
		if genDeclaration, ok := declaration.(*ast.GenDecl); ok {
			if genDeclaration.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDeclaration.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						interfaceCount += 1
						g.IFaceFrom = g.FileSet.Position(interfaceType.Pos()).Line
						g.IFaceTo = g.FileSet.Position(interfaceType.End()).Line
					}
				}
			}
		}
	}

	if interfaceCount != 1 {
		logger.Log.Errorf("gpa supports one interface per source file, found: %d", interfaceCount)
		os.Exit(-2)
	}
}

func (g *GoImplGenerator) parsePackage(f *ast.File) {
	g.Package = f.Name.Name
}

func (g *GoImplGenerator) parseImports(f *ast.File) {
	g.Imports = make(map[string]string, 0)

	for _, importSpec := range f.Imports {
		if importSpec.Path == nil {
			continue
		}

		g.Imports[importSpec.Path.Value] = func() string {
			if importSpec.Name != nil {
				return importSpec.Name.Name
			} else {
				return ""
			}
		}()
	}
}

func (g *GoImplGenerator) parseInterface(f *ast.File) {
	g.IFace = &GoInterface{}
}

func (g *GoImplGenerator) parseComments(f *ast.File) {
	annotationRegex := regexp.MustCompile(`^@\w+\(`)

	comments := make([]*GoComment, 0)
	for _, commentGroup := range f.Comments {
		for _, comment := range commentGroup.List {
			commentText := comment.Text

			if strings.HasPrefix(commentText, "//") {
				commentText = comment.Text[2:] // remove `//`
			} else {
				rawCommentLength := len(commentText)
				commentText = commentText[2 : rawCommentLength-2] // remove `/* ... */`
			}

			// multi-line parsing support
			commentText = strings.TrimSpace(strings.ReplaceAll(commentText, "\n", " "))

			// catch annotation text
			if annotationRegex.FindStringIndex(commentText) != nil && annotationRegex.FindStringIndex(commentText)[0] == 0 {
				comments = append(comments, &GoComment{
					Content: commentText,
					Line:    g.FileSet.Position(comment.Pos()).Line,
				})
			}
		}
	}
	g.IFace.Comments = comments
}

func (g *GoImplGenerator) parseMethods(f *ast.File) {
	for _, declaration := range f.Decls {
		if genDeclaration, ok := declaration.(*ast.GenDecl); ok {
			if genDeclaration.Tok != token.TYPE {
				continue
			}

			g.IFace.Methods = make([]*GoMethod, 0)
			for _, spec := range genDeclaration.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						g.IFace.StructName = typeSpec.Name.Name
						for _, method := range interfaceType.Methods.List {
							if funcType, ok := method.Type.(*ast.FuncType); ok {
								if !method.Names[0].IsExported() {
									logger.Log.Errorf("Method `%s` is not exported... quit.", method.Names[0].String())
									os.Exit(-10)
								}

								method := &GoMethod{
									Line: g.FileSet.Position(method.Names[0].Pos()).Line,
									Name: method.Names[0].String(),
								}

								if funcType.Params != nil {
									method.Params = funcType.Params.List
								}

								if funcType.Results != nil {
									method.Results = funcType.Results.List
								}

								g.IFace.Methods = append(g.IFace.Methods, method)
							}
						}
					}
				}
			}
		}
	}
}

func (g *GoImplGenerator) bindAnnotations() {
	commentCursor, methodCursor := 0, 0
	for commentCursor < len(g.IFace.Comments) && methodCursor < len(g.IFace.Methods) {
		comment, method := g.IFace.Comments[commentCursor], g.IFace.Methods[methodCursor]
		if comment.Line < method.Line {
			if method.DML != Unknown {
				logger.Log.Errorf("Multiple annotations found for method `%s`", method.Name)
				os.Exit(-110)
			}

			p := crudParser.NewCRUDAnnotationParser(antlr.NewCommonTokenStream(crudParser.NewCRUDAnnotationLexer(antlr.NewInputStream(comment.Content)), antlr.TokenDefaultChannel))
			p.BuildParseTrees = true
			listener := crudListener.CRUDListener{
				Annotation: comment.Content,
			}
			antlr.ParseTreeWalkerDefault.Walk(&listener, p.Crud())

			if strings.HasPrefix(strings.ToLower(comment.Content), "@insert") {
				method.DML = Insert
				g.Imports[`"fmt"`] = ""
				g.Imports[`"reflect"`] = ""
				g.Imports[`"strings"`] = ""
			} else if strings.HasPrefix(strings.ToLower(comment.Content), "@update") {
				method.DML = Update
				method.DMLStatement = listener.UpdateStatement
			} else if strings.HasPrefix(strings.ToLower(comment.Content), "@delete") {
				method.DML = Delete
				method.DMLStatement = listener.DeleteStatement
			} else if strings.HasPrefix(strings.ToLower(comment.Content), "@select") {
				method.DML = Select
				method.DMLStatement = listener.SelectStatement
			}

			commentCursor += 1
		} else {
			methodCursor += 1
		}
	}
}
