package cmd

import (
	"bytes"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
	"github.com/liupangzi/gpa/renderer"
	"github.com/spf13/cobra"
)

var implSourceFile string

func init() {
	implCmd.Flags().StringVarP(&implSourceFile, "source", "s", "", "source go interface file")
	if err := implCmd.MarkFlagRequired("source"); err != nil {
		logger.Log.Error(err.Error())
		os.Exit(-1)
	}
}

var implCmd = &cobra.Command{
	Use:   "impl",
	Short: "Generate impl from go interface",
	Long:  "Generate impl from go interface",
	Run: func(cmd *cobra.Command, args []string) {
		Implement(implSourceFile)
	},
}

const implementTemplate = `// Code generated from [[ base .Source ]] by gpa. DO NOT EDIT.

package [[ .Package ]]

import (
[[- range $path, $alias := .Imports ]]
	[[ if $alias ]][[ $alias ]] [[ $path -]][[ else ]][[ $path -]][[ end -]]
[[ end ]]
)

type [[ .IFace.StructName ]]Impl struct {
	DB *sqlx.DB
}

type Tx[[ .IFace.StructName ]]Impl struct {
	TX *sqlx.Tx
}

func New[[ .IFace.StructName ]](db *sqlx.DB) *[[ .IFace.StructName ]]Impl {
	return &[[ .IFace.StructName ]]Impl{
		DB: db,
	}
}

func (db *[[ .IFace.StructName ]]Impl) Begin(ctx context.Context, options *sql.TxOptions) *Tx[[ .IFace.StructName ]]Impl {
	return &Tx[[ .IFace.StructName ]]Impl{
		TX: db.DB.MustBeginTx(ctx, options),
	}
}

func (tx *Tx[[ .IFace.StructName ]]Impl) Commit(ctx context.Context) error {
	return tx.TX.Commit()
}

func (tx *Tx[[ .IFace.StructName ]]Impl) Rollback(ctx context.Context) error {
	return tx.TX.Rollback()
}
`

func Implement(source string) {
	targetFile := source[:len(source)-2] + "impl.go"
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, source, nil, parser.ParseComments)
	if err != nil {
		logger.Log.Errorf("ParseFile(`%s`) err: %s", source, err.Error())
		os.Exit(-1)
	}

	goParser := &gpaparser.Parser{
		FileSet: fileSet,
		File:    file,
		Source:  source,
	}
	goParser.Parse()

	t, tplErr := template.New("MySQLImplement").
		Delims("[[", "]]").
		Funcs(template.FuncMap{
			"base": path.Base,
		}).
		Parse(implementTemplate)
	if tplErr != nil {
		logger.Log.Errorf("init template err: %s", tplErr.Error())
		os.Exit(-1)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, goParser); tplExecuteErr != nil {
		logger.Log.Errorf("render template err: %s", tplExecuteErr.Error())
		os.Exit(-1)
	}

	// sort by the original order defined in the interface
	sort.Slice(goParser.IFace.Methods, func(i, j int) bool {
		return goParser.IFace.Methods[i].Position < goParser.IFace.Methods[j].Position
	})

	for _, method := range goParser.IFace.Methods {
		if method.DML == gpaparser.Update {
			if strings.Contains(method.DMLStatement, ":") {
				buf.Write(renderer.RenderNamedUpdateTpl(goParser, method))
			} else {
				buf.Write(renderer.RenderBindvarsUpdateTpl(goParser, method))
			}
		} else if method.DML == gpaparser.Delete {
			if strings.Contains(method.DMLStatement, ":") {
				buf.Write(renderer.RenderNamedDeleteTpl(goParser, method))
			} else {
				buf.Write(renderer.RenderBindvarsDeleteTpl(goParser, method))
			}
		} else if method.DML == gpaparser.Insert {
			buf.Write(renderer.RenderBindvarsInsertTpl(goParser, method))
		} else if method.DML == gpaparser.Select {
			if strings.Contains(method.DMLStatement, ":") {
				buf.Write(renderer.RenderNamedSelectTpl(goParser, method))
			} else {
				buf.Write(renderer.RenderBindvarsSelectTpl(goParser, method))
			}
		}
	}

	if writeFileErr := ioutil.WriteFile(targetFile, buf.Bytes(), 0644); writeFileErr != nil {
		logger.Log.Errorf("dump file `%s` writeFileErr: %s", targetFile, writeFileErr.Error())
		os.Exit(-1)
	}
}
