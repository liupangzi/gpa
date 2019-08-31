package selectlistener

import (
	"strings"

	parser "github.com/liupangzi/gpa/parser/select"
)

type SelectListener struct {
	*parser.BaseSelectStatementParserListener

	Limit string

	// param limit(?)'s reversed position of []Params
	// e.g. limit ?, ? == 0, limit ? offset ? == 1, limit ? == 0, limit ? offset 10 == 0
	ReversedIndex int

	IsInQuery bool
}

func (s *SelectListener) EnterOffsetValue(ctx *parser.OffsetValueContext) {
	// `limit ? offset ?` makes ReversedIndex++
	if s.Limit == "?" && ctx.GetText() == "?" {
		s.ReversedIndex += 1
	}
}

func (s *SelectListener) EnterLimitValue(ctx *parser.LimitValueContext) {
	s.Limit = ctx.GetText()
}

func (s *SelectListener) EnterInExpression(ctx *parser.InExpressionContext) {
	s.IsInQuery = s.IsInQuery || strings.ContainsAny(ctx.GetText(), "?:")
}
