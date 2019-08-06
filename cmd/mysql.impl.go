package cmd

import (
	"bytes"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/liupangzi/gpa/logger"
	"github.com/liupangzi/gpa/parser/gpa"
	"github.com/spf13/cobra"
)

var implSourceFile string

func init() {
	implCmd.Flags().StringVarP(&implSourceFile, "source", "s", "", "source go interface file")
	if err := implCmd.MarkFlagRequired("source"); err != nil {
		logger.Log.Error(err.Error())
		os.Exit(-10)
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
	"github.com/jmoiron/sqlx"
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

func Implement(s string) {
	targetFile := s[:len(s)-2] + "impl.go"

	goImplGenerator := &gpaparser.GoImplGenerator{
		FileSet: token.NewFileSet(),
		Source:  s,
	}
	goImplGenerator.Parse()

	t, tplErr := template.New("MySQLImplement").
		Delims("[[", "]]").
		Funcs(template.FuncMap{
			"base": path.Base,
		}).
		Parse(implementTemplate)
	if tplErr != nil {
		logger.Log.Errorf("init text/template tplErr: %s", tplErr.Error())
		os.Exit(-11)
	}

	buf := &bytes.Buffer{}
	if tplExecuteErr := t.Execute(buf, goImplGenerator); tplExecuteErr != nil {
		logger.Log.Errorf("render text/template tplExecuteErr: %s", tplExecuteErr.Error())
		os.Exit(-12)
	}

	for _, method := range goImplGenerator.IFace.Methods {
		if method.DML == gpaparser.Update || method.DML == gpaparser.Delete {
			buf.Write(gpaparser.RenderUpdateOrDeleteTpl(goImplGenerator, method))
		} else if method.DML == gpaparser.Insert {
			buf.Write(gpaparser.RenderInsertTpl(goImplGenerator, method))
		} else if method.DML == gpaparser.Select {
			buf.Write(gpaparser.RenderSelectTpl(goImplGenerator, method))
		}
	}

	if writeFileErr := ioutil.WriteFile(targetFile, buf.Bytes(), 0644); writeFileErr != nil {
		logger.Log.Errorf("dump file `%s` writeFileErr: %s", targetFile, writeFileErr.Error())
		os.Exit(-13)
	}
}
