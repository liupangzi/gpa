package updatelistener

import (
	"strings"

	parser "github.com/liupangzi/gpa/parser/update"
)

type UpdateListener struct {
	*parser.BaseUpdateStatementParserListener

	IsInQuery bool
}

func (u *UpdateListener) EnterInExpression(ctx *parser.InExpressionContext) {
	u.IsInQuery = u.IsInQuery || strings.ContainsAny(ctx.GetText(), "?:")
}
