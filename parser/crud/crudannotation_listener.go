// Code generated from CRUDAnnotation.g4 by ANTLR 4.7.2. DO NOT EDIT.

package crudannotation // CRUDAnnotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

// CRUDAnnotationListener is a complete listener for a parse tree produced by CRUDAnnotationParser.
type CRUDAnnotationListener interface {
	antlr.ParseTreeListener

	// EnterCrud is called when entering the crud production.
	EnterCrud(c *CrudContext)

	// EnterInsertClause is called when entering the insertClause production.
	EnterInsertClause(c *InsertClauseContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterUpdateClause is called when entering the updateClause production.
	EnterUpdateClause(c *UpdateClauseContext)

	// EnterDeleteClause is called when entering the deleteClause production.
	EnterDeleteClause(c *DeleteClauseContext)

	// ExitCrud is called when exiting the crud production.
	ExitCrud(c *CrudContext)

	// ExitInsertClause is called when exiting the insertClause production.
	ExitInsertClause(c *InsertClauseContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitUpdateClause is called when exiting the updateClause production.
	ExitUpdateClause(c *UpdateClauseContext)

	// ExitDeleteClause is called when exiting the deleteClause production.
	ExitDeleteClause(c *DeleteClauseContext)
}
