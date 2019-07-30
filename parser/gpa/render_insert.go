package gpaparser

import (
	"bytes"
	"go/ast"
	"os"
	"text/template"

	"github.com/liupangzi/gpa/logger"
)

type BindvarsInsert struct {
	StructName string
	Statement  string
	Method     string
	Param      string
	Arg        string
	StarExpr   bool
}

func RenderInsertTpl(goImplGenerator *GoImplGenerator, method *GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).NumField(); i++ {
		columnNames = append(columnNames, "` + "`" + `"+reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).Field(i).Tag.Get("db")+"` + "`" + `")
		values = append(values, ":"+reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).Field(i).Tag.Get("db"))
	}
	return db.DB.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", [[ .Arg ]].TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), [[ if not .StarExpr ]]&[[ end ]][[ .Arg ]])
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	columnNames, values := make([]string, 0), make([]string, 0)
	for i := 0; i < reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).NumField(); i++ {
		columnNames = append(columnNames, "` + "`" + `"+reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).Field(i).Tag.Get("db")+"` + "`" + `")
		values = append(values, ":"+reflect.TypeOf([[ if .StarExpr ]]*[[ .Arg ]][[ else ]][[ .Arg ]][[ end ]]).Field(i).Tag.Get("db"))
	}
	return tx.TX.NamedExec(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", [[ .Arg ]].TableName(), strings.Join(columnNames, ", "), strings.Join(values, ", ")), [[ if not .StarExpr ]]&[[ end ]][[ .Arg ]])
}
`

	if len(method.Params) != 2 {
		logger.Log.Errorf("Params of method `%s` err", method.Name)
		os.Exit(-1)
	}

	bindvarsInsert := &BindvarsInsert{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Arg:        method.Params[1].Names[0].Name,
	}

	if selector, ok := method.Params[1].Type.(*ast.SelectorExpr); ok {
		bindvarsInsert.StarExpr = false
		bindvarsInsert.Param = method.Params[1].Names[0].Name + " " + selector.X.(*ast.Ident).Name + "." + selector.Sel.Name
	}

	if selector, ok := method.Params[1].Type.(*ast.StarExpr); ok {
		bindvarsInsert.StarExpr = true
		bindvarsInsert.Param = method.Params[1].Names[0].Name + " *" + selector.X.(*ast.SelectorExpr).X.(*ast.Ident).Name + "." + selector.X.(*ast.SelectorExpr).Sel.Name
	}

	t, err := template.New("insertTpl").Delims("[[", "]]").Parse(tpl)
	if err != nil {
		logger.Log.Errorf("init text/template err: %s", err.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, bindvarsInsert); err != nil {
		logger.Log.Errorf("render text/template err: %s", err.Error())
		os.Exit(-2)
	}

	return buf.Bytes()
}
