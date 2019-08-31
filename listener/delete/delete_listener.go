package deletelistener

import (
	"strings"

	parser "github.com/liupangzi/gpa/parser/delete"
)

type DeleteListener struct {
	*parser.BaseDeleteStatementParserListener

	IsInQuery bool
}

func (d *DeleteListener) EnterInExpression(ctx *parser.InExpressionContext) {
	d.IsInQuery = d.IsInQuery || strings.ContainsAny(ctx.GetText(), "?:")
}
