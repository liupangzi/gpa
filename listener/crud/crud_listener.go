package crudlistener

import parser "github.com/liupangzi/gpa/parser/crud"

type CRUDListener struct {
	*parser.BaseCRUDAnnotationListener

	Annotation string

	UpdateStatement string
	DeleteStatement string
	SelectStatement string
}

func (c *CRUDListener) EnterSelectClause(ctx *parser.SelectClauseContext) {
	c.SelectStatement = ctx.GetSql().GetText()
}

func (c *CRUDListener) EnterUpdateClause(ctx *parser.UpdateClauseContext) {
	c.UpdateStatement = ctx.GetSql().GetText()
}

func (c *CRUDListener) EnterDeleteClause(ctx *parser.DeleteClauseContext) {
	c.DeleteStatement = ctx.GetSql().GetText()
}
