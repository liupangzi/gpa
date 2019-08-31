package renderer

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	selectListener "github.com/liupangzi/gpa/listener/select"
	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
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
	IsInQuery bool
	Result    string
	Limit     string
}

func RenderBindvarsSelectTpl(goImplGenerator *gpaparser.Parser, method *gpaparser.GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if .Limit ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[- end ]]
	[[- if .IsInQuery ]]

	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return result, inErr
	}

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	[[- else ]]
	err := db.DB.SelectContext(ctx, &result, [[ .Statement ]][[ .Args ]])
	[[- end ]]
	return result, err
[[- else ]]
	var result [[ .Result ]]
	[[- if .IsInQuery ]]

	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return result, inErr
	}

	err := db.DB.GetContext(ctx, &result, db.DB.Rebind(query), args...)
	[[- else ]]
	err := db.DB.GetContext(ctx, &result, [[ .Statement ]][[ .Args ]])
	[[- end ]]
	return [[ if .IsPointer ]]&[[ end ]]result, err
[[- end ]]
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if .Limit ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[- end ]]
	[[- if .IsInQuery ]]

	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return result, inErr
	}

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	[[- else ]]
	err := tx.TX.SelectContext(ctx, &result, [[ .Statement ]][[ .Args ]])
	[[- end ]]
	return result, err
[[- else ]]
	var result [[ .Result ]]
	[[- if .IsInQuery ]]

	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return result, inErr
	}

	err := tx.TX.GetContext(ctx, &result, tx.TX.Rebind(query), args...)
	[[- else ]]
	err := tx.TX.GetContext(ctx, &result, [[ .Statement ]][[ .Args ]])
	[[- end ]]
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
			args.WriteString(fmt.Sprintf(", %s", name.Name))
			typedNames = append(typedNames, name.Name)
			allTypedNames = append(allTypedNames, name.Name)
		}

		params.WriteString(", ")
		// either []primitiveType or primitiveType/customType
		if tp, isArrayType := param.Type.(*ast.ArrayType); isArrayType {
			params.WriteString(fmt.Sprintf("%s []%s", strings.Join(typedNames, ", "), tp.Elt.(*ast.Ident).Name))
		} else {
			params.WriteString(fmt.Sprintf("%s %s", strings.Join(typedNames, ", "), param.Type.(*ast.Ident).Name))
		}
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

	// * type
	if star, isPointer := astType.(*ast.StarExpr); isPointer {
		bindvarsSelect.IsPointer = true
		astType = star.X
	}

	if selector, isSelector := astType.(*ast.SelectorExpr); isSelector { // a.B
		bindvarsSelect.Result = fmt.Sprintf("%s.%s", selector.X.(*ast.Ident).Name, selector.Sel.Name)
	} else if identity, isIdentity := astType.(*ast.Ident); isIdentity { // raw name
		bindvarsSelect.Result = identity.Name
	}

	p := selectParser.NewSelectStatementParserParser(antlr.NewCommonTokenStream(selectParser.NewSelectStatementParserLexer(antlr.NewInputStream(method.DMLStatement[1:len(method.DMLStatement)-1])), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true
	listener := selectListener.SelectListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	// statement contains `in (?)`
	bindvarsSelect.IsInQuery = listener.IsInQuery

	// init slice with known cap
	if listener.Limit != "" {
		if listener.Limit == "?" {
			bindvarsSelect.Limit = allTypedNames[len(allTypedNames)-1-listener.ReversedIndex]
		} else {
			bindvarsSelect.Limit = listener.Limit
		}
	}

	t, tplErr := template.New("bindvarsSelect").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, bindvarsSelect); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	return buf.Bytes()
}
