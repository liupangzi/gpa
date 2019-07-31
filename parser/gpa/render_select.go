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

	var selector *ast.SelectorExpr
	if s, isSlice := method.Results[0].Type.(*ast.ArrayType); isSlice { // return slice
		bindvarsSelect.IsSlice = true
		if t, isPointer := s.Elt.(*ast.StarExpr); isPointer {
			bindvarsSelect.IsPointer = true
			selector = t.X.(*ast.SelectorExpr)
		} else if t, isSelector := s.Elt.(*ast.SelectorExpr); isSelector {
			selector = t
		} else {
			logger.Log.Error("Unsupported method.Result type found")
			os.Exit(-31)
		}
	} else {
		if t, isPointer := method.Results[0].Type.(*ast.StarExpr); isPointer {
			bindvarsSelect.IsPointer = true
			selector = t.X.(*ast.SelectorExpr)
		} else if t, isSelector := method.Results[0].Type.(*ast.SelectorExpr); isSelector {
			selector = t
		} else {
			logger.Log.Error("Unsupported method.Result type found")
			os.Exit(-31)
		}
	}

	bindvarsSelect.Result = selector.X.(*ast.Ident).Name + "." + selector.Sel.Name

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
