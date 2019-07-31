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
	[[ $alias ]] [[ $path -]]
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

	t, err := template.New("MySQLImplement").
		Delims("[[", "]]").
		Funcs(template.FuncMap{
			"base": path.Base,
		}).
		Parse(implementTemplate)
	if err != nil {
		logger.Log.Errorf("init text/template err: %s", err.Error())
		os.Exit(-11)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, goImplGenerator); err != nil {
		logger.Log.Errorf("render text/template err: %s", err.Error())
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

	if err := ioutil.WriteFile(targetFile, buf.Bytes(), 0644); err != nil {
		logger.Log.Errorf("dump file `%s` err: %s", targetFile, err.Error())
		os.Exit(-13)
	}
}
