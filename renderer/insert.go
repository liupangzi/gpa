package renderer

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"strings"
	"text/template"

	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
)

type BindvarsInsert struct {
	StructName  string
	Statement   string
	Method      string
	Param       string
	Arg         string
	ColumnNames string
	BindValues  string

	DTOIsPointer bool // Param is a pointer to struct
	DTOPackage   string
	DTOName      string
}

func RenderBindvarsInsertTpl(goImplGenerator *gpaparser.Parser, method *gpaparser.GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	return db.DB.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s ([[ .ColumnNames ]]) VALUES ([[ .BindValues ]])", [[ .Arg ]].TableName()), [[ if not .DTOIsPointer ]]&[[ end ]][[ .Arg ]])
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	return tx.TX.NamedExecContext(ctx, fmt.Sprintf("INSERT INTO %s ([[ .ColumnNames ]]) VALUES ([[ .BindValues ]])", [[ .Arg ]].TableName()), [[ if not .DTOIsPointer ]]&[[ end ]][[ .Arg ]])
}
`

	bindvarsInsert := &BindvarsInsert{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Arg:        method.Params[1].Names[0].Name,
	}

	astType := method.Params[1].Type

	if star, isStar := astType.(*ast.StarExpr); isStar {
		bindvarsInsert.DTOIsPointer = true
		astType = star.X
	}

	if selector, isSelector := astType.(*ast.SelectorExpr); isSelector {
		if bindvarsInsert.DTOIsPointer {
			bindvarsInsert.Param = fmt.Sprintf("%s *%s.%s", method.Params[1].Names[0].Name, selector.X.(*ast.Ident).Name, selector.Sel.Name)
		} else {
			bindvarsInsert.Param = fmt.Sprintf("%s %s.%s", method.Params[1].Names[0].Name, selector.X.(*ast.Ident).Name, selector.Sel.Name)
		}
		bindvarsInsert.DTOPackage = selector.X.(*ast.Ident).Name
		bindvarsInsert.DTOName = selector.Sel.Name
	}

	if identity, isIdentity := astType.(*ast.Ident); isIdentity {
		logger.Log.Errorf("Struct `%s` and interface `%s` should not in the same package", identity.Name, goImplGenerator.IFace.StructName)
		os.Exit(-1)
	}

	columnNames, bindValues := make([]string, 0), make([]string, 0)
	loopFunc := goImplGenerator.ParseStructFromImports(bindvarsInsert.DTOPackage, bindvarsInsert.DTOName).IterFunc()
	for kv, ok := loopFunc(); ok; kv, ok = loopFunc() {
		columnNames = append(columnNames, fmt.Sprintf("`%s`", kv.Key))
		bindValues = append(bindValues, fmt.Sprintf(":%s", kv.Key))
	}

	bindvarsInsert.ColumnNames = strings.Join(columnNames, ", ")
	bindvarsInsert.BindValues = strings.Join(bindValues, ", ")

	t, tplErr := template.New("bindvarsInsert").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, bindvarsInsert); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	return buf.Bytes()
}
