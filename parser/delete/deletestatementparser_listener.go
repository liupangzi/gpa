// Code generated from DeleteStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package deleteparser // DeleteStatementParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DeleteStatementParserListener is a complete listener for a parse tree produced by DeleteStatementParserParser.
type DeleteStatementParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByField is called when entering the orderByField production.
	EnterOrderByField(c *OrderByFieldContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterSubExpression is called when entering the subExpression production.
	EnterSubExpression(c *SubExpressionContext)

	// EnterAtomExpression is called when entering the atomExpression production.
	EnterAtomExpression(c *AtomExpressionContext)

	// EnterLikeExpression is called when entering the likeExpression production.
	EnterLikeExpression(c *LikeExpressionContext)

	// EnterCompareExpression is called when entering the compareExpression production.
	EnterCompareExpression(c *CompareExpressionContext)

	// EnterInExpression is called when entering the inExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterColonField is called when entering the colonField production.
	EnterColonField(c *ColonFieldContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByField is called when exiting the orderByField production.
	ExitOrderByField(c *OrderByFieldContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitSubExpression is called when exiting the subExpression production.
	ExitSubExpression(c *SubExpressionContext)

	// ExitAtomExpression is called when exiting the atomExpression production.
	ExitAtomExpression(c *AtomExpressionContext)

	// ExitLikeExpression is called when exiting the likeExpression production.
	ExitLikeExpression(c *LikeExpressionContext)

	// ExitCompareExpression is called when exiting the compareExpression production.
	ExitCompareExpression(c *CompareExpressionContext)

	// ExitInExpression is called when exiting the inExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitColonField is called when exiting the colonField production.
	ExitColonField(c *ColonFieldContext)
}
