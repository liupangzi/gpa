// Code generated from UpdateStatementParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package updateparser // UpdateStatementParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 241,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 5,
	2, 36, 10, 2, 3, 2, 5, 2, 39, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 48, 10, 3, 3, 3, 5, 3, 51, 10, 3, 3, 3, 5, 3, 54, 10, 3, 3,
	4, 3, 4, 3, 4, 7, 4, 59, 10, 4, 12, 4, 14, 4, 62, 11, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 70, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 80, 10, 7, 12, 7, 14, 7, 83, 11, 7, 3, 8, 3, 8, 5,
	8, 87, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 93, 10, 9, 3, 10, 3, 10, 5,
	10, 97, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10, 102, 10, 10, 7, 10, 104, 10,
	10, 12, 10, 14, 10, 107, 11, 10, 3, 11, 3, 11, 5, 11, 111, 10, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 116, 10, 11, 7, 11, 118, 10, 11, 12, 11, 14, 11, 121,
	11, 11, 3, 12, 3, 12, 3, 12, 5, 12, 126, 10, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 136, 10, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 141, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 154, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14,
	159, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 165, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 7, 16, 172, 10, 16, 12, 16, 14, 16, 175, 11, 16,
	3, 16, 3, 16, 3, 16, 7, 16, 180, 10, 16, 12, 16, 14, 16, 183, 11, 16, 5,
	16, 185, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 239, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 6, 3, 2, 45, 46, 3, 2, 33, 34,
	3, 2, 15, 16, 3, 2, 18, 21, 2, 285, 2, 35, 3, 2, 2, 2, 4, 42, 3, 2, 2,
	2, 6, 55, 3, 2, 2, 2, 8, 63, 3, 2, 2, 2, 10, 71, 3, 2, 2, 2, 12, 74, 3,
	2, 2, 2, 14, 84, 3, 2, 2, 2, 16, 88, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20,
	110, 3, 2, 2, 2, 22, 122, 3, 2, 2, 2, 24, 153, 3, 2, 2, 2, 26, 158, 3,
	2, 2, 2, 28, 164, 3, 2, 2, 2, 30, 184, 3, 2, 2, 2, 32, 238, 3, 2, 2, 2,
	34, 36, 5, 4, 3, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3,
	2, 2, 2, 37, 39, 7, 11, 2, 2, 38, 37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39,
	40, 3, 2, 2, 2, 40, 41, 7, 2, 2, 3, 41, 3, 3, 2, 2, 2, 42, 43, 7, 3, 2,
	2, 43, 44, 9, 2, 2, 2, 44, 45, 7, 8, 2, 2, 45, 47, 5, 6, 4, 2, 46, 48,
	5, 10, 6, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 50, 3, 2, 2, 2,
	49, 51, 5, 12, 7, 2, 50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 53, 3,
	2, 2, 2, 52, 54, 5, 16, 9, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	5, 3, 2, 2, 2, 55, 60, 5, 8, 5, 2, 56, 57, 7, 13, 2, 2, 57, 59, 5, 8, 5,
	2, 58, 56, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 7, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 64, 9, 2, 2, 2,
	64, 69, 7, 15, 2, 2, 65, 70, 7, 17, 2, 2, 66, 70, 5, 32, 17, 2, 67, 70,
	7, 44, 2, 2, 68, 70, 7, 43, 2, 2, 69, 65, 3, 2, 2, 2, 69, 66, 3, 2, 2,
	2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 9, 3, 2, 2, 2, 71, 72, 7,
	25, 2, 2, 72, 73, 5, 18, 10, 2, 73, 11, 3, 2, 2, 2, 74, 75, 7, 30, 2, 2,
	75, 76, 7, 32, 2, 2, 76, 81, 5, 14, 8, 2, 77, 78, 7, 13, 2, 2, 78, 80,
	5, 14, 8, 2, 79, 77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2,
	81, 82, 3, 2, 2, 2, 82, 13, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 86, 9,
	2, 2, 2, 85, 87, 9, 3, 2, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87,
	15, 3, 2, 2, 2, 88, 92, 7, 36, 2, 2, 89, 93, 5, 32, 17, 2, 90, 93, 7, 17,
	2, 2, 91, 93, 7, 43, 2, 2, 92, 89, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92,
	91, 3, 2, 2, 2, 93, 17, 3, 2, 2, 2, 94, 97, 5, 22, 12, 2, 95, 97, 5, 20,
	11, 2, 96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 105, 3, 2, 2, 2, 98,
	101, 7, 29, 2, 2, 99, 102, 5, 22, 12, 2, 100, 102, 5, 20, 11, 2, 101, 99,
	3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 98, 3, 2,
	2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2,
	106, 19, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 111, 5, 22, 12, 2, 109,
	111, 5, 24, 13, 2, 110, 108, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 119,
	3, 2, 2, 2, 112, 115, 7, 28, 2, 2, 113, 116, 5, 22, 12, 2, 114, 116, 5,
	24, 13, 2, 115, 113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 118, 3, 2,
	2, 2, 117, 112, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2,
	119, 120, 3, 2, 2, 2, 120, 21, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 125,
	7, 9, 2, 2, 123, 126, 5, 20, 11, 2, 124, 126, 5, 18, 10, 2, 125, 123, 3,
	2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 10, 2,
	2, 128, 23, 3, 2, 2, 2, 129, 130, 9, 2, 2, 2, 130, 131, 7, 27, 2, 2, 131,
	154, 5, 26, 14, 2, 132, 133, 9, 2, 2, 2, 133, 135, 7, 4, 2, 2, 134, 136,
	7, 6, 2, 2, 135, 134, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 3, 2,
	2, 2, 137, 154, 7, 7, 2, 2, 138, 140, 9, 2, 2, 2, 139, 141, 7, 6, 2, 2,
	140, 139, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142,
	143, 7, 5, 2, 2, 143, 144, 7, 9, 2, 2, 144, 145, 5, 30, 16, 2, 145, 146,
	7, 10, 2, 2, 146, 154, 3, 2, 2, 2, 147, 148, 9, 2, 2, 2, 148, 149, 9, 4,
	2, 2, 149, 154, 5, 28, 15, 2, 150, 151, 9, 2, 2, 2, 151, 152, 9, 5, 2,
	2, 152, 154, 5, 28, 15, 2, 153, 129, 3, 2, 2, 2, 153, 132, 3, 2, 2, 2,
	153, 138, 3, 2, 2, 2, 153, 147, 3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 154,
	25, 3, 2, 2, 2, 155, 159, 5, 32, 17, 2, 156, 159, 7, 44, 2, 2, 157, 159,
	7, 17, 2, 2, 158, 155, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2,
	2, 2, 159, 27, 3, 2, 2, 2, 160, 165, 5, 32, 17, 2, 161, 165, 7, 43, 2,
	2, 162, 165, 7, 44, 2, 2, 163, 165, 7, 17, 2, 2, 164, 160, 3, 2, 2, 2,
	164, 161, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 165,
	29, 3, 2, 2, 2, 166, 185, 5, 32, 17, 2, 167, 185, 7, 17, 2, 2, 168, 173,
	7, 43, 2, 2, 169, 170, 7, 13, 2, 2, 170, 172, 7, 43, 2, 2, 171, 169, 3,
	2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2,
	2, 174, 185, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 181, 7, 44, 2, 2, 177,
	178, 7, 13, 2, 2, 178, 180, 7, 44, 2, 2, 179, 177, 3, 2, 2, 2, 180, 183,
	3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 185, 3, 2,
	2, 2, 183, 181, 3, 2, 2, 2, 184, 166, 3, 2, 2, 2, 184, 167, 3, 2, 2, 2,
	184, 168, 3, 2, 2, 2, 184, 176, 3, 2, 2, 2, 185, 31, 3, 2, 2, 2, 186, 187,
	7, 12, 2, 2, 187, 239, 7, 46, 2, 2, 188, 189, 7, 12, 2, 2, 189, 239, 7,
	4, 2, 2, 190, 191, 7, 12, 2, 2, 191, 239, 7, 5, 2, 2, 192, 193, 7, 12,
	2, 2, 193, 239, 7, 6, 2, 2, 194, 195, 7, 12, 2, 2, 195, 239, 7, 7, 2, 2,
	196, 197, 7, 12, 2, 2, 197, 239, 7, 22, 2, 2, 198, 199, 7, 12, 2, 2, 199,
	239, 7, 23, 2, 2, 200, 201, 7, 12, 2, 2, 201, 239, 7, 24, 2, 2, 202, 203,
	7, 12, 2, 2, 203, 239, 7, 25, 2, 2, 204, 205, 7, 12, 2, 2, 205, 239, 7,
	26, 2, 2, 206, 207, 7, 12, 2, 2, 207, 239, 7, 27, 2, 2, 208, 209, 7, 12,
	2, 2, 209, 239, 7, 28, 2, 2, 210, 211, 7, 12, 2, 2, 211, 239, 7, 29, 2,
	2, 212, 213, 7, 12, 2, 2, 213, 239, 7, 30, 2, 2, 214, 215, 7, 12, 2, 2,
	215, 239, 7, 31, 2, 2, 216, 217, 7, 12, 2, 2, 217, 239, 7, 32, 2, 2, 218,
	219, 7, 12, 2, 2, 219, 239, 7, 33, 2, 2, 220, 221, 7, 12, 2, 2, 221, 239,
	7, 34, 2, 2, 222, 223, 7, 12, 2, 2, 223, 239, 7, 35, 2, 2, 224, 225, 7,
	12, 2, 2, 225, 239, 7, 36, 2, 2, 226, 227, 7, 12, 2, 2, 227, 239, 7, 37,
	2, 2, 228, 229, 7, 12, 2, 2, 229, 239, 7, 38, 2, 2, 230, 231, 7, 12, 2,
	2, 231, 239, 7, 39, 2, 2, 232, 233, 7, 12, 2, 2, 233, 239, 7, 40, 2, 2,
	234, 235, 7, 12, 2, 2, 235, 239, 7, 41, 2, 2, 236, 237, 7, 12, 2, 2, 237,
	239, 7, 42, 2, 2, 238, 186, 3, 2, 2, 2, 238, 188, 3, 2, 2, 2, 238, 190,
	3, 2, 2, 2, 238, 192, 3, 2, 2, 2, 238, 194, 3, 2, 2, 2, 238, 196, 3, 2,
	2, 2, 238, 198, 3, 2, 2, 2, 238, 200, 3, 2, 2, 2, 238, 202, 3, 2, 2, 2,
	238, 204, 3, 2, 2, 2, 238, 206, 3, 2, 2, 2, 238, 208, 3, 2, 2, 2, 238,
	210, 3, 2, 2, 2, 238, 212, 3, 2, 2, 2, 238, 214, 3, 2, 2, 2, 238, 216,
	3, 2, 2, 2, 238, 218, 3, 2, 2, 2, 238, 220, 3, 2, 2, 2, 238, 222, 3, 2,
	2, 2, 238, 224, 3, 2, 2, 2, 238, 226, 3, 2, 2, 2, 238, 228, 3, 2, 2, 2,
	238, 230, 3, 2, 2, 2, 238, 232, 3, 2, 2, 2, 238, 234, 3, 2, 2, 2, 238,
	236, 3, 2, 2, 2, 239, 33, 3, 2, 2, 2, 28, 35, 38, 47, 50, 53, 60, 69, 81,
	86, 92, 96, 101, 105, 110, 115, 119, 125, 135, 140, 153, 158, 164, 173,
	181, 184, 238,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'('", "')'", "';'", "':'", "','", "'*'", "'='",
	"'!='", "'?'", "'>'", "'>='", "'<'", "'<='",
}
var symbolicNames = []string{
	"", "Update", "Is", "In", "Not", "Null", "Set", "LeftParenthesis", "RightParenthesis",
	"Semicolon", "Colon", "Comma", "Asterisk", "Equal", "NotEqual", "QuestionMark",
	"GT", "GTE", "LT", "LTE", "Select", "Delete", "From", "Where", "Having",
	"Like", "And", "Or", "Order", "Group", "By", "Asc", "Desc", "Offset", "Limit",
	"As", "Count", "Sum", "Max", "Min", "Avg", "Number", "Literal", "BackQuotedString",
	"RawString", "WS",
}

var ruleNames = []string{
	"statement", "updateStatement", "assignmentList", "assignment", "whereClause",
	"orderByClause", "orderByField", "limitClause", "orExpression", "andExpression",
	"subExpression", "atomExpression", "likeExpression", "compareExpression",
	"inExpression", "colonField",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type UpdateStatementParserParser struct {
	*antlr.BaseParser
}

func NewUpdateStatementParserParser(input antlr.TokenStream) *UpdateStatementParserParser {
	this := new(UpdateStatementParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "UpdateStatementParser.g4"

	return this
}

// UpdateStatementParserParser tokens.
const (
	UpdateStatementParserParserEOF              = antlr.TokenEOF
	UpdateStatementParserParserUpdate           = 1
	UpdateStatementParserParserIs               = 2
	UpdateStatementParserParserIn               = 3
	UpdateStatementParserParserNot              = 4
	UpdateStatementParserParserNull             = 5
	UpdateStatementParserParserSet              = 6
	UpdateStatementParserParserLeftParenthesis  = 7
	UpdateStatementParserParserRightParenthesis = 8
	UpdateStatementParserParserSemicolon        = 9
	UpdateStatementParserParserColon            = 10
	UpdateStatementParserParserComma            = 11
	UpdateStatementParserParserAsterisk         = 12
	UpdateStatementParserParserEqual            = 13
	UpdateStatementParserParserNotEqual         = 14
	UpdateStatementParserParserQuestionMark     = 15
	UpdateStatementParserParserGT               = 16
	UpdateStatementParserParserGTE              = 17
	UpdateStatementParserParserLT               = 18
	UpdateStatementParserParserLTE              = 19
	UpdateStatementParserParserSelect           = 20
	UpdateStatementParserParserDelete           = 21
	UpdateStatementParserParserFrom             = 22
	UpdateStatementParserParserWhere            = 23
	UpdateStatementParserParserHaving           = 24
	UpdateStatementParserParserLike             = 25
	UpdateStatementParserParserAnd              = 26
	UpdateStatementParserParserOr               = 27
	UpdateStatementParserParserOrder            = 28
	UpdateStatementParserParserGroup            = 29
	UpdateStatementParserParserBy               = 30
	UpdateStatementParserParserAsc              = 31
	UpdateStatementParserParserDesc             = 32
	UpdateStatementParserParserOffset           = 33
	UpdateStatementParserParserLimit            = 34
	UpdateStatementParserParserAs               = 35
	UpdateStatementParserParserCount            = 36
	UpdateStatementParserParserSum              = 37
	UpdateStatementParserParserMax              = 38
	UpdateStatementParserParserMin              = 39
	UpdateStatementParserParserAvg              = 40
	UpdateStatementParserParserNumber           = 41
	UpdateStatementParserParserLiteral          = 42
	UpdateStatementParserParserBackQuotedString = 43
	UpdateStatementParserParserRawString        = 44
	UpdateStatementParserParserWS               = 45
)

// UpdateStatementParserParser rules.
const (
	UpdateStatementParserParserRULE_statement         = 0
	UpdateStatementParserParserRULE_updateStatement   = 1
	UpdateStatementParserParserRULE_assignmentList    = 2
	UpdateStatementParserParserRULE_assignment        = 3
	UpdateStatementParserParserRULE_whereClause       = 4
	UpdateStatementParserParserRULE_orderByClause     = 5
	UpdateStatementParserParserRULE_orderByField      = 6
	UpdateStatementParserParserRULE_limitClause       = 7
	UpdateStatementParserParserRULE_orExpression      = 8
	UpdateStatementParserParserRULE_andExpression     = 9
	UpdateStatementParserParserRULE_subExpression     = 10
	UpdateStatementParserParserRULE_atomExpression    = 11
	UpdateStatementParserParserRULE_likeExpression    = 12
	UpdateStatementParserParserRULE_compareExpression = 13
	UpdateStatementParserParserRULE_inExpression      = 14
	UpdateStatementParserParserRULE_colonField        = 15
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
	p.RuleIndex = UpdateStatementParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserEOF, 0)
}

func (s *StatementContext) UpdateStatement() IUpdateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *StatementContext) Semicolon() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserSemicolon, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *UpdateStatementParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UpdateStatementParserParserRULE_statement)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserUpdate {
		{
			p.SetState(32)
			p.UpdateStatement()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserSemicolon {
		{
			p.SetState(35)
			p.Match(UpdateStatementParserParserSemicolon)
		}

	}
	{
		p.SetState(38)
		p.Match(UpdateStatementParserParserEOF)
	}

	return localctx
}

// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTable returns the table token.
	GetTable() antlr.Token

	// SetTable sets the table token.
	SetTable(antlr.Token)

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	table  antlr.Token
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpdateStatementParserParserRULE_updateStatement
	return p
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) GetTable() antlr.Token { return s.table }

func (s *UpdateStatementContext) SetTable(v antlr.Token) { s.table = v }

func (s *UpdateStatementContext) Update() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserUpdate, 0)
}

func (s *UpdateStatementContext) Set() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserSet, 0)
}

func (s *UpdateStatementContext) AssignmentList() IAssignmentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentListContext)
}

func (s *UpdateStatementContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBackQuotedString, 0)
}

func (s *UpdateStatementContext) RawString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRawString, 0)
}

func (s *UpdateStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *UpdateStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *UpdateStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}

func (p *UpdateStatementParserParser) UpdateStatement() (localctx IUpdateStatementContext) {
	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UpdateStatementParserParserRULE_updateStatement)
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
		p.SetState(40)
		p.Match(UpdateStatementParserParserUpdate)
	}
	{
		p.SetState(41)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UpdateStatementContext).table = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UpdateStatementContext).table = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(42)
		p.Match(UpdateStatementParserParserSet)
	}
	{
		p.SetState(43)
		p.AssignmentList()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserWhere {
		{
			p.SetState(44)
			p.WhereClause()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserOrder {
		{
			p.SetState(47)
			p.OrderByClause()
		}

	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserLimit {
		{
			p.SetState(50)
			p.LimitClause()
		}

	}

	return localctx
}

// IAssignmentListContext is an interface to support dynamic dispatch.
type IAssignmentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentListContext differentiates from other interfaces.
	IsAssignmentListContext()
}

type AssignmentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentListContext() *AssignmentListContext {
	var p = new(AssignmentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpdateStatementParserParserRULE_assignmentList
	return p
}

func (*AssignmentListContext) IsAssignmentListContext() {}

func NewAssignmentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentListContext {
	var p = new(AssignmentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_assignmentList

	return p
}

func (s *AssignmentListContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentListContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *AssignmentListContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignmentListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(UpdateStatementParserParserComma)
}

func (s *AssignmentListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserComma, i)
}

func (s *AssignmentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterAssignmentList(s)
	}
}

func (s *AssignmentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitAssignmentList(s)
	}
}

func (p *UpdateStatementParserParser) AssignmentList() (localctx IAssignmentListContext) {
	localctx = NewAssignmentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UpdateStatementParserParserRULE_assignmentList)
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
		p.SetState(53)
		p.Assignment()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpdateStatementParserParserComma {
		{
			p.SetState(54)
			p.Match(UpdateStatementParserParserComma)
		}
		{
			p.SetState(55)
			p.Assignment()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UpdateStatementParserParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Equal() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserEqual, 0)
}

func (s *AssignmentContext) RawString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRawString, 0)
}

func (s *AssignmentContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBackQuotedString, 0)
}

func (s *AssignmentContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserQuestionMark, 0)
}

func (s *AssignmentContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *AssignmentContext) Literal() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLiteral, 0)
}

func (s *AssignmentContext) Number() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNumber, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *UpdateStatementParserParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UpdateStatementParserParserRULE_assignment)
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
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(62)
		p.Match(UpdateStatementParserParserEqual)
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserQuestionMark:
		{
			p.SetState(63)
			p.Match(UpdateStatementParserParserQuestionMark)
		}

	case UpdateStatementParserParserColon:
		{
			p.SetState(64)
			p.ColonField()
		}

	case UpdateStatementParserParserLiteral:
		{
			p.SetState(65)
			p.Match(UpdateStatementParserParserLiteral)
		}

	case UpdateStatementParserParserNumber:
		{
			p.SetState(66)
			p.Match(UpdateStatementParserParserNumber)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = UpdateStatementParserParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) Where() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserWhere, 0)
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
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *UpdateStatementParserParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UpdateStatementParserParserRULE_whereClause)

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
		p.SetState(69)
		p.Match(UpdateStatementParserParserWhere)
	}
	{
		p.SetState(70)
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
	p.RuleIndex = UpdateStatementParserParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) Order() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserOrder, 0)
}

func (s *OrderByClauseContext) By() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBy, 0)
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
	return s.GetTokens(UpdateStatementParserParserComma)
}

func (s *OrderByClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserComma, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (p *UpdateStatementParserParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UpdateStatementParserParserRULE_orderByClause)
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
		p.SetState(72)
		p.Match(UpdateStatementParserParserOrder)
	}
	{
		p.SetState(73)
		p.Match(UpdateStatementParserParserBy)
	}
	{
		p.SetState(74)
		p.OrderByField()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpdateStatementParserParserComma {
		{
			p.SetState(75)
			p.Match(UpdateStatementParserParserComma)
		}
		{
			p.SetState(76)
			p.OrderByField()
		}

		p.SetState(81)
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
	p.RuleIndex = UpdateStatementParserParserRULE_orderByField
	return p
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_orderByField

	return p
}

func (s *OrderByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByFieldContext) GetField() antlr.Token { return s.field }

func (s *OrderByFieldContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByFieldContext) SetField(v antlr.Token) { s.field = v }

func (s *OrderByFieldContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByFieldContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBackQuotedString, 0)
}

func (s *OrderByFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRawString, 0)
}

func (s *OrderByFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAsc, 0)
}

func (s *OrderByFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserDesc, 0)
}

func (s *OrderByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterOrderByField(s)
	}
}

func (s *OrderByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitOrderByField(s)
	}
}

func (p *UpdateStatementParserParser) OrderByField() (localctx IOrderByFieldContext) {
	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, UpdateStatementParserParserRULE_orderByField)
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
		p.SetState(82)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OrderByFieldContext).field = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OrderByFieldContext).field = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == UpdateStatementParserParserAsc || _la == UpdateStatementParserParserDesc {
		{
			p.SetState(83)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByFieldContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserAsc || _la == UpdateStatementParserParserDesc) {
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
	p.RuleIndex = UpdateStatementParserParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) Limit() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLimit, 0)
}

func (s *LimitClauseContext) ColonField() IColonFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColonFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColonFieldContext)
}

func (s *LimitClauseContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserQuestionMark, 0)
}

func (s *LimitClauseContext) Number() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNumber, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *UpdateStatementParserParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UpdateStatementParserParserRULE_limitClause)

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
		p.SetState(86)
		p.Match(UpdateStatementParserParserLimit)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserColon:
		{
			p.SetState(87)
			p.ColonField()
		}

	case UpdateStatementParserParserQuestionMark:
		{
			p.SetState(88)
			p.Match(UpdateStatementParserParserQuestionMark)
		}

	case UpdateStatementParserParserNumber:
		{
			p.SetState(89)
			p.Match(UpdateStatementParserParserNumber)
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
	p.RuleIndex = UpdateStatementParserParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_orExpression

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
	return s.GetTokens(UpdateStatementParserParserOr)
}

func (s *OrExpressionContext) Or(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserOr, i)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *UpdateStatementParserParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, UpdateStatementParserParserRULE_orExpression)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(92)
			p.SubExpression()
		}

	case 2:
		{
			p.SetState(93)
			p.AndExpression()
		}

	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpdateStatementParserParserOr {
		{
			p.SetState(96)
			p.Match(UpdateStatementParserParserOr)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(97)
				p.SubExpression()
			}

		case 2:
			{
				p.SetState(98)
				p.AndExpression()
			}

		}

		p.SetState(105)
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
	p.RuleIndex = UpdateStatementParserParserRULE_andExpression
	return p
}

func (*AndExpressionContext) IsAndExpressionContext() {}

func NewAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_andExpression

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
	return s.GetTokens(UpdateStatementParserParserAnd)
}

func (s *AndExpressionContext) And(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAnd, i)
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (p *UpdateStatementParserParser) AndExpression() (localctx IAndExpressionContext) {
	localctx = NewAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, UpdateStatementParserParserRULE_andExpression)
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
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserLeftParenthesis:
		{
			p.SetState(106)
			p.SubExpression()
		}

	case UpdateStatementParserParserBackQuotedString, UpdateStatementParserParserRawString:
		{
			p.SetState(107)
			p.AtomExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == UpdateStatementParserParserAnd {
		{
			p.SetState(110)
			p.Match(UpdateStatementParserParserAnd)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case UpdateStatementParserParserLeftParenthesis:
			{
				p.SetState(111)
				p.SubExpression()
			}

		case UpdateStatementParserParserBackQuotedString, UpdateStatementParserParserRawString:
			{
				p.SetState(112)
				p.AtomExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(119)
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
	p.RuleIndex = UpdateStatementParserParserRULE_subExpression
	return p
}

func (*SubExpressionContext) IsSubExpressionContext() {}

func NewSubExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubExpressionContext {
	var p = new(SubExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_subExpression

	return p
}

func (s *SubExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLeftParenthesis, 0)
}

func (s *SubExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRightParenthesis, 0)
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
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterSubExpression(s)
	}
}

func (s *SubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitSubExpression(s)
	}
}

func (p *UpdateStatementParserParser) SubExpression() (localctx ISubExpressionContext) {
	localctx = NewSubExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, UpdateStatementParserParserRULE_subExpression)

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
		p.SetState(120)
		p.Match(UpdateStatementParserParserLeftParenthesis)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(121)
			p.AndExpression()
		}

	case 2:
		{
			p.SetState(122)
			p.OrExpression()
		}

	}
	{
		p.SetState(125)
		p.Match(UpdateStatementParserParserRightParenthesis)
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
	p.RuleIndex = UpdateStatementParserParserRULE_atomExpression
	return p
}

func (*AtomExpressionContext) IsAtomExpressionContext() {}

func NewAtomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_atomExpression

	return p
}

func (s *AtomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomExpressionContext) GetField() antlr.Token { return s.field }

func (s *AtomExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AtomExpressionContext) SetField(v antlr.Token) { s.field = v }

func (s *AtomExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AtomExpressionContext) Like() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLike, 0)
}

func (s *AtomExpressionContext) LikeExpression() ILikeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeExpressionContext)
}

func (s *AtomExpressionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBackQuotedString, 0)
}

func (s *AtomExpressionContext) RawString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRawString, 0)
}

func (s *AtomExpressionContext) Is() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserIs, 0)
}

func (s *AtomExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNull, 0)
}

func (s *AtomExpressionContext) Not() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNot, 0)
}

func (s *AtomExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserIn, 0)
}

func (s *AtomExpressionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLeftParenthesis, 0)
}

func (s *AtomExpressionContext) InExpression() IInExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInExpressionContext)
}

func (s *AtomExpressionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRightParenthesis, 0)
}

func (s *AtomExpressionContext) CompareExpression() ICompareExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareExpressionContext)
}

func (s *AtomExpressionContext) Equal() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserEqual, 0)
}

func (s *AtomExpressionContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNotEqual, 0)
}

func (s *AtomExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserGT, 0)
}

func (s *AtomExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLT, 0)
}

func (s *AtomExpressionContext) GTE() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserGTE, 0)
}

func (s *AtomExpressionContext) LTE() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLTE, 0)
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterAtomExpression(s)
	}
}

func (s *AtomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitAtomExpression(s)
	}
}

func (p *UpdateStatementParserParser) AtomExpression() (localctx IAtomExpressionContext) {
	localctx = NewAtomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, UpdateStatementParserParserRULE_atomExpression)
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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(128)
			p.Match(UpdateStatementParserParserLike)
		}
		{
			p.SetState(129)
			p.LikeExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(131)
			p.Match(UpdateStatementParserParserIs)
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == UpdateStatementParserParserNot {
			{
				p.SetState(132)
				p.Match(UpdateStatementParserParserNot)
			}

		}
		{
			p.SetState(135)
			p.Match(UpdateStatementParserParserNull)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(136)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == UpdateStatementParserParserNot {
			{
				p.SetState(137)
				p.Match(UpdateStatementParserParserNot)
			}

		}
		{
			p.SetState(140)
			p.Match(UpdateStatementParserParserIn)
		}
		{
			p.SetState(141)
			p.Match(UpdateStatementParserParserLeftParenthesis)
		}
		{
			p.SetState(142)
			p.InExpression()
		}
		{
			p.SetState(143)
			p.Match(UpdateStatementParserParserRightParenthesis)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(146)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserEqual || _la == UpdateStatementParserParserNotEqual) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(147)
			p.CompareExpression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == UpdateStatementParserParserBackQuotedString || _la == UpdateStatementParserParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(149)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AtomExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<UpdateStatementParserParserGT)|(1<<UpdateStatementParserParserGTE)|(1<<UpdateStatementParserParserLT)|(1<<UpdateStatementParserParserLTE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AtomExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(150)
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
	p.RuleIndex = UpdateStatementParserParserRULE_likeExpression
	return p
}

func (*LikeExpressionContext) IsLikeExpressionContext() {}

func NewLikeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeExpressionContext {
	var p = new(LikeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_likeExpression

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
	return s.GetToken(UpdateStatementParserParserLiteral, 0)
}

func (s *LikeExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserQuestionMark, 0)
}

func (s *LikeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterLikeExpression(s)
	}
}

func (s *LikeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitLikeExpression(s)
	}
}

func (p *UpdateStatementParserParser) LikeExpression() (localctx ILikeExpressionContext) {
	localctx = NewLikeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, UpdateStatementParserParserRULE_likeExpression)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.ColonField()
		}

	case UpdateStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(UpdateStatementParserParserLiteral)
		}

	case UpdateStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Match(UpdateStatementParserParserQuestionMark)
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
	p.RuleIndex = UpdateStatementParserParserRULE_compareExpression
	return p
}

func (*CompareExpressionContext) IsCompareExpressionContext() {}

func NewCompareExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareExpressionContext {
	var p = new(CompareExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_compareExpression

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
	return s.GetToken(UpdateStatementParserParserNumber, 0)
}

func (s *CompareExpressionContext) Literal() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLiteral, 0)
}

func (s *CompareExpressionContext) QuestionMark() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserQuestionMark, 0)
}

func (s *CompareExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterCompareExpression(s)
	}
}

func (s *CompareExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitCompareExpression(s)
	}
}

func (p *UpdateStatementParserParser) CompareExpression() (localctx ICompareExpressionContext) {
	localctx = NewCompareExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, UpdateStatementParserParserRULE_compareExpression)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.ColonField()
		}

	case UpdateStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(UpdateStatementParserParserNumber)
		}

	case UpdateStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(UpdateStatementParserParserLiteral)
		}

	case UpdateStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.Match(UpdateStatementParserParserQuestionMark)
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
	p.RuleIndex = UpdateStatementParserParserRULE_inExpression
	return p
}

func (*InExpressionContext) IsInExpressionContext() {}

func NewInExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InExpressionContext {
	var p = new(InExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_inExpression

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
	return s.GetToken(UpdateStatementParserParserQuestionMark, 0)
}

func (s *InExpressionContext) AllNumber() []antlr.TerminalNode {
	return s.GetTokens(UpdateStatementParserParserNumber)
}

func (s *InExpressionContext) Number(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNumber, i)
}

func (s *InExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(UpdateStatementParserParserComma)
}

func (s *InExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserComma, i)
}

func (s *InExpressionContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(UpdateStatementParserParserLiteral)
}

func (s *InExpressionContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLiteral, i)
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitInExpression(s)
	}
}

func (p *UpdateStatementParserParser) InExpression() (localctx IInExpressionContext) {
	localctx = NewInExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, UpdateStatementParserParserRULE_inExpression)
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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UpdateStatementParserParserColon:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.ColonField()
		}

	case UpdateStatementParserParserQuestionMark:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(UpdateStatementParserParserQuestionMark)
		}

	case UpdateStatementParserParserNumber:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(UpdateStatementParserParserNumber)
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == UpdateStatementParserParserComma {
			{
				p.SetState(167)
				p.Match(UpdateStatementParserParserComma)
			}
			{
				p.SetState(168)
				p.Match(UpdateStatementParserParserNumber)
			}

			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case UpdateStatementParserParserLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.Match(UpdateStatementParserParserLiteral)
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == UpdateStatementParserParserComma {
			{
				p.SetState(175)
				p.Match(UpdateStatementParserParserComma)
			}
			{
				p.SetState(176)
				p.Match(UpdateStatementParserParserLiteral)
			}

			p.SetState(181)
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
	p.RuleIndex = UpdateStatementParserParserRULE_colonField
	return p
}

func (*ColonFieldContext) IsColonFieldContext() {}

func NewColonFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColonFieldContext {
	var p = new(ColonFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UpdateStatementParserParserRULE_colonField

	return p
}

func (s *ColonFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ColonFieldContext) Colon() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserColon, 0)
}

func (s *ColonFieldContext) RawString() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserRawString, 0)
}

func (s *ColonFieldContext) Is() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserIs, 0)
}

func (s *ColonFieldContext) In() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserIn, 0)
}

func (s *ColonFieldContext) Not() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNot, 0)
}

func (s *ColonFieldContext) Null() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserNull, 0)
}

func (s *ColonFieldContext) Select() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserSelect, 0)
}

func (s *ColonFieldContext) Delete() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserDelete, 0)
}

func (s *ColonFieldContext) From() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserFrom, 0)
}

func (s *ColonFieldContext) Where() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserWhere, 0)
}

func (s *ColonFieldContext) Having() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserHaving, 0)
}

func (s *ColonFieldContext) Like() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLike, 0)
}

func (s *ColonFieldContext) And() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAnd, 0)
}

func (s *ColonFieldContext) Or() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserOr, 0)
}

func (s *ColonFieldContext) Order() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserOrder, 0)
}

func (s *ColonFieldContext) Group() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserGroup, 0)
}

func (s *ColonFieldContext) By() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserBy, 0)
}

func (s *ColonFieldContext) Asc() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAsc, 0)
}

func (s *ColonFieldContext) Desc() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserDesc, 0)
}

func (s *ColonFieldContext) Offset() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserOffset, 0)
}

func (s *ColonFieldContext) Limit() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserLimit, 0)
}

func (s *ColonFieldContext) As() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAs, 0)
}

func (s *ColonFieldContext) Count() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserCount, 0)
}

func (s *ColonFieldContext) Sum() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserSum, 0)
}

func (s *ColonFieldContext) Max() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserMax, 0)
}

func (s *ColonFieldContext) Min() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserMin, 0)
}

func (s *ColonFieldContext) Avg() antlr.TerminalNode {
	return s.GetToken(UpdateStatementParserParserAvg, 0)
}

func (s *ColonFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColonFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColonFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.EnterColonField(s)
	}
}

func (s *ColonFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UpdateStatementParserListener); ok {
		listenerT.ExitColonField(s)
	}
}

func (p *UpdateStatementParserParser) ColonField() (localctx IColonFieldContext) {
	localctx = NewColonFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, UpdateStatementParserParserRULE_colonField)

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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(185)
			p.Match(UpdateStatementParserParserRawString)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(187)
			p.Match(UpdateStatementParserParserIs)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(189)
			p.Match(UpdateStatementParserParserIn)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(190)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(191)
			p.Match(UpdateStatementParserParserNot)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(192)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(193)
			p.Match(UpdateStatementParserParserNull)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(194)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(195)
			p.Match(UpdateStatementParserParserSelect)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(196)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(197)
			p.Match(UpdateStatementParserParserDelete)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(198)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(199)
			p.Match(UpdateStatementParserParserFrom)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(200)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(201)
			p.Match(UpdateStatementParserParserWhere)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(202)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(203)
			p.Match(UpdateStatementParserParserHaving)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(204)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(205)
			p.Match(UpdateStatementParserParserLike)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(206)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(207)
			p.Match(UpdateStatementParserParserAnd)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(208)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(209)
			p.Match(UpdateStatementParserParserOr)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(210)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(211)
			p.Match(UpdateStatementParserParserOrder)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(212)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(213)
			p.Match(UpdateStatementParserParserGroup)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(214)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(215)
			p.Match(UpdateStatementParserParserBy)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(216)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(217)
			p.Match(UpdateStatementParserParserAsc)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(218)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(219)
			p.Match(UpdateStatementParserParserDesc)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(220)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(221)
			p.Match(UpdateStatementParserParserOffset)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(222)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(223)
			p.Match(UpdateStatementParserParserLimit)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(224)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(225)
			p.Match(UpdateStatementParserParserAs)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(226)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(227)
			p.Match(UpdateStatementParserParserCount)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(228)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(229)
			p.Match(UpdateStatementParserParserSum)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(230)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(231)
			p.Match(UpdateStatementParserParserMax)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(232)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(233)
			p.Match(UpdateStatementParserParserMin)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(234)
			p.Match(UpdateStatementParserParserColon)
		}
		{
			p.SetState(235)
			p.Match(UpdateStatementParserParserAvg)
		}

	}

	return localctx
}
