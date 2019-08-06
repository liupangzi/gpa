// Code generated from SelectStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package selectparser // SelectStatementParser
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 219,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 5, 2, 42, 10, 2, 3, 2, 5, 2, 45,
	10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 52, 10, 3, 3, 3, 3, 3, 5, 3,
	56, 10, 3, 3, 3, 5, 3, 59, 10, 3, 3, 3, 5, 3, 62, 10, 3, 3, 3, 5, 3, 65,
	10, 3, 3, 4, 5, 4, 68, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 74, 10, 4,
	3, 4, 7, 4, 77, 10, 4, 12, 4, 14, 4, 80, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 7, 8, 99, 10, 8, 12, 8, 14, 8, 102, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	7, 8, 108, 10, 8, 12, 8, 14, 8, 111, 11, 8, 5, 8, 113, 10, 8, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 122, 10, 10, 12, 10, 14, 10,
	125, 11, 10, 3, 11, 3, 11, 5, 11, 129, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 5, 12, 135, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 142,
	10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 150, 10, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 155, 10, 15, 7, 15, 157, 10, 15, 12, 15, 14, 15,
	160, 11, 15, 3, 16, 3, 16, 5, 16, 164, 10, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 169, 10, 16, 7, 16, 171, 10, 16, 12, 16, 14, 16, 174, 11, 16, 3, 17,
	3, 17, 3, 17, 5, 17, 179, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 189, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 194,
	10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 207, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 212, 10,
	19, 12, 19, 14, 19, 215, 11, 19, 3, 20, 3, 20, 3, 20, 2, 2, 21, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 11, 3,
	2, 42, 43, 3, 2, 25, 26, 3, 2, 30, 31, 4, 2, 14, 14, 40, 40, 4, 2, 14,
	14, 41, 41, 3, 2, 12, 13, 4, 2, 14, 14, 40, 41, 3, 2, 15, 18, 3, 2, 35,
	39, 2, 230, 2, 41, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 67, 3, 2, 2, 2, 8,
	81, 3, 2, 2, 2, 10, 87, 3, 2, 2, 2, 12, 90, 3, 2, 2, 2, 14, 93, 3, 2, 2,
	2, 16, 114, 3, 2, 2, 2, 18, 116, 3, 2, 2, 2, 20, 126, 3, 2, 2, 2, 22, 130,
	3, 2, 2, 2, 24, 143, 3, 2, 2, 2, 26, 145, 3, 2, 2, 2, 28, 149, 3, 2, 2,
	2, 30, 163, 3, 2, 2, 2, 32, 175, 3, 2, 2, 2, 34, 206, 3, 2, 2, 2, 36, 208,
	3, 2, 2, 2, 38, 216, 3, 2, 2, 2, 40, 42, 5, 4, 3, 2, 41, 40, 3, 2, 2, 2,
	41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2, 2, 43, 45, 7, 9, 2, 2, 44, 43, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 47, 7, 2, 2, 3, 47,
	3, 3, 2, 2, 2, 48, 51, 7, 19, 2, 2, 49, 52, 7, 11, 2, 2, 50, 52, 5, 6,
	4, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55,
	5, 10, 6, 2, 54, 56, 5, 12, 7, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2,
	2, 56, 58, 3, 2, 2, 2, 57, 59, 5, 14, 8, 2, 58, 57, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 62, 5, 18, 10, 2, 61, 60, 3, 2, 2,
	2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 65, 5, 22, 12, 2, 64, 63,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 5, 3, 2, 2, 2, 66, 68, 5, 8, 5, 2,
	67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 9,
	2, 2, 2, 70, 78, 3, 2, 2, 2, 71, 73, 7, 10, 2, 2, 72, 74, 5, 8, 5, 2, 73,
	72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 77, 9, 2, 2,
	2, 76, 71, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79,
	3, 2, 2, 2, 79, 7, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 82, 5, 38, 20, 2,
	82, 83, 7, 7, 2, 2, 83, 84, 9, 2, 2, 2, 84, 85, 7, 8, 2, 2, 85, 86, 7,
	34, 2, 2, 86, 9, 3, 2, 2, 2, 87, 88, 7, 21, 2, 2, 88, 89, 9, 2, 2, 2, 89,
	11, 3, 2, 2, 2, 90, 91, 7, 22, 2, 2, 91, 92, 5, 28, 15, 2, 92, 13, 3, 2,
	2, 2, 93, 94, 7, 28, 2, 2, 94, 95, 7, 29, 2, 2, 95, 100, 5, 16, 9, 2, 96,
	97, 7, 10, 2, 2, 97, 99, 5, 16, 9, 2, 98, 96, 3, 2, 2, 2, 99, 102, 3, 2,
	2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 112, 3, 2, 2, 2,
	102, 100, 3, 2, 2, 2, 103, 104, 7, 23, 2, 2, 104, 109, 5, 34, 18, 2, 105,
	106, 9, 3, 2, 2, 106, 108, 5, 34, 18, 2, 107, 105, 3, 2, 2, 2, 108, 111,
	3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 113, 3, 2,
	2, 2, 111, 109, 3, 2, 2, 2, 112, 103, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2,
	113, 15, 3, 2, 2, 2, 114, 115, 9, 2, 2, 2, 115, 17, 3, 2, 2, 2, 116, 117,
	7, 27, 2, 2, 117, 118, 7, 29, 2, 2, 118, 123, 5, 20, 11, 2, 119, 120, 7,
	10, 2, 2, 120, 122, 5, 20, 11, 2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2,
	2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 19, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 126, 128, 9, 2, 2, 2, 127, 129, 9, 4, 2, 2, 128,
	127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 21, 3, 2, 2, 2, 130, 141, 7,
	33, 2, 2, 131, 132, 5, 24, 13, 2, 132, 133, 7, 10, 2, 2, 133, 135, 3, 2,
	2, 2, 134, 131, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2,
	136, 142, 5, 26, 14, 2, 137, 138, 5, 26, 14, 2, 138, 139, 7, 32, 2, 2,
	139, 140, 5, 24, 13, 2, 140, 142, 3, 2, 2, 2, 141, 134, 3, 2, 2, 2, 141,
	137, 3, 2, 2, 2, 142, 23, 3, 2, 2, 2, 143, 144, 9, 5, 2, 2, 144, 25, 3,
	2, 2, 2, 145, 146, 9, 5, 2, 2, 146, 27, 3, 2, 2, 2, 147, 150, 5, 32, 17,
	2, 148, 150, 5, 30, 16, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2,
	150, 158, 3, 2, 2, 2, 151, 154, 7, 26, 2, 2, 152, 155, 5, 32, 17, 2, 153,
	155, 5, 30, 16, 2, 154, 152, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 157,
	3, 2, 2, 2, 156, 151, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 158, 159, 3, 2, 2, 2, 159, 29, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2,
	161, 164, 5, 32, 17, 2, 162, 164, 5, 34, 18, 2, 163, 161, 3, 2, 2, 2, 163,
	162, 3, 2, 2, 2, 164, 172, 3, 2, 2, 2, 165, 168, 7, 25, 2, 2, 166, 169,
	5, 32, 17, 2, 167, 169, 5, 34, 18, 2, 168, 166, 3, 2, 2, 2, 168, 167, 3,
	2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 165, 3, 2, 2, 2, 171, 174, 3, 2, 2,
	2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 31, 3, 2, 2, 2, 174,
	172, 3, 2, 2, 2, 175, 178, 7, 7, 2, 2, 176, 179, 5, 30, 16, 2, 177, 179,
	5, 28, 15, 2, 178, 176, 3, 2, 2, 2, 178, 177, 3, 2, 2, 2, 179, 180, 3,
	2, 2, 2, 180, 181, 7, 8, 2, 2, 181, 33, 3, 2, 2, 2, 182, 183, 9, 2, 2,
	2, 183, 184, 7, 24, 2, 2, 184, 207, 9, 6, 2, 2, 185, 186, 9, 2, 2, 2, 186,
	188, 7, 3, 2, 2, 187, 189, 7, 5, 2, 2, 188, 187, 3, 2, 2, 2, 188, 189,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 207, 7, 6, 2, 2, 191, 193, 9, 2,
	2, 2, 192, 194, 7, 5, 2, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2,
	194, 195, 3, 2, 2, 2, 195, 196, 7, 4, 2, 2, 196, 197, 7, 7, 2, 2, 197,
	198, 5, 36, 19, 2, 198, 199, 7, 8, 2, 2, 199, 207, 3, 2, 2, 2, 200, 201,
	9, 2, 2, 2, 201, 202, 9, 7, 2, 2, 202, 207, 9, 8, 2, 2, 203, 204, 9, 2,
	2, 2, 204, 205, 9, 9, 2, 2, 205, 207, 9, 5, 2, 2, 206, 182, 3, 2, 2, 2,
	206, 185, 3, 2, 2, 2, 206, 191, 3, 2, 2, 2, 206, 200, 3, 2, 2, 2, 206,
	203, 3, 2, 2, 2, 207, 35, 3, 2, 2, 2, 208, 213, 9, 8, 2, 2, 209, 210, 7,
	10, 2, 2, 210, 212, 9, 8, 2, 2, 211, 209, 3, 2, 2, 2, 212, 215, 3, 2, 2,
	2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 37, 3, 2, 2, 2, 215,
	213, 3, 2, 2, 2, 216, 217, 9, 10, 2, 2, 217, 39, 3, 2, 2, 2, 30, 41, 44,
	51, 55, 58, 61, 64, 67, 73, 78, 100, 109, 112, 123, 128, 134, 141, 149,
	154, 158, 163, 168, 172, 178, 188, 193, 206, 213,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "';'", "','", "'*'", "'='", "'!='", "'?'",
	"'>'", "'>='", "'<'", "'<='",
}
var symbolicNames = []string{
	"", "Is", "In", "Not", "Null", "LeftParenthesis", "RightParenthesis", "Semicolon",
	"Comma", "Asterisk", "Equal", "NotEqual", "QuestionMark", "GT", "GTE",
	"LT", "LTE", "Select", "Delete", "From", "Where", "Having", "Like", "And",
	"Or", "Order", "Group", "By", "Asc", "Desc", "Offset", "Limit", "As", "Count",
	"Sum", "Max", "Min", "Avg", "Number", "Literal", "BackQuotedString", "RawString",
	"WS",
}

var ruleNames = []string{
	"statement", "selectStatement", "selectField", "selectAsPrefix", "fromClause",
	"whereClause", "groupByClause", "groupByField", "orderByClause", "orderByField",
	"limitClause", "offsetValue", "limitValue", "orExpression", "andExpression",
	"subExpression", "atomExpression", "inExpression", "groupFunction",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SelectStatementParserParser struct {
	*antlr.BaseParser
}

func NewSelectStatementParserParser(input antlr.TokenStream) *SelectStatementParserParser {
	this := new(SelectStatementParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SelectStatementParser.g4"

	return this
}

// SelectStatementParserParser tokens.
const (
	SelectStatementParserParserEOF              = antlr.TokenEOF
	SelectStatementParserParserIs               = 1
	SelectStatementParserParserIn               = 2
	SelectStatementParserParserNot              = 3
	SelectStatementParserParserNull             = 4
	SelectStatementParserParserLeftParenthesis  = 5
	SelectStatementParserParserRightParenthesis = 6
	SelectStatementParserParserSemicolon        = 7
	SelectStatementParserParserComma            = 8
	SelectStatementParserParserAsterisk         = 9
	SelectStatementParserParserEqual            = 10
	SelectStatementParserParserNotEqual         = 11
	SelectStatementParserParserQuestionMark     = 12
	SelectStatementParserParserGT               = 13
	SelectStatementParserParserGTE              = 14
	SelectStatementParserParserLT               = 15
	SelectStatementParserParserLTE              = 16
	SelectStatementParserParserSelect           = 17
	SelectStatementParserParserDelete           = 18
	SelectStatementParserParserFrom             = 19
	SelectStatementParserParserWhere            = 20
	SelectStatementParserParserHaving           = 21
	SelectStatementParserParserLike             = 22
	SelectStatementParserParserAnd              = 23
	SelectStatementParserParserOr               = 24
	SelectStatementParserParserOrder            = 25
	SelectStatementParserParserGroup            = 26
	SelectStatementParserParserBy               = 27
	SelectStatementParserParserAsc              = 28
	SelectStatementParserParserDesc             = 29
	SelectStatementParserParserOffset           = 30
	SelectStatementParserParserLimit            = 31
	SelectStatementParserParserAs               = 32
	SelectStatementParserParserCount            = 33
	SelectStatementParserParserSum              = 34
	SelectStatementParserParserMax              = 35
	SelectStatementParserParserMin              = 36
	SelectStatementParserParserAvg              = 37
	SelectStatementParserParserNumber           = 38
	SelectStatementParserParserLiteral          = 39
	SelectStatementParserParserBackQuotedString = 40
	SelectStatementParserParserRawString        = 41
	SelectStatementParserParserWS               = 42
)

// SelectStatementParserParser rules.
const (
	SelectStatementParserParserRULE_statement       = 0
	SelectStatementParserParserRULE_selectStatement = 1
	SelectStatementParserParserRULE_selectField     = 2
	SelectStatementParserParserRULE_selectAsPrefix  = 3
	SelectStatementParserParserRULE_fromClause      = 4
	SelectStatementParserParserRULE_whereClause     = 5
	SelectStatementParserParserRULE_groupByClause   = 6
	SelectStatementParserParserRULE_groupByField    = 7
	SelectStatementParserParserRULE_orderByClause   = 8
	SelectStatementParserParserRULE_orderByField    = 9
	SelectStatementParserParserRULE_limitClause     = 10
	SelectStatementParserParserRULE_offsetValue     = 11
	SelectStatementParserParserRULE_limitValue      = 12
	SelectStatementParserParserRULE_orExpression    = 13
	SelectStatementParserParserRULE_andExpression   = 14
	SelectStatementParserParserRULE_subExpression   = 15
	SelectStatementParserParserRULE_atomExpression  = 16
	SelectStatementParserParserRULE_inExpression    = 17
	SelectStatementParserParserRULE_groupFunction   = 18
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserEOF, 0)
}

func (s *StatementContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *StatementContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSemicolon, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SelectStatementParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SelectStatementParserParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserSelect {
		{
			p.SetState(38)
			p.SelectStatement()
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserSemicolon {
		{
			p.SetState(41)
			p.Match(SelectStatementParserParserSemicolon)
		}

	}
	{
		p.SetState(44)
		p.Match(SelectStatementParserParserEOF)
	}

	return localctx
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) Select() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSelect, 0)
}

func (s *SelectStatementContext) FromClause() IFromClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *SelectStatementContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAsterisk, 0)
}

func (s *SelectStatementContext) SelectField() ISelectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectFieldContext)
}

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupByClause() IGroupByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupByClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (p *SelectStatementParserParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SelectStatementParserParserRULE_selectStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(SelectStatementParserParserSelect)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserAsterisk:
		{
			p.SetState(47)
			p.Match(SelectStatementParserParserAsterisk)
		}

	case SelectStatementParserParserCount, SelectStatementParserParserSum, SelectStatementParserParserMax, SelectStatementParserParserMin, SelectStatementParserParserAvg, SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
		{
			p.SetState(48)
			p.SelectField()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(51)
		p.FromClause()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserWhere {
		{
			p.SetState(52)
			p.WhereClause()
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserGroup {
		{
			p.SetState(55)
			p.GroupByClause()
		}

	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserOrder {
		{
			p.SetState(58)
			p.OrderByClause()
		}

	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserLimit {
		{
			p.SetState(61)
			p.LimitClause()
		}

	}

	return localctx
}

// ISelectFieldContext is an interface to support dynamic dispatch.
type ISelectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectFieldContext differentiates from other interfaces.
	IsSelectFieldContext()
}

type SelectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFieldContext() *SelectFieldContext {
	var p = new(SelectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectField
	return p
}

func (*SelectFieldContext) IsSelectFieldContext() {}

func NewSelectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFieldContext {
	var p = new(SelectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectField

	return p
}

func (s *SelectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFieldContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserRawString)
}

func (s *SelectFieldContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, i)
}

func (s *SelectFieldContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserBackQuotedString)
}

func (s *SelectFieldContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, i)
}

func (s *SelectFieldContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *SelectFieldContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *SelectFieldContext) AllSelectAsPrefix() []ISelectAsPrefixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectAsPrefixContext)(nil)).Elem())
	var tst = make([]ISelectAsPrefixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectAsPrefixContext)
		}
	}

	return tst
}

func (s *SelectFieldContext) SelectAsPrefix(i int) ISelectAsPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectAsPrefixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectAsPrefixContext)
}

func (s *SelectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectField(s)
	}
}

func (s *SelectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectField(s)
	}
}

func (p *SelectStatementParserParser) SelectField() (localctx ISelectFieldContext) {
	localctx = NewSelectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SelectStatementParserParserRULE_selectField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SelectStatementParserParserCount-33))|(1<<(SelectStatementParserParserSum-33))|(1<<(SelectStatementParserParserMax-33))|(1<<(SelectStatementParserParserMin-33))|(1<<(SelectStatementParserParserAvg-33)))) != 0 {
		{
			p.SetState(64)
			p.SelectAsPrefix()
		}

	}
	{
		p.SetState(67)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(69)
			p.Match(SelectStatementParserParserComma)
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SelectStatementParserParserCount-33))|(1<<(SelectStatementParserParserSum-33))|(1<<(SelectStatementParserParserMax-33))|(1<<(SelectStatementParserParserMin-33))|(1<<(SelectStatementParserParserAvg-33)))) != 0 {
			{
				p.SetState(70)
				p.SelectAsPrefix()
			}

		}
		{
			p.SetState(73)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectAsPrefixContext is an interface to support dynamic dispatch.
type ISelectAsPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectAsPrefixContext differentiates from other interfaces.
	IsSelectAsPrefixContext()
}

type SelectAsPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectAsPrefixContext() *SelectAsPrefixContext {
	var p = new(SelectAsPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_selectAsPrefix
	return p
}

func (*SelectAsPrefixContext) IsSelectAsPrefixContext() {}

func NewSelectAsPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectAsPrefixContext {
	var p = new(SelectAsPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_selectAsPrefix

	return p
}

func (s *SelectAsPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectAsPrefixContext) GroupFunction() IGroupFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupFunctionContext)
}

func (s *SelectAsPrefixContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *SelectAsPrefixContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *SelectAsPrefixContext) As() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAs, 0)
}

func (s *SelectAsPrefixContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *SelectAsPrefixContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *SelectAsPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectAsPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectAsPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSelectAsPrefix(s)
	}
}

func (s *SelectAsPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSelectAsPrefix(s)
	}
}

func (p *SelectStatementParserParser) SelectAsPrefix() (localctx ISelectAsPrefixContext) {
	localctx = NewSelectAsPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SelectStatementParserParserRULE_selectAsPrefix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.GroupFunction()
	}
	{
		p.SetState(80)
		p.Match(SelectStatementParserParserLeftParenthesis)
	}
	{
		p.SetState(81)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(82)
		p.Match(SelectStatementParserParserRightParenthesis)
	}
	{
		p.SetState(83)
		p.Match(SelectStatementParserParserAs)
	}

	return localctx
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTable returns the table token.
	GetTable() antlr.Token

	// SetTable sets the table token.
	SetTable(antlr.Token)

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	table  antlr.Token
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) GetTable() antlr.Token { return s.table }

func (s *FromClauseContext) SetTable(v antlr.Token) { s.table = v }

func (s *FromClauseContext) From() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserFrom, 0)
}

func (s *FromClauseContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *FromClauseContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (p *SelectStatementParserParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SelectStatementParserParserRULE_fromClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(SelectStatementParserParserFrom)
	}
	{
		p.SetState(86)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*FromClauseContext).table = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*FromClauseContext).table = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) Where() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserWhere, 0)
}

func (s *WhereClauseContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *SelectStatementParserParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SelectStatementParserParserRULE_whereClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(SelectStatementParserParserWhere)
	}
	{
		p.SetState(89)
		p.OrExpression()
	}

	return localctx
}

// IGroupByClauseContext is an interface to support dynamic dispatch.
type IGroupByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByClauseContext differentiates from other interfaces.
	IsGroupByClauseContext()
}

type GroupByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByClauseContext() *GroupByClauseContext {
	var p = new(GroupByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupByClause
	return p
}

func (*GroupByClauseContext) IsGroupByClauseContext() {}

func NewGroupByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByClauseContext {
	var p = new(GroupByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupByClause

	return p
}

func (s *GroupByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByClauseContext) Group() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGroup, 0)
}

func (s *GroupByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBy, 0)
}

func (s *GroupByClauseContext) AllGroupByField() []IGroupByFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupByFieldContext)(nil)).Elem())
	var tst = make([]IGroupByFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupByFieldContext)
		}
	}

	return tst
}

func (s *GroupByClauseContext) GroupByField(i int) IGroupByFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupByFieldContext)
}

func (s *GroupByClauseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *GroupByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *GroupByClauseContext) Having() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserHaving, 0)
}

func (s *GroupByClauseContext) AllAtomExpression() []IAtomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem())
	var tst = make([]IAtomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomExpressionContext)
		}
	}

	return tst
}

func (s *GroupByClauseContext) AtomExpression(i int) IAtomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomExpressionContext)
}

func (s *GroupByClauseContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserAnd)
}

func (s *GroupByClauseContext) And(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAnd, i)
}

func (s *GroupByClauseContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserOr)
}

func (s *GroupByClauseContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOr, i)
}

func (s *GroupByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupByClause(s)
	}
}

func (s *GroupByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupByClause(s)
	}
}

func (p *SelectStatementParserParser) GroupByClause() (localctx IGroupByClauseContext) {
	localctx = NewGroupByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SelectStatementParserParserRULE_groupByClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(SelectStatementParserParserGroup)
	}
	{
		p.SetState(92)
		p.Match(SelectStatementParserParserBy)
	}
	{
		p.SetState(93)
		p.GroupByField()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(94)
			p.Match(SelectStatementParserParserComma)
		}
		{
			p.SetState(95)
			p.GroupByField()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserHaving {
		{
			p.SetState(101)
			p.Match(SelectStatementParserParserHaving)
		}
		{
			p.SetState(102)
			p.AtomExpression()
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SelectStatementParserParserAnd || _la == SelectStatementParserParserOr {
			{
				p.SetState(103)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SelectStatementParserParserAnd || _la == SelectStatementParserParserOr) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(104)
				p.AtomExpression()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IGroupByFieldContext is an interface to support dynamic dispatch.
type IGroupByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// IsGroupByFieldContext differentiates from other interfaces.
	IsGroupByFieldContext()
}

type GroupByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
}

func NewEmptyGroupByFieldContext() *GroupByFieldContext {
	var p = new(GroupByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupByField
	return p
}

func (*GroupByFieldContext) IsGroupByFieldContext() {}

func NewGroupByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByFieldContext {
	var p = new(GroupByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupByField

	return p
}

func (s *GroupByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByFieldContext) GetField() antlr.Token { return s.field }

func (s *GroupByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *GroupByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *GroupByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *GroupByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupByField(s)
	}
}

func (s *GroupByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupByField(s)
	}
}

func (p *SelectStatementParserParser) GroupByField() (localctx IGroupByFieldContext) {
	localctx = NewGroupByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SelectStatementParserParserRULE_groupByField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*GroupByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*GroupByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) Order() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOrder, 0)
}

func (s *OrderByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBy, 0)
}

func (s *OrderByClauseContext) AllOrderByField() []IOrderByFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderByFieldContext)(nil)).Elem())
	var tst = make([]IOrderByFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderByFieldContext)
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByField(i int) IOrderByFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderByFieldContext)
}

func (s *OrderByClauseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *OrderByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (p *SelectStatementParserParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SelectStatementParserParserRULE_orderByClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(SelectStatementParserParserOrder)
	}
	{
		p.SetState(115)
		p.Match(SelectStatementParserParserBy)
	}
	{
		p.SetState(116)
		p.OrderByField()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(117)
			p.Match(SelectStatementParserParserComma)
		}
		{
			p.SetState(118)
			p.OrderByField()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByFieldContext is an interface to support dynamic dispatch.
type IOrderByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// GetOrder returns the order token.
	GetOrder() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// SetOrder sets the order token.
	SetOrder(antlr.Token)

	// IsOrderByFieldContext differentiates from other interfaces.
	IsOrderByFieldContext()
}

type OrderByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	order  antlr.Token
}

func NewEmptyOrderByFieldContext() *OrderByFieldContext {
	var p = new(OrderByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orderByField
	return p
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orderByField

	return p
}

func (s *OrderByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByFieldContext) GetField() antlr.Token { return s.field }

func (s *OrderByFieldContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *OrderByFieldContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *OrderByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *OrderByFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAsc, 0)
}

func (s *OrderByFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserDesc, 0)
}

func (s *OrderByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrderByField(s)
	}
}

func (s *OrderByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrderByField(s)
	}
}

func (p *SelectStatementParserParser) OrderByField() (localctx IOrderByFieldContext) {
	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SelectStatementParserParserRULE_orderByField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OrderByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OrderByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SelectStatementParserParserAsc || _la == SelectStatementParserParserDesc {
		{
			p.SetState(125)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByFieldContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserAsc || _la == SelectStatementParserParserDesc) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderByFieldContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) Limit() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLimit, 0)
}

func (s *LimitClauseContext) LimitValue() ILimitValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitValueContext)
}

func (s *LimitClauseContext) Offset() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOffset, 0)
}

func (s *LimitClauseContext) OffsetValue() IOffsetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffsetValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOffsetValueContext)
}

func (s *LimitClauseContext) Comma() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *SelectStatementParserParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SelectStatementParserParserRULE_limitClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(SelectStatementParserParserLimit)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.SetState(132)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(129)
				p.OffsetValue()
			}
			{
				p.SetState(130)
				p.Match(SelectStatementParserParserComma)
			}

		}
		{
			p.SetState(134)
			p.LimitValue()
		}

	case 2:
		{
			p.SetState(135)
			p.LimitValue()
		}
		{
			p.SetState(136)
			p.Match(SelectStatementParserParserOffset)
		}
		{
			p.SetState(137)
			p.OffsetValue()
		}

	}

	return localctx
}

// IOffsetValueContext is an interface to support dynamic dispatch.
type IOffsetValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffsetValueContext differentiates from other interfaces.
	IsOffsetValueContext()
}

type OffsetValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetValueContext() *OffsetValueContext {
	var p = new(OffsetValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_offsetValue
	return p
}

func (*OffsetValueContext) IsOffsetValueContext() {}

func NewOffsetValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetValueContext {
	var p = new(OffsetValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_offsetValue

	return p
}

func (s *OffsetValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetValueContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *OffsetValueContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *OffsetValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOffsetValue(s)
	}
}

func (s *OffsetValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOffsetValue(s)
	}
}

func (p *SelectStatementParserParser) OffsetValue() (localctx IOffsetValueContext) {
	localctx = NewOffsetValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SelectStatementParserParserRULE_offsetValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserQuestionMark || _la == SelectStatementParserParserNumber) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILimitValueContext is an interface to support dynamic dispatch.
type ILimitValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitValueContext differentiates from other interfaces.
	IsLimitValueContext()
}

type LimitValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitValueContext() *LimitValueContext {
	var p = new(LimitValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_limitValue
	return p
}

func (*LimitValueContext) IsLimitValueContext() {}

func NewLimitValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitValueContext {
	var p = new(LimitValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_limitValue

	return p
}

func (s *LimitValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitValueContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *LimitValueContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *LimitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterLimitValue(s)
	}
}

func (s *LimitValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitLimitValue(s)
	}
}

func (p *SelectStatementParserParser) LimitValue() (localctx ILimitValueContext) {
	localctx = NewLimitValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SelectStatementParserParserRULE_limitValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SelectStatementParserParserQuestionMark || _la == SelectStatementParserParserNumber) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) AllSubExpression() []ISubExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem())
	var tst = make([]ISubExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) SubExpression(i int) ISubExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubExpressionContext)
}

func (s *OrExpressionContext) AllAndExpression() []IAndExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem())
	var tst = make([]IAndExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAndExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) AndExpression(i int) IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *OrExpressionContext) AllOr() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserOr)
}

func (s *OrExpressionContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserOr, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *SelectStatementParserParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SelectStatementParserParserRULE_orExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(145)
			p.SubExpression()
		}

	case 2:
		{
			p.SetState(146)
			p.AndExpression()
		}

	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserOr {
		{
			p.SetState(149)
			p.Match(SelectStatementParserParserOr)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(150)
				p.SubExpression()
			}

		case 2:
			{
				p.SetState(151)
				p.AndExpression()
			}

		}

		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAndExpressionContext is an interface to support dynamic dispatch.
type IAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndExpressionContext differentiates from other interfaces.
	IsAndExpressionContext()
}

type AndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExpressionContext() *AndExpressionContext {
	var p = new(AndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_andExpression

	return p
}

func (s *AndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExpressionContext) AllSubExpression() []ISubExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem())
	var tst = make([]ISubExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) SubExpression(i int) ISubExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubExpressionContext)
}

func (s *AndExpressionContext) AllAtomExpression() []IAtomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem())
	var tst = make([]IAtomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) AtomExpression(i int) IAtomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomExpressionContext)
}

func (s *AndExpressionContext) AllAnd() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserAnd)
}

func (s *AndExpressionContext) And(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAnd, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *SelectStatementParserParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SelectStatementParserParserRULE_andExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SelectStatementParserParserLeftParenthesis:
		{
			p.SetState(159)
			p.SubExpression()
		}

	case SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
		{
			p.SetState(160)
			p.AtomExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserAnd {
		{
			p.SetState(163)
			p.Match(SelectStatementParserParserAnd)
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SelectStatementParserParserLeftParenthesis:
			{
				p.SetState(164)
				p.SubExpression()
			}

		case SelectStatementParserParserBackQuotedString, SelectStatementParserParserRawString:
			{
				p.SetState(165)
				p.AtomExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISubExpressionContext is an interface to support dynamic dispatch.
type ISubExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubExpressionContext differentiates from other interfaces.
	IsSubExpressionContext()
}

type SubExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubExpressionContext() *SubExpressionContext {
	var p = new(SubExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_subExpression
	return p
}

func (*SubExpressionContext) IsSubExpressionContext() {}

func NewSubExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubExpressionContext {
	var p = new(SubExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_subExpression

	return p
}

func (s *SubExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *SubExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *SubExpressionContext) AndExpression() IAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndExpressionContext)
}

func (s *SubExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *SubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterSubExpression(s)
	}
}

func (s *SubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitSubExpression(s)
	}
}

func (p *SelectStatementParserParser) SubExpression() (localctx ISubExpressionContext) {
	localctx = NewSubExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SelectStatementParserParserRULE_subExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(SelectStatementParserParserLeftParenthesis)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(174)
			p.AndExpression()
		}

	case 2:
		{
			p.SetState(175)
			p.OrExpression()
		}

	}
	{
		p.SetState(178)
		p.Match(SelectStatementParserParserRightParenthesis)
	}

	return localctx
}

// IAtomExpressionContext is an interface to support dynamic dispatch.
type IAtomExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// GetVal returns the val token.
	GetVal() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAtomExpressionContext differentiates from other interfaces.
	IsAtomExpressionContext()
}

type AtomExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	val    antlr.Token
	op     antlr.Token
}

func NewEmptyAtomExpressionContext() *AtomExpressionContext {
	var p = new(AtomExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_atomExpression
	return p
}

func (*AtomExpressionContext) IsAtomExpressionContext() {}

func NewAtomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_atomExpression

	return p
}

func (s *AtomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomExpressionContext) GetField() antlr.Token { return s.field }

func (s *AtomExpressionContext) GetVal() antlr.Token { return s.val }

func (s *AtomExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AtomExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *AtomExpressionContext) SetVal(v antlr.Token) { s.val = v }

func (s *AtomExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AtomExpressionContext) Like() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLike, 0)
}

func (s *AtomExpressionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserBackQuotedString, 0)
}

func (s *AtomExpressionContext) RawString() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRawString, 0)
}

func (s *AtomExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLiteral, 0)
}

func (s *AtomExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, 0)
}

func (s *AtomExpressionContext) Is() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIs, 0)
}

func (s *AtomExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNull, 0)
}

func (s *AtomExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNot, 0)
}

func (s *AtomExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserIn, 0)
}

func (s *AtomExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLeftParenthesis, 0)
}

func (s *AtomExpressionContext) InExpression() IInExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInExpressionContext)
}

func (s *AtomExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserRightParenthesis, 0)
}

func (s *AtomExpressionContext) Equal() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserEqual, 0)
}

func (s *AtomExpressionContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNotEqual, 0)
}

func (s *AtomExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, 0)
}

func (s *AtomExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGT, 0)
}

func (s *AtomExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLT, 0)
}

func (s *AtomExpressionContext) GTE() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserGTE, 0)
}

func (s *AtomExpressionContext) LTE() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLTE, 0)
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterAtomExpression(s)
	}
}

func (s *AtomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitAtomExpression(s)
	}
}

func (p *SelectStatementParserParser) AtomExpression() (localctx IAtomExpressionContext) {
	localctx = NewAtomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SelectStatementParserParserRULE_atomExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(181)
			p.Match(SelectStatementParserParserLike)
		}
		{
			p.SetState(182)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).val = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserQuestionMark || _la == SelectStatementParserParserLiteral) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).val = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(184)
			p.Match(SelectStatementParserParserIs)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SelectStatementParserParserNot {
			{
				p.SetState(185)
				p.Match(SelectStatementParserParserNot)
			}

		}
		{
			p.SetState(188)
			p.Match(SelectStatementParserParserNull)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SelectStatementParserParserNot {
			{
				p.SetState(190)
				p.Match(SelectStatementParserParserNot)
			}

		}
		{
			p.SetState(193)
			p.Match(SelectStatementParserParserIn)
		}
		{
			p.SetState(194)
			p.Match(SelectStatementParserParserLeftParenthesis)
		}
		{
			p.SetState(195)
			p.InExpression()
		}
		{
			p.SetState(196)
			p.Match(SelectStatementParserParserRightParenthesis)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(199)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserEqual || _la == SelectStatementParserParserNotEqual) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(200)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).val = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(SelectStatementParserParserQuestionMark-12))|(1<<(SelectStatementParserParserNumber-12))|(1<<(SelectStatementParserParserLiteral-12)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).val = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserBackQuotedString || _la == SelectStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(202)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SelectStatementParserParserGT)|(1<<SelectStatementParserParserGTE)|(1<<SelectStatementParserParserLT)|(1<<SelectStatementParserParserLTE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(203)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).val = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SelectStatementParserParserQuestionMark || _la == SelectStatementParserParserNumber) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).val = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IInExpressionContext is an interface to support dynamic dispatch.
type IInExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInExpressionContext differentiates from other interfaces.
	IsInExpressionContext()
}

type InExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInExpressionContext() *InExpressionContext {
	var p = new(InExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_inExpression
	return p
}

func (*InExpressionContext) IsInExpressionContext() {}

func NewInExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExpressionContext {
	var p = new(InExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_inExpression

	return p
}

func (s *InExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InExpressionContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserNumber)
}

func (s *InExpressionContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserNumber, i)
}

func (s *InExpressionContext) AllQuestionMark() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserQuestionMark)
}

func (s *InExpressionContext) QuestionMark(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserQuestionMark, i)
}

func (s *InExpressionContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserLiteral)
}

func (s *InExpressionContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserLiteral, i)
}

func (s *InExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(SelectStatementParserParserComma)
}

func (s *InExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserComma, i)
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

func (p *SelectStatementParserParser) InExpression() (localctx IInExpressionContext) {
	localctx = NewInExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SelectStatementParserParserRULE_inExpression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(SelectStatementParserParserQuestionMark-12))|(1<<(SelectStatementParserParserNumber-12))|(1<<(SelectStatementParserParserLiteral-12)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SelectStatementParserParserComma {
		{
			p.SetState(207)
			p.Match(SelectStatementParserParserComma)
		}
		{
			p.SetState(208)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(SelectStatementParserParserQuestionMark-12))|(1<<(SelectStatementParserParserNumber-12))|(1<<(SelectStatementParserParserLiteral-12)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupFunctionContext is an interface to support dynamic dispatch.
type IGroupFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupFunctionContext differentiates from other interfaces.
	IsGroupFunctionContext()
}

type GroupFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupFunctionContext() *GroupFunctionContext {
	var p = new(GroupFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SelectStatementParserParserRULE_groupFunction
	return p
}

func (*GroupFunctionContext) IsGroupFunctionContext() {}

func NewGroupFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupFunctionContext {
	var p = new(GroupFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SelectStatementParserParserRULE_groupFunction

	return p
}

func (s *GroupFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupFunctionContext) Count() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserCount, 0)
}

func (s *GroupFunctionContext) Sum() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserSum, 0)
}

func (s *GroupFunctionContext) Max() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMax, 0)
}

func (s *GroupFunctionContext) Min() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserMin, 0)
}

func (s *GroupFunctionContext) Avg() antlr.TerminalNode {
	return s.GetToken(SelectStatementParserParserAvg, 0)
}

func (s *GroupFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.EnterGroupFunction(s)
	}
}

func (s *GroupFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SelectStatementParserListener); ok {
		listenerT.ExitGroupFunction(s)
	}
}

func (p *SelectStatementParserParser) GroupFunction() (localctx IGroupFunctionContext) {
	localctx = NewGroupFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SelectStatementParserParserRULE_groupFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SelectStatementParserParserCount-33))|(1<<(SelectStatementParserParserSum-33))|(1<<(SelectStatementParserParserMax-33))|(1<<(SelectStatementParserParserMin-33))|(1<<(SelectStatementParserParserAvg-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
