// Code generated from SelectStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package selectparser // SelectStatementParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SelectStatementParserListener is a complete listener for a parse tree produced by SelectStatementParserParser.
type SelectStatementParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSelectField is called when entering the selectField production.
	EnterSelectField(c *SelectFieldContext)

	// EnterSelectAsPrefix is called when entering the selectAsPrefix production.
	EnterSelectAsPrefix(c *SelectAsPrefixContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupByField is called when entering the groupByField production.
	EnterGroupByField(c *GroupByFieldContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByField is called when entering the orderByField production.
	EnterOrderByField(c *OrderByFieldContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOffsetValue is called when entering the offsetValue production.
	EnterOffsetValue(c *OffsetValueContext)

	// EnterLimitValue is called when entering the limitValue production.
	EnterLimitValue(c *LimitValueContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterSubExpression is called when entering the subExpression production.
	EnterSubExpression(c *SubExpressionContext)

	// EnterAtomExpression is called when entering the atomExpression production.
	EnterAtomExpression(c *AtomExpressionContext)

	// EnterInExpression is called when entering the inExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterGroupFunction is called when entering the groupFunction production.
	EnterGroupFunction(c *GroupFunctionContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSelectField is called when exiting the selectField production.
	ExitSelectField(c *SelectFieldContext)

	// ExitSelectAsPrefix is called when exiting the selectAsPrefix production.
	ExitSelectAsPrefix(c *SelectAsPrefixContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupByField is called when exiting the groupByField production.
	ExitGroupByField(c *GroupByFieldContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByField is called when exiting the orderByField production.
	ExitOrderByField(c *OrderByFieldContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOffsetValue is called when exiting the offsetValue production.
	ExitOffsetValue(c *OffsetValueContext)

	// ExitLimitValue is called when exiting the limitValue production.
	ExitLimitValue(c *LimitValueContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitSubExpression is called when exiting the subExpression production.
	ExitSubExpression(c *SubExpressionContext)

	// ExitAtomExpression is called when exiting the atomExpression production.
	ExitAtomExpression(c *AtomExpressionContext)

	// ExitInExpression is called when exiting the inExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitGroupFunction is called when exiting the groupFunction production.
	ExitGroupFunction(c *GroupFunctionContext)
}
