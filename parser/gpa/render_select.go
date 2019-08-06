package gpaparser

import (
	"bytes"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	selectListener "github.com/liupangzi/gpa/listener/select"
	"github.com/liupangzi/gpa/logger"
	selectParser "github.com/liupangzi/gpa/parser/select"
)

type BindvarsSelect struct {
	StructName string
	Statement  string
	Method     string
	Params     string
	Args       string

	IsPointer bool
	IsSlice   bool
	Result    string
	Limit     string
}

func RenderSelectTpl(goImplGenerator *GoImplGenerator, method *GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if .Limit ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[ end ]]
	err := db.DB.Select(&result, [[ .Statement ]][[ .Args ]])
	return result, err
[[- else ]]
	var result [[ .Result ]]
	err := db.DB.Get(&result, [[ .Statement ]][[ .Args ]])
	return [[ if .IsPointer ]]&[[ end ]]result, err
[[- end ]]
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if .Limit ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[ end ]]
	err := tx.TX.Select(&result, [[ .Statement ]][[ .Args ]])
	return result, err
[[- else ]]
	var result [[ .Result ]]
	err := tx.TX.Get(&result, [[ .Statement ]][[ .Args ]])
	return [[ if .IsPointer ]]&[[ end ]]result, err
[[- end ]]
}
`
	args, params := new(strings.Builder), new(strings.Builder)
	allTypedNames := make([]string, 0)

	for idx, param := range method.Params {
		if idx == 0 {
			continue
		}

		typedNames := make([]string, 0)
		for _, name := range param.Names {
			args.WriteString(", ")
			args.WriteString(name.Name)
			typedNames = append(typedNames, name.Name)
			allTypedNames = append(allTypedNames, name.Name)
		}

		params.WriteString(", ")
		params.WriteString(strings.Join(typedNames, ", ") + " " + param.Type.(*ast.Ident).Name)
	}

	bindvarsSelect := &BindvarsSelect{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Args:       args.String(),
		Params:     params.String(),
	}

	astType := method.Results[0].Type

	// return slice
	if _, isSlice := astType.(*ast.ArrayType); isSlice {
		bindvarsSelect.IsSlice = true
		astType = astType.(*ast.ArrayType).Elt
	}

	if p, isPointer := astType.(*ast.StarExpr); isPointer {
		bindvarsSelect.IsPointer = true
		bindvarsSelect.Result = p.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + p.X.(*ast.SelectorExpr).Sel.Name
	} else if s, isSelector := astType.(*ast.SelectorExpr); isSelector {
		bindvarsSelect.Result = s.X.(*ast.Ident).Name + "." + s.Sel.Name
	} else {
		bindvarsSelect.Result = astType.(*ast.Ident).Name
	}

	p := selectParser.NewSelectStatementParserParser(antlr.NewCommonTokenStream(selectParser.NewSelectStatementParserLexer(antlr.NewInputStream(method.DMLStatement[1:len(method.DMLStatement)-1])), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true
	listener := selectListener.SelectListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	// init slice with known cap
	if listener.Limit != "" {
		if listener.Limit == "?" {
			bindvarsSelect.Limit = allTypedNames[len(allTypedNames)-1-listener.ReversedIndex]
		} else {
			bindvarsSelect.Limit = listener.Limit
		}
	}

	t, tplErr := template.New("selectTpl").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init text/template tplErr: %s", tplErr.Error())
		os.Exit(-31)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, bindvarsSelect); tplExecuteErr != nil {
		logger.Log.Errorf("render text/template tplExecuteErr: %s", tplExecuteErr.Error())
		os.Exit(-33)
	}

	return buf.Bytes()
}
