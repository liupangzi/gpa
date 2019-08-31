package gpaparser

import (
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
	"os"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	om "github.com/cevaris/ordered_map"
	crudListener "github.com/liupangzi/gpa/listener/crud"
	"github.com/liupangzi/gpa/logger"
	crudParser "github.com/liupangzi/gpa/parser/crud"
)

type CRUD int8

const (
	_ CRUD = iota
	Insert
	Select
	Update
	Delete
)

type GoMethod struct {
	Name     string
	Params   []*ast.Field
	Results  []*ast.Field
	Position token.Pos

	DML          CRUD
	DMLStatement string
}

type GoInterface struct {
	StructName string
	CommentMap ast.CommentMap
	Methods    []*GoMethod
}

type Parser struct {
	FileSet *token.FileSet
	File    *ast.File

	Source  string
	Package string
	Imports map[string]string // path -> alias

	IFace     *GoInterface
	IFaceFrom token.Pos
	IFaceTo   token.Pos
}

func (g *Parser) Parse() {
	g.parsePackage()
	g.parseImports()
	g.parseInterface()
	g.parseMethods()
}

func (g *Parser) parsePackage() {
	g.Package = g.File.Name.Name
}

func (g *Parser) parseImports() {
	// default sqlx package
	g.Imports = map[string]string{
		`"github.com/jmoiron/sqlx"`: "",
	}

	for _, importSpec := range g.File.Imports {
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

func (g *Parser) parseInterface() {
	g.IFace = &GoInterface{}
	ifNames := make([]string, 0)

	for _, declaration := range g.File.Decls {
		if _, isGenericDeclaration := declaration.(*ast.GenDecl); !isGenericDeclaration {
			logger.Log.Errorf("source file `%s` should contain interface type only", g.Source)
			os.Exit(-1)
		}

		if genericDeclaration, ok := declaration.(*ast.GenDecl); ok {
			if genericDeclaration.Tok == token.IMPORT {
				continue
			}

			if genericDeclaration.Tok != token.TYPE {
				logger.Log.Errorf("source file `%s` should contain interface type only", g.Source)
				os.Exit(-1)
			}

			for _, spec := range genericDeclaration.Specs {
				if typeSpec, isType := spec.(*ast.TypeSpec); isType {
					if interfaceType, isInterfaceType := typeSpec.Type.(*ast.InterfaceType); isInterfaceType {
						g.IFaceFrom = interfaceType.Pos()
						g.IFaceTo = interfaceType.End()
						g.IFace.StructName = typeSpec.Name.Name
						ifNames = append(ifNames, typeSpec.Name.Name)
					} else {
						logger.Log.Errorf("source file `%s` should contain interface type only", g.Source)
						os.Exit(-1)
					}
				} else {
					logger.Log.Errorf("source file `%s` should contain interface type only", g.Source)
					os.Exit(-1)
				}
			}
		}
	}

	if len(ifNames) != 1 {
		logger.Log.Errorf("gpa supports one interface per source file, found %d: %+v", len(ifNames), ifNames)
		os.Exit(-1)
	}
}

func (g *Parser) parseMethods() {
	for node := range ast.NewCommentMap(g.FileSet, g.File, g.File.Comments) {
		field, isField := node.(*ast.Field)
		if !isField || field.Doc == nil || len(field.Doc.List) == 0 || field.Doc.List[0].Slash < g.IFaceFrom || field.Doc.List[0].Slash > g.IFaceTo {
			// not a field or empty doc or current field beyond interface, skip
			continue
		}

		function, isFunction := field.Type.(*ast.FuncType)
		if len(field.Names) != 1 || !isFunction {
			// invalid function / not a function, skip
			continue
		}

		if !field.Names[0].IsExported() {
			logger.Log.Errorf("Method `%s` cannot be exported... quit.", field.Names[0].String())
			os.Exit(-1)
		}

		if function.Params == nil || function.Results == nil {
			logger.Log.Errorf("Nil param or result found in Method `%s`'s signature... quit.", field.Names[0].String())
			os.Exit(-1)
		}

		var CRUDAnnotation string

		// parse and format annotations
		for _, comment := range field.Doc.List {
			characters := make([]rune, 0, len(comment.Text))
			for _, c := range func(s string) string {
				if strings.HasPrefix(comment.Text, "//") {
					return strings.TrimSpace(comment.Text[2:])
				} else {
					return strings.TrimSpace(comment.Text[2 : len(comment.Text)-2])
				}
			}(comment.Text) {
				if c == '\t' || c == '\n' || c == '\r' || c == ' ' {
					if len(characters) != 0 && characters[len(characters)-1] != ' ' {
						characters = append(characters, ' ')
					}
				} else {
					characters = append(characters, c)
				}
			}

			annotation := string(characters)
			if strings.HasPrefix(strings.ToLower(annotation), "@insert") ||
				strings.HasPrefix(strings.ToLower(annotation), "@update") ||
				strings.HasPrefix(strings.ToLower(annotation), "@delete") ||
				strings.HasPrefix(strings.ToLower(annotation), "@select") {
				CRUDAnnotation = annotation
			}
		}

		if len(CRUDAnnotation) == 0 {
			logger.Log.Errorf("CRUD annotation not found in Method `%s`'s signature... quit.", field.Names[0].String())
			os.Exit(-1)
		}

		method := &GoMethod{
			Name:     field.Names[0].Name,
			Params:   function.Params.List,
			Results:  function.Results.List,
			Position: function.Pos(),
		}

		p := crudParser.NewCRUDAnnotationParser(antlr.NewCommonTokenStream(crudParser.NewCRUDAnnotationLexer(antlr.NewInputStream(CRUDAnnotation)), antlr.TokenDefaultChannel))
		p.BuildParseTrees = true
		listener := crudListener.CRUDListener{
			Annotation: CRUDAnnotation,
		}
		antlr.ParseTreeWalkerDefault.Walk(&listener, p.Crud())

		if strings.HasPrefix(strings.ToLower(CRUDAnnotation), "@insert") {
			method.DML = Insert
			g.Imports[`"fmt"`] = ""

			// @Insert method restrictions
			if len(method.Params) != 2 || len(method.Params[0].Names) != 1 || len(method.Params[1].Names) != 1 {
				logger.Log.Errorf("Method `%s` should contain two params only", method.Name)
				os.Exit(-1)
			}

			if method.Params[0].Names[0].Name != "ctx" {
				logger.Log.Errorf("Method `%s`'s first param's name should be `ctx`", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Params[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "context" || selector.Sel.Name != "Context" {
				logger.Log.Errorf("Method `%s`'s first param's type should be `context.Context`", method.Name)
				os.Exit(-1)
			}

			if len(method.Results) != 2 || len(method.Results[0].Names) != 0 || len(method.Results[1].Names) != 0 {
				logger.Log.Errorf("Method `%s` should contain two unnamed return results", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Results[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "sql" || selector.Sel.Name != "Result" {
				logger.Log.Errorf("Method `%s`'s first result's type should be `sql.Result`", method.Name)
				os.Exit(-1)
			}

			if identity, isIdentity := method.Results[1].Type.(*ast.Ident); !isIdentity || identity.Name != "error" {
				logger.Log.Errorf("Method `%s`'s second result's type should be `error`", method.Name)
				os.Exit(-1)
			}
		} else if strings.HasPrefix(strings.ToLower(CRUDAnnotation), "@update") {
			method.DML = Update
			method.DMLStatement = listener.UpdateStatement

			// @Update method restrictions
			if len(method.Params) < 2 || len(method.Params[0].Names) != 1 {
				logger.Log.Errorf("Method `%s` should contain at least two params", method.Name)
				os.Exit(-1)
			}

			if method.Params[0].Names[0].Name != "ctx" {
				logger.Log.Errorf("Method `%s`'s first param's name should be `ctx`", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Params[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "context" || selector.Sel.Name != "Context" {
				logger.Log.Errorf("Method `%s`'s first param's type should be `context.Context`", method.Name)
				os.Exit(-1)
			}

			if len(method.Results) != 2 || len(method.Results[0].Names) != 0 || len(method.Results[1].Names) != 0 {
				logger.Log.Errorf("Method `%s` should contain two unnamed return results", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Results[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "sql" || selector.Sel.Name != "Result" {
				logger.Log.Errorf("Method `%s`'s first result's type should be `sql.Result`", method.Name)
				os.Exit(-1)
			}

			if identity, isIdentity := method.Results[1].Type.(*ast.Ident); !isIdentity || identity.Name != "error" {
				logger.Log.Errorf("Method `%s`'s second result's type should be `error`", method.Name)
				os.Exit(-1)
			}
		} else if strings.HasPrefix(strings.ToLower(CRUDAnnotation), "@delete") {
			method.DML = Delete
			method.DMLStatement = listener.DeleteStatement

			// @Delete method restrictions
			if len(method.Params) < 2 || len(method.Params[0].Names) != 1 {
				logger.Log.Errorf("Method `%s` should contain at least two params", method.Name)
				os.Exit(-1)
			}

			if method.Params[0].Names[0].Name != "ctx" {
				logger.Log.Errorf("Method `%s`'s first param's name should be `ctx`", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Params[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "context" || selector.Sel.Name != "Context" {
				logger.Log.Errorf("Method `%s`'s first param's type should be `context.Context`", method.Name)
				os.Exit(-1)
			}

			if len(method.Results) != 2 || len(method.Results[0].Names) != 0 || len(method.Results[1].Names) != 0 {
				logger.Log.Errorf("Method `%s` should contain two unnamed return results", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Results[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "sql" || selector.Sel.Name != "Result" {
				logger.Log.Errorf("Method `%s`'s first result's type should be `sql.Result`", method.Name)
				os.Exit(-1)
			}

			if identity, isIdentity := method.Results[1].Type.(*ast.Ident); !isIdentity || identity.Name != "error" {
				logger.Log.Errorf("Method `%s`'s second result's type should be `error`", method.Name)
				os.Exit(-1)
			}
		} else if strings.HasPrefix(strings.ToLower(CRUDAnnotation), "@select") {
			method.DML = Select
			method.DMLStatement = listener.SelectStatement

			// @Select method restrictions
			if len(method.Params) == 0 || len(method.Params[0].Names) != 1 {
				logger.Log.Errorf("Method `%s` should contain at least one param", method.Name)
				os.Exit(-1)
			}

			if method.Params[0].Names[0].Name != "ctx" {
				logger.Log.Errorf("Method `%s`'s first param's name should be `ctx`", method.Name)
				os.Exit(-1)
			}

			if selector, isSelector := method.Params[0].Type.(*ast.SelectorExpr); !isSelector || selector.X.(*ast.Ident).Name != "context" || selector.Sel.Name != "Context" {
				logger.Log.Errorf("Method `%s`'s first param's type should be `context.Context`", method.Name)
				os.Exit(-1)
			}

			if len(method.Results) != 2 || len(method.Results[0].Names) != 0 || len(method.Results[1].Names) != 0 {
				logger.Log.Errorf("Method `%s` should contain two unnamed return results", method.Name)
				os.Exit(-1)
			}

			if identity, isIdentity := method.Results[1].Type.(*ast.Ident); !isIdentity || identity.Name != "error" {
				logger.Log.Errorf("Method `%s`'s second result's type should be `error`", method.Name)
				os.Exit(-1)
			}
		} else {
			logger.Log.Errorf("Nil CRUDAnnotation found in Method `%s`'s doc... quit.", field.Names[0].String())
			os.Exit(-1)
		}

		g.IFace.Methods = append(g.IFace.Methods, method)
	}
}

// ParseStructFromImports return tag => Field
func (g *Parser) ParseStructFromImports(packageName, structName string) *om.OrderedMap {
	// get tag names from ast node
	var pkg *types.Package
	var importerErr error
	for path, alias := range g.Imports {
		// remove ""
		rawPath := path[1 : len(path)-1]

		if alias == packageName {
			// DTO uses an alias as prefix
			if pkg, importerErr = importer.ForCompiler(token.NewFileSet(), "source", nil).Import(rawPath); importerErr != nil {
				logger.Log.Errorf("Parse imports err: ", importerErr.Error())
				os.Exit(-1)
			}
			break
		} else {
			// DTO uses the default package name
			if pkg, importerErr = importer.ForCompiler(token.NewFileSet(), "source", nil).Import(rawPath); importerErr != nil {
				logger.Log.Errorf("Parse imports err: ", importerErr.Error())
				os.Exit(-1)
			}

			if pkg.Name() == packageName {
				break
			}
		}
	}

	result := om.NewOrderedMap()
	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)
		if tn, ok := obj.(*types.TypeName); ok {
			if tn.Name() != structName {
				continue
			}

			// type path struct{}
			structLiteral := strings.Join(strings.Split(tn.String(), " ")[2:], " ")
			fields := strings.TrimRight(strings.TrimLeft(structLiteral, "struct{"), "}")
			for _, field := range strings.Split(fields, ";") {
				// fieldName type tagLiteral
				fieldName := strings.Split(strings.TrimSpace(field), " ")[0]
				tagLiteral, _ := strconv.Unquote(strings.Join(strings.Split(strings.TrimSpace(field), " ")[2:], " "))
				for _, t := range strings.Split(tagLiteral, " ") {
					if strings.HasPrefix(t, "db:") {
						result.Set(t[4:len(t)-1], fieldName)
					}
				}
			}
		}
	}
	return result
}
