package gpaparser

import (
	"bytes"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/liupangzi/gpa/logger"
)

type BindvarsUpdateOrDelete struct {
	StructName string
	Statement  string
	Method     string
	Args       string
	Params     string
}

func RenderUpdateOrDeleteTpl(goImplGenerator *GoImplGenerator, method *GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) (sql.Result, error) {
	return db.DB.ExecContext(ctx, [[ .Statement ]][[ .Args ]])
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context[[ .Params ]]) (sql.Result, error) {
	return tx.TX.ExecContext(ctx, [[ .Statement ]][[ .Args ]])
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

	bindvarsUpdate := &BindvarsUpdateOrDelete{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Args:       args.String(),
		Params:     params.String(),
	}

	t, tplErr := template.New("updateOrDeleteTpl").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init text/template tplErr: %s", tplErr.Error())
		os.Exit(-30)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, bindvarsUpdate); tplExecuteErr != nil {
		logger.Log.Errorf("render text/template tplExecuteErr: %s", tplExecuteErr.Error())
		os.Exit(-32)
	}

	return buf.Bytes()
}
