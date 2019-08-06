// Code generated from SelectStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package selectparser // SelectStatementParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSelectStatementParserListener is a complete listener for a parse tree produced by SelectStatementParserParser.
type BaseSelectStatementParserListener struct{}

var _ SelectStatementParserListener = &BaseSelectStatementParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSelectStatementParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSelectStatementParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSelectStatementParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSelectStatementParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSelectStatementParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSelectStatementParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseSelectStatementParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseSelectStatementParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSelectField is called when production selectField is entered.
func (s *BaseSelectStatementParserListener) EnterSelectField(ctx *SelectFieldContext) {}

// ExitSelectField is called when production selectField is exited.
func (s *BaseSelectStatementParserListener) ExitSelectField(ctx *SelectFieldContext) {}

// EnterSelectAsPrefix is called when production selectAsPrefix is entered.
func (s *BaseSelectStatementParserListener) EnterSelectAsPrefix(ctx *SelectAsPrefixContext) {}

// ExitSelectAsPrefix is called when production selectAsPrefix is exited.
func (s *BaseSelectStatementParserListener) ExitSelectAsPrefix(ctx *SelectAsPrefixContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSelectStatementParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSelectStatementParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSelectStatementParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSelectStatementParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSelectStatementParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSelectStatementParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupByField is called when production groupByField is entered.
func (s *BaseSelectStatementParserListener) EnterGroupByField(ctx *GroupByFieldContext) {}

// ExitGroupByField is called when production groupByField is exited.
func (s *BaseSelectStatementParserListener) ExitGroupByField(ctx *GroupByFieldContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSelectStatementParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSelectStatementParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByField is called when production orderByField is entered.
func (s *BaseSelectStatementParserListener) EnterOrderByField(ctx *OrderByFieldContext) {}

// ExitOrderByField is called when production orderByField is exited.
func (s *BaseSelectStatementParserListener) ExitOrderByField(ctx *OrderByFieldContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSelectStatementParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSelectStatementParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetValue is called when production offsetValue is entered.
func (s *BaseSelectStatementParserListener) EnterOffsetValue(ctx *OffsetValueContext) {}

// ExitOffsetValue is called when production offsetValue is exited.
func (s *BaseSelectStatementParserListener) ExitOffsetValue(ctx *OffsetValueContext) {}

// EnterLimitValue is called when production limitValue is entered.
func (s *BaseSelectStatementParserListener) EnterLimitValue(ctx *LimitValueContext) {}

// ExitLimitValue is called when production limitValue is exited.
func (s *BaseSelectStatementParserListener) ExitLimitValue(ctx *LimitValueContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseSelectStatementParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseSelectStatementParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseSelectStatementParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseSelectStatementParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterSubExpression is called when production subExpression is entered.
func (s *BaseSelectStatementParserListener) EnterSubExpression(ctx *SubExpressionContext) {}

// ExitSubExpression is called when production subExpression is exited.
func (s *BaseSelectStatementParserListener) ExitSubExpression(ctx *SubExpressionContext) {}

// EnterAtomExpression is called when production atomExpression is entered.
func (s *BaseSelectStatementParserListener) EnterAtomExpression(ctx *AtomExpressionContext) {}

// ExitAtomExpression is called when production atomExpression is exited.
func (s *BaseSelectStatementParserListener) ExitAtomExpression(ctx *AtomExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseSelectStatementParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseSelectStatementParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterGroupFunction is called when production groupFunction is entered.
func (s *BaseSelectStatementParserListener) EnterGroupFunction(ctx *GroupFunctionContext) {}

// ExitGroupFunction is called when production groupFunction is exited.
func (s *BaseSelectStatementParserListener) ExitGroupFunction(ctx *GroupFunctionContext) {}
