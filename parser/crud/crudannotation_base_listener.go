// Code generated from CRUDAnnotation.g4 by ANTLR 4.7.2. DO NOT EDIT.

package crudannotation // CRUDAnnotation
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCRUDAnnotationListener is a complete listener for a parse tree produced by CRUDAnnotationParser.
type BaseCRUDAnnotationListener struct{}

var _ CRUDAnnotationListener = &BaseCRUDAnnotationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCRUDAnnotationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCRUDAnnotationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCRUDAnnotationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCRUDAnnotationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCrud is called when production crud is entered.
func (s *BaseCRUDAnnotationListener) EnterCrud(ctx *CrudContext) {}

// ExitCrud is called when production crud is exited.
func (s *BaseCRUDAnnotationListener) ExitCrud(ctx *CrudContext) {}

// EnterInsertClause is called when production insertClause is entered.
func (s *BaseCRUDAnnotationListener) EnterInsertClause(ctx *InsertClauseContext) {}

// ExitInsertClause is called when production insertClause is exited.
func (s *BaseCRUDAnnotationListener) ExitInsertClause(ctx *InsertClauseContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseCRUDAnnotationListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseCRUDAnnotationListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterUpdateClause is called when production updateClause is entered.
func (s *BaseCRUDAnnotationListener) EnterUpdateClause(ctx *UpdateClauseContext) {}

// ExitUpdateClause is called when production updateClause is exited.
func (s *BaseCRUDAnnotationListener) ExitUpdateClause(ctx *UpdateClauseContext) {}

// EnterDeleteClause is called when production deleteClause is entered.
func (s *BaseCRUDAnnotationListener) EnterDeleteClause(ctx *DeleteClauseContext) {}

// ExitDeleteClause is called when production deleteClause is exited.
func (s *BaseCRUDAnnotationListener) ExitDeleteClause(ctx *DeleteClauseContext) {}
