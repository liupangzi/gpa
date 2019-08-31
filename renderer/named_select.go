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

type NamedSelect struct {
	StructName     string
	Statement      string
	Method         string
	Param          string
	Arg            string
	ParamIsMap     bool
	ParamIsPointer bool

	IsPointer  bool
	IsSlice    bool
	IsInQuery  bool
	Result     string
	Limit      string
	DTOPackage string
	DTOName    string
}

func RenderNamedSelectTpl(goImplGenerator *gpaparser.Parser, method *gpaparser.GoMethod) []byte {
	tpl := `
func (db *[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if and .Limit (not .ParamIsMap) ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[- end ]]

	query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
	if namedErr != nil {
		return result, namedErr
	}

	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return result, inErr
	}
	[[- end ]]

	err := db.DB.SelectContext(ctx, &result, db.DB.Rebind(query), args...)
	return result, err
[[- else ]]
	var result [[ .Result ]]
	query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
	if namedErr != nil {
		return [[ if .IsPointer ]]&[[ end ]]result, namedErr
	}

	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return [[ if .IsPointer ]]&[[ end ]]result, inErr
	}
	[[- end ]]

	err := db.DB.GetContext(ctx, &result, db.DB.Rebind(query), args...)
	return [[ if .IsPointer ]]&[[ end ]]result, err
[[- end ]]
}

func (tx *Tx[[ .StructName ]]Impl) [[ .Method ]](ctx context.Context, [[ .Param ]]) ([[ if .IsSlice ]][][[ end ]][[ if .IsPointer ]]*[[ end ]][[ .Result ]], error) {
[[- if .IsSlice ]]
	[[- if and .Limit (not .ParamIsMap) ]]
	result := make([][[ if .IsPointer ]]*[[ end ]][[ .Result ]], 0, [[ .Limit ]])
	[[- else ]]
	var result [][[ if .IsPointer ]]*[[ end ]][[ .Result ]]
	[[- end ]]

	query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
	if namedErr != nil {
		return result, namedErr
	}

	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return result, inErr
	}
	[[- end ]]

	err := tx.TX.SelectContext(ctx, &result, tx.TX.Rebind(query), args...)
	return result, err
[[- else ]]
	var result [[ .Result ]]
	query, args, namedErr := sqlx.Named([[ .Statement ]], [[ .Arg ]])
	if namedErr != nil {
		return [[ if .IsPointer ]]&[[ end ]]result, namedErr
	}

	[[- if .IsInQuery ]]
	query, args, inErr := sqlx.In(query, args...)
	if inErr != nil {
		return [[ if .IsPointer ]]&[[ end ]]result, inErr
	}
	[[- end ]]

	err := tx.TX.GetContext(ctx, &result, tx.TX.Rebind(query), args...)
	return [[ if .IsPointer ]]&[[ end ]]result, err
[[- end ]]
}
`

	if len(method.Params) != 2 || len(method.Params[1].Names) != 1 {
		logger.Log.Errorf("Named select should contain two params only.")
		os.Exit(-1)
	}

	namedSelect := &NamedSelect{
		StructName: goImplGenerator.IFace.StructName,
		Statement:  method.DMLStatement,
		Method:     method.Name,
		Arg:        method.Params[1].Names[0].Name,
	}

	// either map or struct
	if mt, isMapType := method.Params[1].Type.(*ast.MapType); isMapType {
		if keyType, isString := mt.Key.(*ast.Ident); !isString || keyType.Name != "string" {
			logger.Log.Errorf("The key type of map in named select should be `string`")
			os.Exit(-1)
		}

		if _, isInterface := mt.Value.(*ast.InterfaceType); !isInterface {
			logger.Log.Errorf("The value type of map in named select should be `interface{}`")
			os.Exit(-1)
		}

		namedSelect.ParamIsMap = true
		namedSelect.Param = fmt.Sprintf("%s map[string]interface{}", namedSelect.Arg)
	} else {
		tp := method.Params[1].Type
		if star, isStar := method.Params[1].Type.(*ast.StarExpr); isStar {
			tp = star.X
			namedSelect.ParamIsPointer = true
		}

		if selector, isSelector := tp.(*ast.SelectorExpr); isSelector {
			namedSelect.DTOPackage = selector.X.(*ast.Ident).Name
			namedSelect.DTOName = selector.Sel.Name
			if namedSelect.ParamIsPointer {
				namedSelect.Param = fmt.Sprintf("%s *%s.%s", namedSelect.Arg, namedSelect.DTOPackage, namedSelect.DTOName)
			} else {
				namedSelect.Param = fmt.Sprintf("%s %s.%s", namedSelect.Arg, namedSelect.DTOPackage, namedSelect.DTOName)
			}
		} else {
			logger.Log.Errorf("DTO in named select should either be a map or a custom struct")
			os.Exit(-1)
		}
	}

	astType := method.Results[0].Type

	// return slice
	if _, isSlice := astType.(*ast.ArrayType); isSlice {
		namedSelect.IsSlice = true
		astType = astType.(*ast.ArrayType).Elt
	}

	// * type
	if star, isPointer := astType.(*ast.StarExpr); isPointer {
		namedSelect.IsPointer = true
		astType = star.X
	}

	if selector, isSelector := astType.(*ast.SelectorExpr); isSelector { // a.B
		namedSelect.Result = fmt.Sprintf("%s.%s", selector.X.(*ast.Ident).Name, selector.Sel.Name)
	} else if identity, isIdentity := astType.(*ast.Ident); isIdentity { // raw name
		namedSelect.Result = identity.Name
	}

	p := selectParser.NewSelectStatementParserParser(antlr.NewCommonTokenStream(selectParser.NewSelectStatementParserLexer(antlr.NewInputStream(method.DMLStatement[1:len(method.DMLStatement)-1])), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true
	listener := selectListener.SelectListener{}
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Statement())

	// statement contains `in (:namedField)`
	namedSelect.IsInQuery = listener.IsInQuery

	// init slice with known cap
	if listener.Limit != "" {
		if strings.HasPrefix(listener.Limit, ":") {
			k := strings.TrimSpace(listener.Limit[1:])
			if namedSelect.ParamIsMap {
				namedSelect.Limit = fmt.Sprintf(`%s["%s"].(int64)`, namedSelect.Arg, k)
			} else {
				if field, tagExists := goImplGenerator.ParseStructFromImports(namedSelect.DTOPackage, namedSelect.DTOName).Get(k); tagExists {
					namedSelect.Limit = fmt.Sprintf(`%s.%s`, namedSelect.Arg, field)
				} else {
					logger.Log.Errorf("Named param field `%s` not found in struct tags", k)
					os.Exit(-1)
				}
			}
		} else {
			namedSelect.Limit = listener.Limit
		}
	}

	t, tplErr := template.New("namedSelect").Delims("[[", "]]").Parse(tpl)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, namedSelect); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	return buf.Bytes()
}
