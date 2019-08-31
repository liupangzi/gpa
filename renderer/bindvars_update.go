package renderer

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	updateListener "github.com/liupangzi/gpa/listener/update"
	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
	updateParser "github.com/liupangzi/gpa/parser/update"
)

type BindvarsUpdate struct {
	StructName string
	Statement  string
	Method     string
	Args       string
	Params     string
	IsInQuery  bool
}

func RenderBindvarsUpdateTpl(goImplGenerator *gpaparser.Parser, method *gpaparser.GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) (sql.Result, error) {
	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return nil, inErr
	}

	return db.DB.ExecContext(ctx, db.DB.Rebind(query), args...)
	[[- else ]]
	return db.DB.ExecContext(ctx, [[ .Statement ]][[ .Args ]])
	[[- end ]]
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) (sql.Result, error) {
	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In([[ .Statement ]][[ .Args ]])
	if inErr != nil {
		return nil, inErr
	}

	return tx.TX.ExecContext(ctx, tx.TX.Rebind(query), args...)
	[[- else ]]
	return tx.TX.ExecContext(ctx, [[ .Statement ]][[ .Args ]])
	[[- end ]]
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

		if identity, isIdentity := param.Type.(*ast.Ident); isIdentity {
			params.WriteString(fmt.Sprintf(", %s %s", strings.Join(typedNames, ", "), identity.Name))
		} else if arr, isArrType := param.Type.(*ast.ArrayType); isArrType {
			params.WriteString(fmt.Sprintf(", %s []%s", strings.Join(typedNames, ", "), arr.Elt.(*ast.Ident).Name))
		} else {
			logger.Log.Errorf("param.Type should be a primitive type or an array consists of primitive type")
			os.Exit(-1)
		}
	}

	bindvarsUpdate := &BindvarsUpdate{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Args:       args.String(),
		Params:     params.String(),
	}

	p := updateParser.NewUpdateStatementParserParser(antlr.NewCommonTokenStream(updateParser.NewUpdateStatementParserLexer(antlr.NewInputStream(method.DMLStatement[1:len(method.DMLStatement)-1])), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true
	listener := updateListener.UpdateListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	bindvarsUpdate.IsInQuery = listener.IsInQuery

	t, tplErr := template.New("bindvarsUpdate").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, bindvarsUpdate); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	return buf.Bytes()
}
