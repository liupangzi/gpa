// Code generated from UpdateStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package updateparser // UpdateStatementParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUpdateStatementParserListener is a complete listener for a parse tree produced by UpdateStatementParserParser.
type BaseUpdateStatementParserListener struct{}

var _ UpdateStatementParserListener = &BaseUpdateStatementParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUpdateStatementParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUpdateStatementParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUpdateStatementParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUpdateStatementParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseUpdateStatementParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseUpdateStatementParserListener) ExitStatement(ctx *StatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseUpdateStatementParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseUpdateStatementParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseUpdateStatementParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseUpdateStatementParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseUpdateStatementParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseUpdateStatementParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseUpdateStatementParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseUpdateStatementParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseUpdateStatementParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseUpdateStatementParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByField is called when production orderByField is entered.
func (s *BaseUpdateStatementParserListener) EnterOrderByField(ctx *OrderByFieldContext) {}

// ExitOrderByField is called when production orderByField is exited.
func (s *BaseUpdateStatementParserListener) ExitOrderByField(ctx *OrderByFieldContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseUpdateStatementParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseUpdateStatementParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterSubExpression is called when production subExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterSubExpression(ctx *SubExpressionContext) {}

// ExitSubExpression is called when production subExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitSubExpression(ctx *SubExpressionContext) {}

// EnterAtomExpression is called when production atomExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterAtomExpression(ctx *AtomExpressionContext) {}

// ExitAtomExpression is called when production atomExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitAtomExpression(ctx *AtomExpressionContext) {}

// EnterLikeExpression is called when production likeExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterLikeExpression(ctx *LikeExpressionContext) {}

// ExitLikeExpression is called when production likeExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitLikeExpression(ctx *LikeExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseUpdateStatementParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseUpdateStatementParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterColonField is called when production colonField is entered.
func (s *BaseUpdateStatementParserListener) EnterColonField(ctx *ColonFieldContext) {}

// ExitColonField is called when production colonField is exited.
func (s *BaseUpdateStatementParserListener) ExitColonField(ctx *ColonFieldContext) {}
