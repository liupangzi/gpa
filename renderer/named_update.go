package renderer

import (
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"text/template"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	updateListener "github.com/liupangzi/gpa/listener/update"
	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
	updateParser "github.com/liupangzi/gpa/parser/update"
)

type NamedUpdate struct {
	StructName     string
	Statement      string
	Method         string
	Arg            string
	Param          string
	ParamIsMap     bool
	ParamIsPointer bool
	IsInQuery      bool
	DTOPackage     string
	DTOName        string
}

func RenderNamedUpdateTpl(goImplGenerator *gpaparser.Parser, method *gpaparser.GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	[[- if .IsInQuery ]]
    query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
    if namedErr != nil {
		return nil, namedErr
    }

    query, args, inErr := sqlx.In(query, args...)
    if inErr != nil {
		return nil, inErr
    }

	return db.DB.ExecContext(ctx, db.DB.Rebind(query), args...)
	[[- else ]]
	return db.DB.NamedExecContext(ctx, [[ .Statement ]], [[ .Arg ]])
	[[- end ]]
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) (sql.Result, error) {
	[[- if .IsInQuery ]]
    query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
    if namedErr != nil {
		return nil, namedErr
    }

    query, args, inErr := sqlx.In(query, args...)
    if inErr != nil {
		return nil, inErr
    }

	return tx.TX.ExecContext(ctx, tx.TX.Rebind(query), args...)
	[[- else ]]
	return tx.TX.NamedExecContext(ctx, [[ .Statement ]], [[ .Arg ]])
	[[- end ]]
}
`

	if len(method.Params) != 2 || len(method.Params[1].Names) != 1 {
		logger.Log.Errorf("Named update/delete should contain two params only.")
		os.Exit(-1)
	}

	namedUpdate := &NamedUpdate{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Arg:        method.Params[1].Names[0].Name,
	}

	// either map or struct
	if mt, isMapType := method.Params[1].Type.(*ast.MapType); isMapType {
		if keyType, isString := mt.Key.(*ast.Ident); !isString || keyType.Name != "string" {
			logger.Log.Errorf("The key type of map in named update/delete should be `string`")
			os.Exit(-1)
		}

		if _, isInterface := mt.Value.(*ast.InterfaceType); !isInterface {
			logger.Log.Errorf("The value type of map in named update/delete should be `interface{}`")
			os.Exit(-1)
		}

		namedUpdate.ParamIsMap = true
		namedUpdate.Param = fmt.Sprintf("%s map[string]interface{}", namedUpdate.Arg)
	} else {
		tp := method.Params[1].Type
		if star, isStar := method.Params[1].Type.(*ast.StarExpr); isStar {
			tp = star.X
			namedUpdate.ParamIsPointer = true
		}

		if selector, isSelector := tp.(*ast.SelectorExpr); isSelector {
			namedUpdate.DTOPackage = selector.X.(*ast.Ident).Name
			namedUpdate.DTOName = selector.Sel.Name
			if namedUpdate.ParamIsPointer {
				namedUpdate.Param = fmt.Sprintf("%s *%s.%s", namedUpdate.Arg, namedUpdate.DTOPackage, namedUpdate.DTOName)
			} else {
				namedUpdate.Param = fmt.Sprintf("%s %s.%s", namedUpdate.Arg, namedUpdate.DTOPackage, namedUpdate.DTOName)
			}
		} else {
			logger.Log.Errorf("DTO in named update/delete should either be a map or a custom struct")
			os.Exit(-1)
		}
	}

	p := updateParser.NewUpdateStatementParserParser(antlr.NewCommonTokenStream(updateParser.NewUpdateStatementParserLexer(antlr.NewInputStream(method.DMLStatement[1:len(method.DMLStatement)-1])), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true
	listener := updateListener.UpdateListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	namedUpdate.IsInQuery = listener.IsInQuery

	t, tplErr := template.New("namedUpdate").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, namedUpdate); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	return buf.Bytes()
}
