// Code generated from CRUDAnnotation.g4 by ANTLR 4.7.2. DO NOT EDIT.

package crudannotation // CRUDAnnotation
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 50, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 25, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 47, 2, 24, 3, 2, 2, 2, 4, 26, 3, 2,
	2, 2, 6, 31, 3, 2, 2, 2, 8, 37, 3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 13,
	5, 4, 3, 2, 13, 14, 7, 2, 2, 3, 14, 25, 3, 2, 2, 2, 15, 16, 5, 6, 4, 2,
	16, 17, 7, 2, 2, 3, 17, 25, 3, 2, 2, 2, 18, 19, 5, 8, 5, 2, 19, 20, 7,
	2, 2, 3, 20, 25, 3, 2, 2, 2, 21, 22, 5, 10, 6, 2, 22, 23, 7, 2, 2, 3, 23,
	25, 3, 2, 2, 2, 24, 12, 3, 2, 2, 2, 24, 15, 3, 2, 2, 2, 24, 18, 3, 2, 2,
	2, 24, 21, 3, 2, 2, 2, 25, 3, 3, 2, 2, 2, 26, 27, 7, 3, 2, 2, 27, 28, 7,
	6, 2, 2, 28, 29, 7, 4, 2, 2, 29, 30, 7, 5, 2, 2, 30, 5, 3, 2, 2, 2, 31,
	32, 7, 3, 2, 2, 32, 33, 7, 7, 2, 2, 33, 34, 7, 4, 2, 2, 34, 35, 7, 10,
	2, 2, 35, 36, 7, 5, 2, 2, 36, 7, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39,
	7, 8, 2, 2, 39, 40, 7, 4, 2, 2, 40, 41, 7, 10, 2, 2, 41, 42, 7, 5, 2, 2,
	42, 9, 3, 2, 2, 2, 43, 44, 7, 3, 2, 2, 44, 45, 7, 9, 2, 2, 45, 46, 7, 4,
	2, 2, 46, 47, 7, 10, 2, 2, 47, 48, 7, 5, 2, 2, 48, 11, 3, 2, 2, 2, 3, 24,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'@'", "'('", "')'",
}
var symbolicNames = []string{
	"", "At", "LeftParenthesis", "RightParenthesis", "Insert", "Select", "Update",
	"Delete", "Literal", "WS",
}

var ruleNames = []string{
	"crud", "insertClause", "selectClause", "updateClause", "deleteClause",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CRUDAnnotationParser struct {
	*antlr.BaseParser
}

func NewCRUDAnnotationParser(input antlr.TokenStream) *CRUDAnnotationParser {
	this := new(CRUDAnnotationParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CRUDAnnotation.g4"

	return this
}

// CRUDAnnotationParser tokens.
const (
	CRUDAnnotationParserEOF              = antlr.TokenEOF
	CRUDAnnotationParserAt               = 1
	CRUDAnnotationParserLeftParenthesis  = 2
	CRUDAnnotationParserRightParenthesis = 3
	CRUDAnnotationParserInsert           = 4
	CRUDAnnotationParserSelect           = 5
	CRUDAnnotationParserUpdate           = 6
	CRUDAnnotationParserDelete           = 7
	CRUDAnnotationParserLiteral          = 8
	CRUDAnnotationParserWS               = 9
)

// CRUDAnnotationParser rules.
const (
	CRUDAnnotationParserRULE_crud         = 0
	CRUDAnnotationParserRULE_insertClause = 1
	CRUDAnnotationParserRULE_selectClause = 2
	CRUDAnnotationParserRULE_updateClause = 3
	CRUDAnnotationParserRULE_deleteClause = 4
)

// ICrudContext is an interface to support dynamic dispatch.
type ICrudContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCrudContext differentiates from other interfaces.
	IsCrudContext()
}

type CrudContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrudContext() *CrudContext {
	var p = new(CrudContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CRUDAnnotationParserRULE_crud
	return p
}

func (*CrudContext) IsCrudContext() {}

func NewCrudContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CrudContext {
	var p = new(CrudContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CRUDAnnotationParserRULE_crud

	return p
}

func (s *CrudContext) GetParser() antlr.Parser { return s.parser }

func (s *CrudContext) InsertClause() IInsertClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertClauseContext)
}

func (s *CrudContext) EOF() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserEOF, 0)
}

func (s *CrudContext) SelectClause() ISelectClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectClauseContext)
}

func (s *CrudContext) UpdateClause() IUpdateClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateClauseContext)
}

func (s *CrudContext) DeleteClause() IDeleteClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteClauseContext)
}

func (s *CrudContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrudContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CrudContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.EnterCrud(s)
	}
}

func (s *CrudContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.ExitCrud(s)
	}
}

func (p *CRUDAnnotationParser) Crud() (localctx ICrudContext) {
	localctx = NewCrudContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CRUDAnnotationParserRULE_crud)

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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.InsertClause()
		}
		{
			p.SetState(11)
			p.Match(CRUDAnnotationParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(13)
			p.SelectClause()
		}
		{
			p.SetState(14)
			p.Match(CRUDAnnotationParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(16)
			p.UpdateClause()
		}
		{
			p.SetState(17)
			p.Match(CRUDAnnotationParserEOF)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(19)
			p.DeleteClause()
		}
		{
			p.SetState(20)
			p.Match(CRUDAnnotationParserEOF)
		}

	}

	return localctx
}

// IInsertClauseContext is an interface to support dynamic dispatch.
type IInsertClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertClauseContext differentiates from other interfaces.
	IsInsertClauseContext()
}

type InsertClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertClauseContext() *InsertClauseContext {
	var p = new(InsertClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CRUDAnnotationParserRULE_insertClause
	return p
}

func (*InsertClauseContext) IsInsertClauseContext() {}

func NewInsertClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertClauseContext {
	var p = new(InsertClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CRUDAnnotationParserRULE_insertClause

	return p
}

func (s *InsertClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertClauseContext) At() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserAt, 0)
}

func (s *InsertClauseContext) Insert() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserInsert, 0)
}

func (s *InsertClauseContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLeftParenthesis, 0)
}

func (s *InsertClauseContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserRightParenthesis, 0)
}

func (s *InsertClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.EnterInsertClause(s)
	}
}

func (s *InsertClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.ExitInsertClause(s)
	}
}

func (p *CRUDAnnotationParser) InsertClause() (localctx IInsertClauseContext) {
	localctx = NewInsertClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CRUDAnnotationParserRULE_insertClause)

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
		p.SetState(24)
		p.Match(CRUDAnnotationParserAt)
	}
	{
		p.SetState(25)
		p.Match(CRUDAnnotationParserInsert)
	}
	{
		p.SetState(26)
		p.Match(CRUDAnnotationParserLeftParenthesis)
	}
	{
		p.SetState(27)
		p.Match(CRUDAnnotationParserRightParenthesis)
	}

	return localctx
}

// ISelectClauseContext is an interface to support dynamic dispatch.
type ISelectClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSql returns the sql token.
	GetSql() antlr.Token

	// SetSql sets the sql token.
	SetSql(antlr.Token)

	// IsSelectClauseContext differentiates from other interfaces.
	IsSelectClauseContext()
}

type SelectClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sql    antlr.Token
}

func NewEmptySelectClauseContext() *SelectClauseContext {
	var p = new(SelectClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CRUDAnnotationParserRULE_selectClause
	return p
}

func (*SelectClauseContext) IsSelectClauseContext() {}

func NewSelectClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectClauseContext {
	var p = new(SelectClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CRUDAnnotationParserRULE_selectClause

	return p
}

func (s *SelectClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectClauseContext) GetSql() antlr.Token { return s.sql }

func (s *SelectClauseContext) SetSql(v antlr.Token) { s.sql = v }

func (s *SelectClauseContext) At() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserAt, 0)
}

func (s *SelectClauseContext) Select() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserSelect, 0)
}

func (s *SelectClauseContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLeftParenthesis, 0)
}

func (s *SelectClauseContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserRightParenthesis, 0)
}

func (s *SelectClauseContext) Literal() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLiteral, 0)
}

func (s *SelectClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.EnterSelectClause(s)
	}
}

func (s *SelectClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.ExitSelectClause(s)
	}
}

func (p *CRUDAnnotationParser) SelectClause() (localctx ISelectClauseContext) {
	localctx = NewSelectClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CRUDAnnotationParserRULE_selectClause)

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
		p.SetState(29)
		p.Match(CRUDAnnotationParserAt)
	}
	{
		p.SetState(30)
		p.Match(CRUDAnnotationParserSelect)
	}
	{
		p.SetState(31)
		p.Match(CRUDAnnotationParserLeftParenthesis)
	}
	{
		p.SetState(32)

		var _m = p.Match(CRUDAnnotationParserLiteral)

		localctx.(*SelectClauseContext).sql = _m
	}
	{
		p.SetState(33)
		p.Match(CRUDAnnotationParserRightParenthesis)
	}

	return localctx
}

// IUpdateClauseContext is an interface to support dynamic dispatch.
type IUpdateClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSql returns the sql token.
	GetSql() antlr.Token

	// SetSql sets the sql token.
	SetSql(antlr.Token)

	// IsUpdateClauseContext differentiates from other interfaces.
	IsUpdateClauseContext()
}

type UpdateClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sql    antlr.Token
}

func NewEmptyUpdateClauseContext() *UpdateClauseContext {
	var p = new(UpdateClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CRUDAnnotationParserRULE_updateClause
	return p
}

func (*UpdateClauseContext) IsUpdateClauseContext() {}

func NewUpdateClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateClauseContext {
	var p = new(UpdateClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CRUDAnnotationParserRULE_updateClause

	return p
}

func (s *UpdateClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateClauseContext) GetSql() antlr.Token { return s.sql }

func (s *UpdateClauseContext) SetSql(v antlr.Token) { s.sql = v }

func (s *UpdateClauseContext) At() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserAt, 0)
}

func (s *UpdateClauseContext) Update() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserUpdate, 0)
}

func (s *UpdateClauseContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLeftParenthesis, 0)
}

func (s *UpdateClauseContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserRightParenthesis, 0)
}

func (s *UpdateClauseContext) Literal() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLiteral, 0)
}

func (s *UpdateClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.EnterUpdateClause(s)
	}
}

func (s *UpdateClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.ExitUpdateClause(s)
	}
}

func (p *CRUDAnnotationParser) UpdateClause() (localctx IUpdateClauseContext) {
	localctx = NewUpdateClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CRUDAnnotationParserRULE_updateClause)

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
		p.SetState(35)
		p.Match(CRUDAnnotationParserAt)
	}
	{
		p.SetState(36)
		p.Match(CRUDAnnotationParserUpdate)
	}
	{
		p.SetState(37)
		p.Match(CRUDAnnotationParserLeftParenthesis)
	}
	{
		p.SetState(38)

		var _m = p.Match(CRUDAnnotationParserLiteral)

		localctx.(*UpdateClauseContext).sql = _m
	}
	{
		p.SetState(39)
		p.Match(CRUDAnnotationParserRightParenthesis)
	}

	return localctx
}

// IDeleteClauseContext is an interface to support dynamic dispatch.
type IDeleteClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSql returns the sql token.
	GetSql() antlr.Token

	// SetSql sets the sql token.
	SetSql(antlr.Token)

	// IsDeleteClauseContext differentiates from other interfaces.
	IsDeleteClauseContext()
}

type DeleteClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sql    antlr.Token
}

func NewEmptyDeleteClauseContext() *DeleteClauseContext {
	var p = new(DeleteClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CRUDAnnotationParserRULE_deleteClause
	return p
}

func (*DeleteClauseContext) IsDeleteClauseContext() {}

func NewDeleteClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteClauseContext {
	var p = new(DeleteClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CRUDAnnotationParserRULE_deleteClause

	return p
}

func (s *DeleteClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteClauseContext) GetSql() antlr.Token { return s.sql }

func (s *DeleteClauseContext) SetSql(v antlr.Token) { s.sql = v }

func (s *DeleteClauseContext) At() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserAt, 0)
}

func (s *DeleteClauseContext) Delete() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserDelete, 0)
}

func (s *DeleteClauseContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLeftParenthesis, 0)
}

func (s *DeleteClauseContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserRightParenthesis, 0)
}

func (s *DeleteClauseContext) Literal() antlr.TerminalNode {
	return s.GetToken(CRUDAnnotationParserLiteral, 0)
}

func (s *DeleteClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.EnterDeleteClause(s)
	}
}

func (s *DeleteClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CRUDAnnotationListener); ok {
		listenerT.ExitDeleteClause(s)
	}
}

func (p *CRUDAnnotationParser) DeleteClause() (localctx IDeleteClauseContext) {
	localctx = NewDeleteClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CRUDAnnotationParserRULE_deleteClause)

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
		p.SetState(41)
		p.Match(CRUDAnnotationParserAt)
	}
	{
		p.SetState(42)
		p.Match(CRUDAnnotationParserDelete)
	}
	{
		p.SetState(43)
		p.Match(CRUDAnnotationParserLeftParenthesis)
	}
	{
		p.SetState(44)

		var _m = p.Match(CRUDAnnotationParserLiteral)

		localctx.(*DeleteClauseContext).sql = _m
	}
	{
		p.SetState(45)
		p.Match(CRUDAnnotationParserRightParenthesis)
	}

	return localctx
}
