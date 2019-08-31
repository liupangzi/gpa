// Code generated from DeleteStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package deleteparser // DeleteStatementParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 227,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 5, 2, 34, 10, 2,
	3, 2, 5, 2, 37, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3,
	3, 5, 3, 47, 10, 3, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	56, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 66, 10,
	6, 12, 6, 14, 6, 69, 11, 6, 3, 7, 3, 7, 5, 7, 73, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 79, 10, 8, 3, 9, 3, 9, 5, 9, 83, 10, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 88, 10, 9, 7, 9, 90, 10, 9, 12, 9, 14, 9, 93, 11, 9, 3, 10, 3, 10,
	5, 10, 97, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 102, 10, 10, 7, 10, 104,
	10, 10, 12, 10, 14, 10, 107, 11, 10, 3, 11, 3, 11, 3, 11, 5, 11, 112, 10,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 122,
	10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 127, 10, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 140, 10, 12,
	3, 13, 3, 13, 3, 13, 5, 13, 145, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 151, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 158, 10, 15,
	12, 15, 14, 15, 161, 11, 15, 3, 15, 3, 15, 3, 15, 7, 15, 166, 10, 15, 12,
	15, 14, 15, 169, 11, 15, 5, 15, 171, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 225, 10, 16, 3, 16, 2,
	2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 6, 3,
	2, 43, 44, 3, 2, 31, 32, 3, 2, 13, 14, 3, 2, 16, 19, 2, 269, 2, 33, 3,
	2, 2, 2, 4, 40, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 57, 3, 2, 2, 2, 10, 60,
	3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 74, 3, 2, 2, 2, 16, 82, 3, 2, 2, 2,
	18, 96, 3, 2, 2, 2, 20, 108, 3, 2, 2, 2, 22, 139, 3, 2, 2, 2, 24, 144,
	3, 2, 2, 2, 26, 150, 3, 2, 2, 2, 28, 170, 3, 2, 2, 2, 30, 224, 3, 2, 2,
	2, 32, 34, 5, 4, 3, 2, 33, 32, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 36,
	3, 2, 2, 2, 35, 37, 7, 9, 2, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2,
	37, 38, 3, 2, 2, 2, 38, 39, 7, 2, 2, 3, 39, 3, 3, 2, 2, 2, 40, 41, 7, 21,
	2, 2, 41, 43, 5, 6, 4, 2, 42, 44, 5, 8, 5, 2, 43, 42, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 47, 5, 10, 6, 2, 46, 45, 3, 2, 2, 2,
	46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 50, 5, 14, 8, 2, 49, 48, 3,
	2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 22, 2, 2, 52,
	55, 9, 2, 2, 2, 53, 54, 7, 35, 2, 2, 54, 56, 9, 2, 2, 2, 55, 53, 3, 2,
	2, 2, 55, 56, 3, 2, 2, 2, 56, 7, 3, 2, 2, 2, 57, 58, 7, 23, 2, 2, 58, 59,
	5, 16, 9, 2, 59, 9, 3, 2, 2, 2, 60, 61, 7, 28, 2, 2, 61, 62, 7, 30, 2,
	2, 62, 67, 5, 12, 7, 2, 63, 64, 7, 11, 2, 2, 64, 66, 5, 12, 7, 2, 65, 63,
	3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 11, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 72, 9, 2, 2, 2, 71, 73, 9,
	3, 2, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 13, 3, 2, 2, 2, 74,
	78, 7, 34, 2, 2, 75, 79, 5, 30, 16, 2, 76, 79, 7, 15, 2, 2, 77, 79, 7,
	41, 2, 2, 78, 75, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79,
	15, 3, 2, 2, 2, 80, 83, 5, 20, 11, 2, 81, 83, 5, 18, 10, 2, 82, 80, 3,
	2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 91, 3, 2, 2, 2, 84, 87, 7, 27, 2, 2, 85,
	88, 5, 20, 11, 2, 86, 88, 5, 18, 10, 2, 87, 85, 3, 2, 2, 2, 87, 86, 3,
	2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 84, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91,
	89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 17, 3, 2, 2, 2, 93, 91, 3, 2, 2,
	2, 94, 97, 5, 20, 11, 2, 95, 97, 5, 22, 12, 2, 96, 94, 3, 2, 2, 2, 96,
	95, 3, 2, 2, 2, 97, 105, 3, 2, 2, 2, 98, 101, 7, 26, 2, 2, 99, 102, 5,
	20, 11, 2, 100, 102, 5, 22, 12, 2, 101, 99, 3, 2, 2, 2, 101, 100, 3, 2,
	2, 2, 102, 104, 3, 2, 2, 2, 103, 98, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2,
	105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 19, 3, 2, 2, 2, 107, 105,
	3, 2, 2, 2, 108, 111, 7, 7, 2, 2, 109, 112, 5, 18, 10, 2, 110, 112, 5,
	16, 9, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2,
	2, 113, 114, 7, 8, 2, 2, 114, 21, 3, 2, 2, 2, 115, 116, 9, 2, 2, 2, 116,
	117, 7, 25, 2, 2, 117, 140, 5, 24, 13, 2, 118, 119, 9, 2, 2, 2, 119, 121,
	7, 3, 2, 2, 120, 122, 7, 5, 2, 2, 121, 120, 3, 2, 2, 2, 121, 122, 3, 2,
	2, 2, 122, 123, 3, 2, 2, 2, 123, 140, 7, 6, 2, 2, 124, 126, 9, 2, 2, 2,
	125, 127, 7, 5, 2, 2, 126, 125, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 129, 7, 4, 2, 2, 129, 130, 7, 7, 2, 2, 130, 131,
	5, 28, 15, 2, 131, 132, 7, 8, 2, 2, 132, 140, 3, 2, 2, 2, 133, 134, 9,
	2, 2, 2, 134, 135, 9, 4, 2, 2, 135, 140, 5, 26, 14, 2, 136, 137, 9, 2,
	2, 2, 137, 138, 9, 5, 2, 2, 138, 140, 5, 26, 14, 2, 139, 115, 3, 2, 2,
	2, 139, 118, 3, 2, 2, 2, 139, 124, 3, 2, 2, 2, 139, 133, 3, 2, 2, 2, 139,
	136, 3, 2, 2, 2, 140, 23, 3, 2, 2, 2, 141, 145, 5, 30, 16, 2, 142, 145,
	7, 42, 2, 2, 143, 145, 7, 15, 2, 2, 144, 141, 3, 2, 2, 2, 144, 142, 3,
	2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 25, 3, 2, 2, 2, 146, 151, 5, 30, 16,
	2, 147, 151, 7, 41, 2, 2, 148, 151, 7, 42, 2, 2, 149, 151, 7, 15, 2, 2,
	150, 146, 3, 2, 2, 2, 150, 147, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150,
	149, 3, 2, 2, 2, 151, 27, 3, 2, 2, 2, 152, 171, 5, 30, 16, 2, 153, 171,
	7, 15, 2, 2, 154, 159, 7, 41, 2, 2, 155, 156, 7, 11, 2, 2, 156, 158, 7,
	41, 2, 2, 157, 155, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2,
	2, 159, 160, 3, 2, 2, 2, 160, 171, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162,
	167, 7, 42, 2, 2, 163, 164, 7, 11, 2, 2, 164, 166, 7, 42, 2, 2, 165, 163,
	3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2,
	2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 152, 3, 2, 2, 2,
	170, 153, 3, 2, 2, 2, 170, 154, 3, 2, 2, 2, 170, 162, 3, 2, 2, 2, 171,
	29, 3, 2, 2, 2, 172, 173, 7, 10, 2, 2, 173, 225, 7, 44, 2, 2, 174, 175,
	7, 10, 2, 2, 175, 225, 7, 3, 2, 2, 176, 177, 7, 10, 2, 2, 177, 225, 7,
	4, 2, 2, 178, 179, 7, 10, 2, 2, 179, 225, 7, 5, 2, 2, 180, 181, 7, 10,
	2, 2, 181, 225, 7, 6, 2, 2, 182, 183, 7, 10, 2, 2, 183, 225, 7, 20, 2,
	2, 184, 185, 7, 10, 2, 2, 185, 225, 7, 21, 2, 2, 186, 187, 7, 10, 2, 2,
	187, 225, 7, 22, 2, 2, 188, 189, 7, 10, 2, 2, 189, 225, 7, 23, 2, 2, 190,
	191, 7, 10, 2, 2, 191, 225, 7, 24, 2, 2, 192, 193, 7, 10, 2, 2, 193, 225,
	7, 25, 2, 2, 194, 195, 7, 10, 2, 2, 195, 225, 7, 26, 2, 2, 196, 197, 7,
	10, 2, 2, 197, 225, 7, 27, 2, 2, 198, 199, 7, 10, 2, 2, 199, 225, 7, 28,
	2, 2, 200, 201, 7, 10, 2, 2, 201, 225, 7, 29, 2, 2, 202, 203, 7, 10, 2,
	2, 203, 225, 7, 30, 2, 2, 204, 205, 7, 10, 2, 2, 205, 225, 7, 31, 2, 2,
	206, 207, 7, 10, 2, 2, 207, 225, 7, 32, 2, 2, 208, 209, 7, 10, 2, 2, 209,
	225, 7, 33, 2, 2, 210, 211, 7, 10, 2, 2, 211, 225, 7, 34, 2, 2, 212, 213,
	7, 10, 2, 2, 213, 225, 7, 35, 2, 2, 214, 215, 7, 10, 2, 2, 215, 225, 7,
	36, 2, 2, 216, 217, 7, 10, 2, 2, 217, 225, 7, 37, 2, 2, 218, 219, 7, 10,
	2, 2, 219, 225, 7, 38, 2, 2, 220, 221, 7, 10, 2, 2, 221, 225, 7, 39, 2,
	2, 222, 223, 7, 10, 2, 2, 223, 225, 7, 40, 2, 2, 224, 172, 3, 2, 2, 2,
	224, 174, 3, 2, 2, 2, 224, 176, 3, 2, 2, 2, 224, 178, 3, 2, 2, 2, 224,
	180, 3, 2, 2, 2, 224, 182, 3, 2, 2, 2, 224, 184, 3, 2, 2, 2, 224, 186,
	3, 2, 2, 2, 224, 188, 3, 2, 2, 2, 224, 190, 3, 2, 2, 2, 224, 192, 3, 2,
	2, 2, 224, 194, 3, 2, 2, 2, 224, 196, 3, 2, 2, 2, 224, 198, 3, 2, 2, 2,
	224, 200, 3, 2, 2, 2, 224, 202, 3, 2, 2, 2, 224, 204, 3, 2, 2, 2, 224,
	206, 3, 2, 2, 2, 224, 208, 3, 2, 2, 2, 224, 210, 3, 2, 2, 2, 224, 212,
	3, 2, 2, 2, 224, 214, 3, 2, 2, 2, 224, 216, 3, 2, 2, 2, 224, 218, 3, 2,
	2, 2, 224, 220, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 31, 3, 2, 2, 2,
	27, 33, 36, 43, 46, 49, 55, 67, 72, 78, 82, 87, 91, 96, 101, 105, 111,
	121, 126, 139, 144, 150, 159, 167, 170, 224,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'('", "')'", "';'", "':'", "','", "'*'", "'='", "'!='",
	"'?'", "'>'", "'>='", "'<'", "'<='",
}
var symbolicNames = []string{
	"", "Is", "In", "Not", "Null", "LeftParenthesis", "RightParenthesis", "Semicolon",
	"Colon", "Comma", "Asterisk", "Equal", "NotEqual", "QuestionMark", "GT",
	"GTE", "LT", "LTE", "Select", "Delete", "From", "Where", "Having", "Like",
	"And", "Or", "Order", "Group", "By", "Asc", "Desc", "Offset", "Limit",
	"As", "Count", "Sum", "Max", "Min", "Avg", "Number", "Literal", "BackQuotedString",
	"RawString", "WS",
}

var ruleNames = []string{
	"statement", "deleteStatement", "fromClause", "whereClause", "orderByClause",
	"orderByField", "limitClause", "orExpression", "andExpression", "subExpression",
	"atomExpression", "likeExpression", "compareExpression", "inExpression",
	"colonField",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DeleteStatementParserParser struct {
	*antlr.BaseParser
}

func NewDeleteStatementParserParser(input antlr.TokenStream) *DeleteStatementParserParser {
	this := new(DeleteStatementParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DeleteStatementParser.g4"

	return this
}

// DeleteStatementParserParser tokens.
const (
	DeleteStatementParserParserEOF              = antlr.TokenEOF
	DeleteStatementParserParserIs               = 1
	DeleteStatementParserParserIn               = 2
	DeleteStatementParserParserNot              = 3
	DeleteStatementParserParserNull             = 4
	DeleteStatementParserParserLeftParenthesis  = 5
	DeleteStatementParserParserRightParenthesis = 6
	DeleteStatementParserParserSemicolon        = 7
	DeleteStatementParserParserColon            = 8
	DeleteStatementParserParserComma            = 9
	DeleteStatementParserParserAsterisk         = 10
	DeleteStatementParserParserEqual            = 11
	DeleteStatementParserParserNotEqual         = 12
	DeleteStatementParserParserQuestionMark     = 13
	DeleteStatementParserParserGT               = 14
	DeleteStatementParserParserGTE              = 15
	DeleteStatementParserParserLT               = 16
	DeleteStatementParserParserLTE              = 17
	DeleteStatementParserParserSelect           = 18
	DeleteStatementParserParserDelete           = 19
	DeleteStatementParserParserFrom             = 20
	DeleteStatementParserParserWhere            = 21
	DeleteStatementParserParserHaving           = 22
	DeleteStatementParserParserLike             = 23
	DeleteStatementParserParserAnd              = 24
	DeleteStatementParserParserOr               = 25
	DeleteStatementParserParserOrder            = 26
	DeleteStatementParserParserGroup            = 27
	DeleteStatementParserParserBy               = 28
	DeleteStatementParserParserAsc              = 29
	DeleteStatementParserParserDesc             = 30
	DeleteStatementParserParserOffset           = 31
	DeleteStatementParserParserLimit            = 32
	DeleteStatementParserParserAs               = 33
	DeleteStatementParserParserCount            = 34
	DeleteStatementParserParserSum              = 35
	DeleteStatementParserParserMax              = 36
	DeleteStatementParserParserMin              = 37
	DeleteStatementParserParserAvg              = 38
	DeleteStatementParserParserNumber           = 39
	DeleteStatementParserParserLiteral          = 40
	DeleteStatementParserParserBackQuotedString = 41
	DeleteStatementParserParserRawString        = 42
	DeleteStatementParserParserWS               = 43
)

// DeleteStatementParserParser rules.
const (
	DeleteStatementParserParserRULE_statement         = 0
	DeleteStatementParserParserRULE_deleteStatement   = 1
	DeleteStatementParserParserRULE_fromClause        = 2
	DeleteStatementParserParserRULE_whereClause       = 3
	DeleteStatementParserParserRULE_orderByClause     = 4
	DeleteStatementParserParserRULE_orderByField      = 5
	DeleteStatementParserParserRULE_limitClause       = 6
	DeleteStatementParserParserRULE_orExpression      = 7
	DeleteStatementParserParserRULE_andExpression     = 8
	DeleteStatementParserParserRULE_subExpression     = 9
	DeleteStatementParserParserRULE_atomExpression    = 10
	DeleteStatementParserParserRULE_likeExpression    = 11
	DeleteStatementParserParserRULE_compareExpression = 12
	DeleteStatementParserParserRULE_inExpression      = 13
	DeleteStatementParserParserRULE_colonField        = 14
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
	p.RuleIndex = DeleteStatementParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserEOF, 0)
}

func (s *StatementContext) DeleteStatement() IDeleteStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *StatementContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserSemicolon, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DeleteStatementParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DeleteStatementParserParserRULE_statement)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserDelete {
		{
			p.SetState(30)
			p.DeleteStatement()
		}

	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserSemicolon {
		{
			p.SetState(33)
			p.Match(DeleteStatementParserParserSemicolon)
		}

	}
	{
		p.SetState(36)
		p.Match(DeleteStatementParserParserEOF)
	}

	return localctx
}

// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_deleteStatement
	return p
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) Delete() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserDelete, 0)
}

func (s *DeleteStatementContext) FromClause() IFromClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *DeleteStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *DeleteStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *DeleteStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (p *DeleteStatementParserParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DeleteStatementParserParserRULE_deleteStatement)
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
		p.SetState(38)
		p.Match(DeleteStatementParserParserDelete)
	}
	{
		p.SetState(39)
		p.FromClause()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserWhere {
		{
			p.SetState(40)
			p.WhereClause()
		}

	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserOrder {
		{
			p.SetState(43)
			p.OrderByClause()
		}

	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserLimit {
		{
			p.SetState(46)
			p.LimitClause()
		}

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

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetTable sets the table token.
	SetTable(antlr.Token)

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	table  antlr.Token
	alias  antlr.Token
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) GetTable() antlr.Token { return s.table }

func (s *FromClauseContext) GetAlias() antlr.Token { return s.alias }

func (s *FromClauseContext) SetTable(v antlr.Token) { s.table = v }

func (s *FromClauseContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *FromClauseContext) From() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserFrom, 0)
}

func (s *FromClauseContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(DeleteStatementParserParserBackQuotedString)
}

func (s *FromClauseContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserBackQuotedString, i)
}

func (s *FromClauseContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(DeleteStatementParserParserRawString)
}

func (s *FromClauseContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRawString, i)
}

func (s *FromClauseContext) As() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAs, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (p *DeleteStatementParserParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DeleteStatementParserParserRULE_fromClause)
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
		p.SetState(49)
		p.Match(DeleteStatementParserParserFrom)
	}
	{
		p.SetState(50)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*FromClauseContext).table = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*FromClauseContext).table = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserAs {
		{
			p.SetState(51)
			p.Match(DeleteStatementParserParserAs)
		}
		{
			p.SetState(52)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FromClauseContext).alias = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FromClauseContext).alias = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
	p.RuleIndex = DeleteStatementParserParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) Where() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserWhere, 0)
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
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *DeleteStatementParserParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DeleteStatementParserParserRULE_whereClause)

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
		p.SetState(55)
		p.Match(DeleteStatementParserParserWhere)
	}
	{
		p.SetState(56)
		p.OrExpression()
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
	p.RuleIndex = DeleteStatementParserParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) Order() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserOrder, 0)
}

func (s *OrderByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserBy, 0)
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
	return s.GetTokens(DeleteStatementParserParserComma)
}

func (s *OrderByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserComma, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (p *DeleteStatementParserParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DeleteStatementParserParserRULE_orderByClause)
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
		p.SetState(58)
		p.Match(DeleteStatementParserParserOrder)
	}
	{
		p.SetState(59)
		p.Match(DeleteStatementParserParserBy)
	}
	{
		p.SetState(60)
		p.OrderByField()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DeleteStatementParserParserComma {
		{
			p.SetState(61)
			p.Match(DeleteStatementParserParserComma)
		}
		{
			p.SetState(62)
			p.OrderByField()
		}

		p.SetState(67)
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
	p.RuleIndex = DeleteStatementParserParserRULE_orderByField
	return p
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_orderByField

	return p
}

func (s *OrderByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByFieldContext) GetField() antlr.Token { return s.field }

func (s *OrderByFieldContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *OrderByFieldContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserBackQuotedString, 0)
}

func (s *OrderByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRawString, 0)
}

func (s *OrderByFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAsc, 0)
}

func (s *OrderByFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserDesc, 0)
}

func (s *OrderByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterOrderByField(s)
	}
}

func (s *OrderByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitOrderByField(s)
	}
}

func (p *DeleteStatementParserParser) OrderByField() (localctx IOrderByFieldContext) {
	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DeleteStatementParserParserRULE_orderByField)
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
		p.SetState(68)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OrderByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OrderByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DeleteStatementParserParserAsc || _la == DeleteStatementParserParserDesc {
		{
			p.SetState(69)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByFieldContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserAsc || _la == DeleteStatementParserParserDesc) {
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
	p.RuleIndex = DeleteStatementParserParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) Limit() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLimit, 0)
}

func (s *LimitClauseContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *LimitClauseContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserQuestionMark, 0)
}

func (s *LimitClauseContext) Number() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNumber, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *DeleteStatementParserParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DeleteStatementParserParserRULE_limitClause)

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
		p.SetState(72)
		p.Match(DeleteStatementParserParserLimit)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DeleteStatementParserParserColon:
		{
			p.SetState(73)
			p.ColonField()
		}

	case DeleteStatementParserParserQuestionMark:
		{
			p.SetState(74)
			p.Match(DeleteStatementParserParserQuestionMark)
		}

	case DeleteStatementParserParserNumber:
		{
			p.SetState(75)
			p.Match(DeleteStatementParserParserNumber)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = DeleteStatementParserParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_orExpression

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
	return s.GetTokens(DeleteStatementParserParserOr)
}

func (s *OrExpressionContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserOr, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *DeleteStatementParserParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DeleteStatementParserParserRULE_orExpression)
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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(78)
			p.SubExpression()
		}

	case 2:
		{
			p.SetState(79)
			p.AndExpression()
		}

	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DeleteStatementParserParserOr {
		{
			p.SetState(82)
			p.Match(DeleteStatementParserParserOr)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(83)
				p.SubExpression()
			}

		case 2:
			{
				p.SetState(84)
				p.AndExpression()
			}

		}

		p.SetState(91)
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
	p.RuleIndex = DeleteStatementParserParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_andExpression

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
	return s.GetTokens(DeleteStatementParserParserAnd)
}

func (s *AndExpressionContext) And(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAnd, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *DeleteStatementParserParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DeleteStatementParserParserRULE_andExpression)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DeleteStatementParserParserLeftParenthesis:
		{
			p.SetState(92)
			p.SubExpression()
		}

	case DeleteStatementParserParserBackQuotedString, DeleteStatementParserParserRawString:
		{
			p.SetState(93)
			p.AtomExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DeleteStatementParserParserAnd {
		{
			p.SetState(96)
			p.Match(DeleteStatementParserParserAnd)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DeleteStatementParserParserLeftParenthesis:
			{
				p.SetState(97)
				p.SubExpression()
			}

		case DeleteStatementParserParserBackQuotedString, DeleteStatementParserParserRawString:
			{
				p.SetState(98)
				p.AtomExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(105)
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
	p.RuleIndex = DeleteStatementParserParserRULE_subExpression
	return p
}

func (*SubExpressionContext) IsSubExpressionContext() {}

func NewSubExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubExpressionContext {
	var p = new(SubExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_subExpression

	return p
}

func (s *SubExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLeftParenthesis, 0)
}

func (s *SubExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRightParenthesis, 0)
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
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterSubExpression(s)
	}
}

func (s *SubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitSubExpression(s)
	}
}

func (p *DeleteStatementParserParser) SubExpression() (localctx ISubExpressionContext) {
	localctx = NewSubExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DeleteStatementParserParserRULE_subExpression)

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
		p.SetState(106)
		p.Match(DeleteStatementParserParserLeftParenthesis)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(107)
			p.AndExpression()
		}

	case 2:
		{
			p.SetState(108)
			p.OrExpression()
		}

	}
	{
		p.SetState(111)
		p.Match(DeleteStatementParserParserRightParenthesis)
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

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAtomExpressionContext differentiates from other interfaces.
	IsAtomExpressionContext()
}

type AtomExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
	op     antlr.Token
}

func NewEmptyAtomExpressionContext() *AtomExpressionContext {
	var p = new(AtomExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_atomExpression
	return p
}

func (*AtomExpressionContext) IsAtomExpressionContext() {}

func NewAtomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_atomExpression

	return p
}

func (s *AtomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomExpressionContext) GetField() antlr.Token { return s.field }

func (s *AtomExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AtomExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *AtomExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AtomExpressionContext) Like() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLike, 0)
}

func (s *AtomExpressionContext) LikeExpression() ILikeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeExpressionContext)
}

func (s *AtomExpressionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserBackQuotedString, 0)
}

func (s *AtomExpressionContext) RawString() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRawString, 0)
}

func (s *AtomExpressionContext) Is() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserIs, 0)
}

func (s *AtomExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNull, 0)
}

func (s *AtomExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNot, 0)
}

func (s *AtomExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserIn, 0)
}

func (s *AtomExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLeftParenthesis, 0)
}

func (s *AtomExpressionContext) InExpression() IInExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInExpressionContext)
}

func (s *AtomExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRightParenthesis, 0)
}

func (s *AtomExpressionContext) CompareExpression() ICompareExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareExpressionContext)
}

func (s *AtomExpressionContext) Equal() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserEqual, 0)
}

func (s *AtomExpressionContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNotEqual, 0)
}

func (s *AtomExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserGT, 0)
}

func (s *AtomExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLT, 0)
}

func (s *AtomExpressionContext) GTE() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserGTE, 0)
}

func (s *AtomExpressionContext) LTE() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLTE, 0)
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterAtomExpression(s)
	}
}

func (s *AtomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitAtomExpression(s)
	}
}

func (p *DeleteStatementParserParser) AtomExpression() (localctx IAtomExpressionContext) {
	localctx = NewAtomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DeleteStatementParserParserRULE_atomExpression)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(114)
			p.Match(DeleteStatementParserParserLike)
		}
		{
			p.SetState(115)
			p.LikeExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(117)
			p.Match(DeleteStatementParserParserIs)
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DeleteStatementParserParserNot {
			{
				p.SetState(118)
				p.Match(DeleteStatementParserParserNot)
			}

		}
		{
			p.SetState(121)
			p.Match(DeleteStatementParserParserNull)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DeleteStatementParserParserNot {
			{
				p.SetState(123)
				p.Match(DeleteStatementParserParserNot)
			}

		}
		{
			p.SetState(126)
			p.Match(DeleteStatementParserParserIn)
		}
		{
			p.SetState(127)
			p.Match(DeleteStatementParserParserLeftParenthesis)
		}
		{
			p.SetState(128)
			p.InExpression()
		}
		{
			p.SetState(129)
			p.Match(DeleteStatementParserParserRightParenthesis)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(132)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserEqual || _la == DeleteStatementParserParserNotEqual) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(133)
			p.CompareExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(134)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DeleteStatementParserParserBackQuotedString || _la == DeleteStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(135)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DeleteStatementParserParserGT)|(1<<DeleteStatementParserParserGTE)|(1<<DeleteStatementParserParserLT)|(1<<DeleteStatementParserParserLTE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(136)
			p.CompareExpression()
		}

	}

	return localctx
}

// ILikeExpressionContext is an interface to support dynamic dispatch.
type ILikeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeExpressionContext differentiates from other interfaces.
	IsLikeExpressionContext()
}

type LikeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeExpressionContext() *LikeExpressionContext {
	var p = new(LikeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_likeExpression
	return p
}

func (*LikeExpressionContext) IsLikeExpressionContext() {}

func NewLikeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeExpressionContext {
	var p = new(LikeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_likeExpression

	return p
}

func (s *LikeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *LikeExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLiteral, 0)
}

func (s *LikeExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserQuestionMark, 0)
}

func (s *LikeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterLikeExpression(s)
	}
}

func (s *LikeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitLikeExpression(s)
	}
}

func (p *DeleteStatementParserParser) LikeExpression() (localctx ILikeExpressionContext) {
	localctx = NewLikeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DeleteStatementParserParserRULE_likeExpression)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DeleteStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.ColonField()
		}

	case DeleteStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(DeleteStatementParserParserLiteral)
		}

	case DeleteStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Match(DeleteStatementParserParserQuestionMark)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompareExpressionContext is an interface to support dynamic dispatch.
type ICompareExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareExpressionContext differentiates from other interfaces.
	IsCompareExpressionContext()
}

type CompareExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareExpressionContext() *CompareExpressionContext {
	var p = new(CompareExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_compareExpression
	return p
}

func (*CompareExpressionContext) IsCompareExpressionContext() {}

func NewCompareExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_compareExpression

	return p
}

func (s *CompareExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *CompareExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNumber, 0)
}

func (s *CompareExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLiteral, 0)
}

func (s *CompareExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserQuestionMark, 0)
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

func (p *DeleteStatementParserParser) CompareExpression() (localctx ICompareExpressionContext) {
	localctx = NewCompareExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DeleteStatementParserParserRULE_compareExpression)

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

	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DeleteStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.ColonField()
		}

	case DeleteStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(DeleteStatementParserParserNumber)
		}

	case DeleteStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(DeleteStatementParserParserLiteral)
		}

	case DeleteStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.Match(DeleteStatementParserParserQuestionMark)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = DeleteStatementParserParserRULE_inExpression
	return p
}

func (*InExpressionContext) IsInExpressionContext() {}

func NewInExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExpressionContext {
	var p = new(InExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_inExpression

	return p
}

func (s *InExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InExpressionContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *InExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserQuestionMark, 0)
}

func (s *InExpressionContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(DeleteStatementParserParserNumber)
}

func (s *InExpressionContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNumber, i)
}

func (s *InExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DeleteStatementParserParserComma)
}

func (s *InExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserComma, i)
}

func (s *InExpressionContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(DeleteStatementParserParserLiteral)
}

func (s *InExpressionContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLiteral, i)
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

func (p *DeleteStatementParserParser) InExpression() (localctx IInExpressionContext) {
	localctx = NewInExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DeleteStatementParserParserRULE_inExpression)
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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DeleteStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.ColonField()
		}

	case DeleteStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Match(DeleteStatementParserParserQuestionMark)
		}

	case DeleteStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(DeleteStatementParserParserNumber)
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DeleteStatementParserParserComma {
			{
				p.SetState(153)
				p.Match(DeleteStatementParserParserComma)
			}
			{
				p.SetState(154)
				p.Match(DeleteStatementParserParserNumber)
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DeleteStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.Match(DeleteStatementParserParserLiteral)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DeleteStatementParserParserComma {
			{
				p.SetState(161)
				p.Match(DeleteStatementParserParserComma)
			}
			{
				p.SetState(162)
				p.Match(DeleteStatementParserParserLiteral)
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColonFieldContext is an interface to support dynamic dispatch.
type IColonFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColonFieldContext differentiates from other interfaces.
	IsColonFieldContext()
}

type ColonFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColonFieldContext() *ColonFieldContext {
	var p = new(ColonFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeleteStatementParserParserRULE_colonField
	return p
}

func (*ColonFieldContext) IsColonFieldContext() {}

func NewColonFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColonFieldContext {
	var p = new(ColonFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeleteStatementParserParserRULE_colonField

	return p
}

func (s *ColonFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ColonFieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserColon, 0)
}

func (s *ColonFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserRawString, 0)
}

func (s *ColonFieldContext) Is() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserIs, 0)
}

func (s *ColonFieldContext) In() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserIn, 0)
}

func (s *ColonFieldContext) Not() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNot, 0)
}

func (s *ColonFieldContext) Null() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserNull, 0)
}

func (s *ColonFieldContext) Select() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserSelect, 0)
}

func (s *ColonFieldContext) Delete() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserDelete, 0)
}

func (s *ColonFieldContext) From() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserFrom, 0)
}

func (s *ColonFieldContext) Where() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserWhere, 0)
}

func (s *ColonFieldContext) Having() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserHaving, 0)
}

func (s *ColonFieldContext) Like() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLike, 0)
}

func (s *ColonFieldContext) And() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAnd, 0)
}

func (s *ColonFieldContext) Or() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserOr, 0)
}

func (s *ColonFieldContext) Order() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserOrder, 0)
}

func (s *ColonFieldContext) Group() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserGroup, 0)
}

func (s *ColonFieldContext) By() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserBy, 0)
}

func (s *ColonFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAsc, 0)
}

func (s *ColonFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserDesc, 0)
}

func (s *ColonFieldContext) Offset() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserOffset, 0)
}

func (s *ColonFieldContext) Limit() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserLimit, 0)
}

func (s *ColonFieldContext) As() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAs, 0)
}

func (s *ColonFieldContext) Count() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserCount, 0)
}

func (s *ColonFieldContext) Sum() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserSum, 0)
}

func (s *ColonFieldContext) Max() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserMax, 0)
}

func (s *ColonFieldContext) Min() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserMin, 0)
}

func (s *ColonFieldContext) Avg() antlr.TerminalNode {
	return s.GetToken(DeleteStatementParserParserAvg, 0)
}

func (s *ColonFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColonFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColonFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.EnterColonField(s)
	}
}

func (s *ColonFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeleteStatementParserListener); ok {
		listenerT.ExitColonField(s)
	}
}

func (p *DeleteStatementParserParser) ColonField() (localctx IColonFieldContext) {
	localctx = NewColonFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DeleteStatementParserParserRULE_colonField)

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

	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(171)
			p.Match(DeleteStatementParserParserRawString)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(173)
			p.Match(DeleteStatementParserParserIs)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(175)
			p.Match(DeleteStatementParserParserIn)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(177)
			p.Match(DeleteStatementParserParserNot)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(178)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(179)
			p.Match(DeleteStatementParserParserNull)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(181)
			p.Match(DeleteStatementParserParserSelect)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(182)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(183)
			p.Match(DeleteStatementParserParserDelete)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(184)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(185)
			p.Match(DeleteStatementParserParserFrom)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(186)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(187)
			p.Match(DeleteStatementParserParserWhere)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(188)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(189)
			p.Match(DeleteStatementParserParserHaving)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(190)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(191)
			p.Match(DeleteStatementParserParserLike)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(192)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(193)
			p.Match(DeleteStatementParserParserAnd)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(194)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(195)
			p.Match(DeleteStatementParserParserOr)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(196)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(197)
			p.Match(DeleteStatementParserParserOrder)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(198)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(199)
			p.Match(DeleteStatementParserParserGroup)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(200)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(201)
			p.Match(DeleteStatementParserParserBy)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(202)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(203)
			p.Match(DeleteStatementParserParserAsc)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(204)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(205)
			p.Match(DeleteStatementParserParserDesc)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(206)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(207)
			p.Match(DeleteStatementParserParserOffset)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(208)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(209)
			p.Match(DeleteStatementParserParserLimit)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(210)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(211)
			p.Match(DeleteStatementParserParserAs)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(212)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(213)
			p.Match(DeleteStatementParserParserCount)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(214)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(215)
			p.Match(DeleteStatementParserParserSum)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(216)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(217)
			p.Match(DeleteStatementParserParserMax)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(218)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(219)
			p.Match(DeleteStatementParserParserMin)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(220)
			p.Match(DeleteStatementParserParserColon)
		}
		{
			p.SetState(221)
			p.Match(DeleteStatementParserParserAvg)
		}

	}

	return localctx
}
