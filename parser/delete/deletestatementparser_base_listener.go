// Code generated from DeleteStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package deleteparser // DeleteStatementParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDeleteStatementParserListener is a complete listener for a parse tree produced by DeleteStatementParserParser.
type BaseDeleteStatementParserListener struct{}

var _ DeleteStatementParserListener = &BaseDeleteStatementParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDeleteStatementParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDeleteStatementParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDeleteStatementParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDeleteStatementParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDeleteStatementParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDeleteStatementParserListener) ExitStatement(ctx *StatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDeleteStatementParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDeleteStatementParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseDeleteStatementParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseDeleteStatementParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseDeleteStatementParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseDeleteStatementParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseDeleteStatementParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseDeleteStatementParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByField is called when production orderByField is entered.
func (s *BaseDeleteStatementParserListener) EnterOrderByField(ctx *OrderByFieldContext) {}

// ExitOrderByField is called when production orderByField is exited.
func (s *BaseDeleteStatementParserListener) ExitOrderByField(ctx *OrderByFieldContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseDeleteStatementParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseDeleteStatementParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterSubExpression is called when production subExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterSubExpression(ctx *SubExpressionContext) {}

// ExitSubExpression is called when production subExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitSubExpression(ctx *SubExpressionContext) {}

// EnterAtomExpression is called when production atomExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterAtomExpression(ctx *AtomExpressionContext) {}

// ExitAtomExpression is called when production atomExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitAtomExpression(ctx *AtomExpressionContext) {}

// EnterLikeExpression is called when production likeExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterLikeExpression(ctx *LikeExpressionContext) {}

// ExitLikeExpression is called when production likeExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitLikeExpression(ctx *LikeExpressionContext) {}

// EnterCompareExpression is called when production compareExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterCompareExpression(ctx *CompareExpressionContext) {}

// ExitCompareExpression is called when production compareExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitCompareExpression(ctx *CompareExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseDeleteStatementParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseDeleteStatementParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterColonField is called when production colonField is entered.
func (s *BaseDeleteStatementParserListener) EnterColonField(ctx *ColonFieldContext) {}

// ExitColonField is called when production colonField is exited.
func (s *BaseDeleteStatementParserListener) ExitColonField(ctx *ColonFieldContext) {}
