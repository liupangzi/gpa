package cmd

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
	"unsafe"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	listener "github.com/liupangzi/gpa/listener/ddl"
	"github.com/liupangzi/gpa/logger"
	parser "github.com/liupangzi/gpa/parser/ddl"
	"github.com/spf13/cobra"
)

var modelDDLFile string
var modelTargetFile string
var modelPackageName string

func init() {
	modelCmd.Flags().StringVarP(&modelDDLFile, "source", "s", "", `mysql ddl file, e.g:
mysql > SHOW CREATE TABLE t`)
	if err := modelCmd.MarkFlagRequired("source"); err != nil {
		logger.Log.Error(err.Error())
		os.Exit(-1)
	}

	modelCmd.Flags().StringVarP(&modelTargetFile, "target", "t", "", "generated go model file")
	if err := modelCmd.MarkFlagRequired("target"); err != nil {
		logger.Log.Error(err.Error())
		os.Exit(-1)
	}

	modelCmd.Flags().StringVarP(&modelPackageName, "package_name", "p", "", "go model package name")
	if err := modelCmd.MarkFlagRequired("package_name"); err != nil {
		logger.Log.Error(err.Error())
		os.Exit(-1)
	}
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate MySQL model from SQL ddl",
	Long:  "Generate MySQL model from SQL ddl",
	Run: func(cmd *cobra.Command, args []string) {
		Model(modelDDLFile, modelPackageName)
	},
}

const modelTemplate = `// Code generated from [[ .SourceFile ]] by gpa. DO NOT EDIT.

package [[ .PackageName ]]

[[ if .Imports ]]
import(
[[ range .Imports -]]
	"[[ . ]]"
[[ end -]]
)
[[ end -]]

[[ range $s := .Structs ]]
func ([[ .StructName ]]) TableName() string {
	return "[[ .TableName ]]";
}

type [[ $s.StructName ]] struct {
[[- range $d := $s.Definitions ]]
	[[ $d.Name ]] [[ $d.Type ]] [[ $d.Tags ]] [[ if $d.SQLComment ]] // [[ $d.SQLComment ]] [[ end -]]
[[ end ]]
}

func New[[ .StructName ]]() *[[ .StructName ]] {
	return &[[ .StructName ]]{}
}
[[ end -]]
`

func Model(sourceFile, packageName string) {
	if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
		logger.Log.Errorf("ddl file `%s` not exists", sourceFile)
		os.Exit(-1)
	}

	fs, err := antlr.NewFileStream(sourceFile)
	if err != nil {
		logger.Log.Errorf("antlr.NewFileStream err: %s", err.Error())
		os.Exit(-2)
	}

	p := parser.NewCreateTableParser(antlr.NewCommonTokenStream(parser.NewDDLLexer(fs), antlr.TokenDefaultChannel))
	p.BuildParseTrees = true

	// antlr and tpl share one model
	createTableListener := listener.CreateTableListener{
		SourceFile:  filepath.Base(sourceFile),
		PackageName: packageName,
	}
	antlr.ParseTreeWalkerDefault.Walk(&createTableListener, p.Root())

	processStructNames(&createTableListener, func(s string) string {
		return strcase.ToCamel(s)
	})
	processDefinitionNames(&createTableListener, func(s string) string {
		return strcase.ToCamel(s)
	})
	processTypes(&createTableListener)
	processImports(&createTableListener)
	processStructAlignment(&createTableListener)
	processTags(&createTableListener, &TagGenerator{
		Prefix: "db",
		Map: func(d *listener.Definition) string {
			return strings.ToLower(d.SQLName)
		},
	}, &TagGenerator{
		Prefix: "json",
		Map: func(d *listener.Definition) string {
			return strings.ToLower(d.SQLName)
		},
	})

	t, err := template.New("MySQLModel").Delims("[[", "]]").Parse(modelTemplate)
	if err != nil {
		logger.Log.Errorf("init text/template err: %s", err.Error())
		os.Exit(-3)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, createTableListener); err != nil {
		logger.Log.Errorf("render text/template err: %s", err.Error())
		os.Exit(-4)
	}

	targetDir := path.Dir(modelTargetFile)
	if _, err := os.Stat(targetDir); err != nil && os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(targetDir, 0644); mkdirErr != nil {
			logger.Log.Errorf("mkdir `%s` err: %s", targetDir, err.Error())
			os.Exit(-5)
		}
	}

	if err := ioutil.WriteFile(modelTargetFile, buf.Bytes(), 0644); err != nil {
		logger.Log.Errorf("dump file `%s` err: %s", modelTargetFile, err.Error())
		os.Exit(-6)
	}
}

// processStructNames generate go struct names
func processStructNames(l *listener.CreateTableListener, mappers ...func(string) string) {
	for _, s := range l.Structs {
		s.StructName = s.TableName
		for _, mapper := range mappers {
			s.StructName = mapper(s.StructName)
		}
	}
}

// processDefinitionNames generate go struct field names
func processDefinitionNames(l *listener.CreateTableListener, mappers ...func(string) string) {
	for _, s := range l.Structs {
		for _, d := range s.Definitions {
			d.Name = d.SQLName
			for _, mapper := range mappers {
				d.Name = mapper(d.Name)
			}
		}
	}
}

// processTypes generate go struct field types
func processTypes(l *listener.CreateTableListener) {
	for _, s := range l.Structs {
		for _, d := range s.Definitions {
			if d.SQLType == "tinyint" {
				if d.NotNull {
					if d.Unsigned {
						d.Type = "uint8"
						d.Size = unsafe.Sizeof(uint8(0))
					} else {
						d.Type = "int8"
						d.Size = unsafe.Sizeof(int8(0))
					}
				} else {
					d.Type = "sql.NullInt64"
					d.Size = unsafe.Sizeof(sql.NullInt64{})
				}
			} else if d.SQLType == "smallint" {
				if d.NotNull {
					if d.Unsigned {
						d.Type = "uint16"
						d.Size = unsafe.Sizeof(uint16(0))
					} else {
						d.Type = "int16"
						d.Size = unsafe.Sizeof(int16(0))
					}
				} else {
					d.Type = "sql.NullInt64"
					d.Size = unsafe.Sizeof(sql.NullInt64{})
				}
			} else if d.SQLType == "int" || d.SQLType == "integer" || d.SQLType == "mediumint" {
				if d.NotNull {
					if d.Unsigned {
						d.Type = "uint32"
						d.Size = unsafe.Sizeof(uint32(0))
					} else {
						d.Type = "int32"
						d.Size = unsafe.Sizeof(int32(0))
					}
				} else {
					d.Type = "sql.NullInt64"
					d.Size = unsafe.Sizeof(sql.NullInt64{})
				}
			} else if d.SQLType == "bigint" {
				if d.NotNull {
					if d.Unsigned {
						d.Type = "uint64"
						d.Size = unsafe.Sizeof(uint64(0))
					} else {
						d.Type = "int64"
						d.Size = unsafe.Sizeof(int64(0))
					}
				} else {
					d.Type = "sql.NullInt64"
					d.Size = unsafe.Sizeof(sql.NullInt64{})
				}
			} else if d.SQLType == "bit" ||
				d.SQLType == "bool" ||
				d.SQLType == "boolean" {
				if d.NotNull {
					d.Type = "uint8"
					d.Size = unsafe.Sizeof(uint8(0))
				} else {
					d.Type = "sql.NullInt64"
					d.Size = unsafe.Sizeof(sql.NullInt64{})
				}
			} else if d.SQLType == "double" ||
				d.SQLType == "float" ||
				d.SQLType == "fixed" ||
				d.SQLType == "numeric" ||
				d.SQLType == "decimal" ||
				d.SQLType == "dec" {
				if d.NotNull {
					d.Type = "float64"
					d.Size = unsafe.Sizeof(float64(0))
				} else {
					d.Type = "sql.NullFloat64"
					d.Size = unsafe.Sizeof(sql.NullFloat64{})
				}
			} else if d.SQLType == "enum" ||
				d.SQLType == "set" ||
				d.SQLType == "char" ||
				d.SQLType == "character" ||
				d.SQLType == "varchar" ||
				d.SQLType == "tinytext" ||
				d.SQLType == "text" ||
				d.SQLType == "mediumtext" ||
				d.SQLType == "longtext" ||
				d.SQLType == "nchar" ||
				d.SQLType == "nvarchar" {
				if d.NotNull {
					d.Type = "string"
					d.Size = unsafe.Sizeof("")
				} else {
					d.Type = "sql.NullString"
					d.Size = unsafe.Sizeof(sql.NullString{})
				}
			} else if d.SQLType == "timestamp" ||
				d.SQLType == "date" ||
				d.SQLType == "datetime" ||
				d.SQLType == "year" ||
				d.SQLType == "time" {
				if d.NotNull {
					d.Type = "time.Time"
					d.Size = unsafe.Sizeof(time.Time{})
				} else {
					d.Type = "mysql.NullTime"
					d.Size = unsafe.Sizeof(mysql.NullTime{})
				}
			} else if d.SQLType == "binary" ||
				d.SQLType == "varbinary" ||
				d.SQLType == "tinyblob" ||
				d.SQLType == "blob" ||
				d.SQLType == "mediumblob" ||
				d.SQLType == "longblob" {
				d.Type = "[]byte"
				d.Size = unsafe.Sizeof([]byte(""))
			} else if d.SQLType == "serial" {
				d.Type = "uint64"
				d.NotNull = true
				d.Size = unsafe.Sizeof(uint64(0))
			}
		}
	}
}

// processImports merge imports
func processImports(l *listener.CreateTableListener) {
	imports, importMapping := make([]string, 0), make(map[string]bool)
	for _, goStruct := range l.Structs {
		for _, definition := range goStruct.Definitions {
			if strings.HasSuffix(definition.Type, "time.Time") {
				if _, exists := importMapping["time"]; !exists {
					importMapping["time"] = true
					imports = append(imports, "time")
				}
			} else if strings.HasPrefix(definition.Type, "sql.Null") {
				if _, exists := importMapping["database/sql"]; !exists {
					importMapping["database/sql"] = true
					imports = append(imports, "database/sql")
				}
			} else if definition.Type == "mysql.NullTime" {
				if _, exists := importMapping["github.com/go-sql-driver/mysql"]; !exists {
					importMapping["github.com/go-sql-driver/mysql"] = true
					imports = append(imports, "github.com/go-sql-driver/mysql")
				}
			}
		}
	}
	l.Imports = imports
}

// processStructAlignment reduce memory costs
func processStructAlignment(l *listener.CreateTableListener) {
	for _, s := range l.Structs {
		sort.Slice(s.Definitions, func(i, j int) bool {
			if s.Definitions[i].Size != s.Definitions[j].Size {
				return s.Definitions[i].Size < s.Definitions[j].Size
			} else if s.Definitions[i].Type != s.Definitions[j].Type {
				return s.Definitions[i].Type < s.Definitions[j].Type
			} else {
				return s.Definitions[i].Name < s.Definitions[j].Name
			}
		})
	}
}

type TagGenerator struct {
	Prefix string
	Map    func(*listener.Definition) string
}

// processTags go struct tags
func processTags(l *listener.CreateTableListener, tagGenerators ...*TagGenerator) {
	for _, s := range l.Structs {
		for _, definition := range s.Definitions {
			tags := make([]string, 0)
			for _, tg := range tagGenerators {
				tags = append(tags, tg.Prefix+`:"`+tg.Map(definition)+`"`)
			}

			definition.Tags = func() string {
				if len(tags) == 0 {
					return ""
				}
				return "`" + strings.Join(tags, " ") + "`"
			}()
		}
	}
}
