package ddllistener

import (
	"strings"

	parser "github.com/liupangzi/gpa/parser/ddl"
)

type CreateTableListener struct {
	*parser.BaseCreateTableParserListener

	SourceFile  string
	TargetFile  string
	PackageName string
	Imports     []string
	Structs     []*GoStruct
}

type GoStruct struct {
	TableName    string
	StructName   string
	TableComment string
	Definitions  []*Definition
}

type Definition struct {
	SQLName       string // SQL field name
	SQLType       string // SQL field type
	Name          string // Go struct field name
	Type          string // Go struct field type
	NotNull       bool
	Unsigned      bool
	SQLComment    string // SQL comment
	InlineComment string // inline comment
	Tags          string
	Size          uintptr
}

func (c *CreateTableListener) EnterTableName(ctx *parser.TableNameContext) {
	nextStruct := &GoStruct{}
	nextStruct.TableName = strings.ReplaceAll(ctx.GetText(), "`", "")

	c.Structs = append(c.Structs, nextStruct)
}

func (c *CreateTableListener) EnterCreateDefinition(ctx *parser.CreateDefinitionContext) {
	if ctx.GetField() == nil {
		return
	}

	fieldName := strings.ReplaceAll(ctx.GetField().GetText(), "`", "")
	currentStruct := c.Structs[len(c.Structs)-1]

	currentStruct.Definitions = append(currentStruct.Definitions, &Definition{
		SQLName: fieldName,
	})
}

func (c *CreateTableListener) EnterDataType(ctx *parser.DataTypeContext) {
	currentStruct := c.Structs[len(c.Structs)-1]
	currentDefinition := currentStruct.Definitions[len(currentStruct.Definitions)-1]

	currentDefinition.SQLType = strings.ToLower(ctx.GetTypeName().GetText())
	currentDefinition.Unsigned = ctx.GetSign() != nil && strings.ToLower(ctx.GetSign().GetText()) == "unsigned"
}

func (c *CreateTableListener) EnterNullNotNull(ctx *parser.NullNotNullContext) {
	currentStruct := c.Structs[len(c.Structs)-1]
	currentDefinition := currentStruct.Definitions[len(currentStruct.Definitions)-1]

	currentDefinition.NotNull = strings.ToLower(ctx.GetText()) == "notnull"
}

func (c *CreateTableListener) EnterComment(ctx *parser.CommentContext) {
	currentStruct := c.Structs[len(c.Structs)-1]
	currentDefinition := currentStruct.Definitions[len(currentStruct.Definitions)-1]

	currentDefinition.SQLComment = ctx.GetContent().GetText()[1 : len(ctx.GetContent().GetText())-1]
}

func (c *CreateTableListener) EnterTableOption(ctx *parser.TableOptionContext) {
	if ctx.GetTableComment() != nil {
		c.Structs[len(c.Structs)-1].TableComment = ctx.GetTableComment().GetText()
	}
}
