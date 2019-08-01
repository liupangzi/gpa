package gpaparser

import (
	"bytes"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/liupangzi/gpa/logger"
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
}

func RenderSelectTpl(goImplGenerator *GoImplGenerator, method *GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
	var result [[ if .IsSlice ]][][[ end ]][[ if and .IsSlice .IsPointer ]]*[[ end ]][[ .Result ]]
	err := db.DB.[[ if .IsSlice ]]Select[[ else ]]Get[[ end ]](&result, [[ .Statement ]][[ .Args ]])
	return [[ if and .IsPointer (not .IsSlice) ]]&[[ end ]]result, err
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
	var result [[ if .IsSlice ]][][[ end ]][[ if and .IsSlice .IsPointer ]]*[[ end ]][[ .Result ]]
	err := tx.TX.[[ if .IsSlice ]]Select[[ else ]]Get[[ end ]](&result, [[ .Statement ]][[ .Args ]])
	return [[ if and .IsPointer (not .IsSlice) ]]&[[ end ]]result, err
}
`
	var args, params strings.Builder

	for idx, param := range method.Params {
		if idx == 0 {
			continue
		}

		typedNames := make([]string, 0)

		for _, name := range param.Names {
			args.WriteString(", ")
			args.WriteString(name.Name)
			typedNames = append(typedNames, name.Name)
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

	t, err := template.New("selectTpl").Delims("[[", "]]").Parse(tpl)
	if err != nil {
		logger.Log.Errorf("init text/template err: %s", err.Error())
		os.Exit(-31)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, bindvarsSelect); err != nil {
		logger.Log.Errorf("render text/template err: %s", err.Error())
		os.Exit(-33)
	}

	return buf.Bytes()
}
