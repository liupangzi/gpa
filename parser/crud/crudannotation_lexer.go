// Code generated from CRUDAnnotation.g4 by ANTLR 4.7.2. DO NOT EDIT.

package crudannotation

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 90, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	62, 10, 9, 12, 9, 14, 9, 65, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	72, 10, 9, 12, 9, 14, 9, 75, 11, 9, 3, 9, 5, 9, 78, 10, 9, 3, 10, 3, 10,
	5, 10, 82, 10, 10, 3, 11, 6, 11, 85, 10, 11, 13, 11, 14, 11, 86, 3, 11,
	3, 11, 4, 63, 73, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 2, 21, 11, 3, 2, 17, 4, 2, 75, 75, 107, 107, 4, 2, 80, 80, 112,
	112, 4, 2, 85, 85, 117, 117, 4, 2, 71, 71, 103, 103, 4, 2, 84, 84, 116,
	116, 4, 2, 86, 86, 118, 118, 4, 2, 78, 78, 110, 110, 4, 2, 69, 69, 101,
	101, 4, 2, 87, 87, 119, 119, 4, 2, 82, 82, 114, 114, 4, 2, 70, 70, 102,
	102, 4, 2, 67, 67, 99, 99, 3, 2, 12, 33, 3, 2, 129, 161, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 95, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25,
	3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 36, 3, 2, 2, 2, 13,
	43, 3, 2, 2, 2, 15, 50, 3, 2, 2, 2, 17, 77, 3, 2, 2, 2, 19, 81, 3, 2, 2,
	2, 21, 84, 3, 2, 2, 2, 23, 24, 7, 66, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26,
	7, 42, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 43, 2, 2, 28, 8, 3, 2, 2, 2,
	29, 30, 9, 2, 2, 2, 30, 31, 9, 3, 2, 2, 31, 32, 9, 4, 2, 2, 32, 33, 9,
	5, 2, 2, 33, 34, 9, 6, 2, 2, 34, 35, 9, 7, 2, 2, 35, 10, 3, 2, 2, 2, 36,
	37, 9, 4, 2, 2, 37, 38, 9, 5, 2, 2, 38, 39, 9, 8, 2, 2, 39, 40, 9, 5, 2,
	2, 40, 41, 9, 9, 2, 2, 41, 42, 9, 7, 2, 2, 42, 12, 3, 2, 2, 2, 43, 44,
	9, 10, 2, 2, 44, 45, 9, 11, 2, 2, 45, 46, 9, 12, 2, 2, 46, 47, 9, 13, 2,
	2, 47, 48, 9, 7, 2, 2, 48, 49, 9, 5, 2, 2, 49, 14, 3, 2, 2, 2, 50, 51,
	9, 12, 2, 2, 51, 52, 9, 5, 2, 2, 52, 53, 9, 8, 2, 2, 53, 54, 9, 5, 2, 2,
	54, 55, 9, 7, 2, 2, 55, 56, 9, 5, 2, 2, 56, 16, 3, 2, 2, 2, 57, 63, 7,
	36, 2, 2, 58, 59, 7, 94, 2, 2, 59, 62, 7, 36, 2, 2, 60, 62, 5, 19, 10,
	2, 61, 58, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	66, 78, 7, 36, 2, 2, 67, 73, 7, 41, 2, 2, 68, 69, 7, 94, 2, 2, 69, 72,
	7, 41, 2, 2, 70, 72, 5, 19, 10, 2, 71, 68, 3, 2, 2, 2, 71, 70, 3, 2, 2,
	2, 72, 75, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 76,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 7, 41, 2, 2, 77, 57, 3, 2, 2, 2,
	77, 67, 3, 2, 2, 2, 78, 18, 3, 2, 2, 2, 79, 82, 10, 14, 2, 2, 80, 82, 10,
	15, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 20, 3, 2, 2, 2, 83,
	85, 9, 16, 2, 2, 84, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 84, 3, 2,
	2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 8, 11, 2, 2, 89,
	22, 3, 2, 2, 2, 10, 2, 61, 63, 71, 73, 77, 81, 86, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'@'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "At", "LeftParenthesis", "RightParenthesis", "Insert", "Select", "Update",
	"Delete", "Literal", "WS",
}

var lexerRuleNames = []string{
	"At", "LeftParenthesis", "RightParenthesis", "Insert", "Select", "Update",
	"Delete", "Literal", "UnicodeCharacter", "WS",
}

type CRUDAnnotationLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCRUDAnnotationLexer(input antlr.CharStream) *CRUDAnnotationLexer {

	l := new(CRUDAnnotationLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CRUDAnnotation.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CRUDAnnotationLexer tokens.
const (
	CRUDAnnotationLexerAt               = 1
	CRUDAnnotationLexerLeftParenthesis  = 2
	CRUDAnnotationLexerRightParenthesis = 3
	CRUDAnnotationLexerInsert           = 4
	CRUDAnnotationLexerSelect           = 5
	CRUDAnnotationLexerUpdate           = 6
	CRUDAnnotationLexerDelete           = 7
	CRUDAnnotationLexerLiteral          = 8
	CRUDAnnotationLexerWS               = 9
)
