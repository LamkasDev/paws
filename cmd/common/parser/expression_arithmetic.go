package parser

import (
	"fmt"

	"github.com/LamkasDev/paws/cmd/common/lexer"
)

type ParserExpressionArithmetic struct {
	Left     *ParserExpression
	Right    *ParserExpression
	Operator *lexer.LexerToken
}

func NewParserExpressionArithmetic(left *ParserExpression, right *ParserExpression, operator *lexer.LexerToken) *ParserExpression {
	return &ParserExpression{
		Type: ParserExpressionTypeArithmetic,
		Data: &ParserExpressionArithmetic{
			Left:     left,
			Right:    right,
			Operator: operator,
		},
	}
}

func (expression *ParserExpressionArithmetic) Sprint() string {
	return fmt.Sprintf("(%+v %s %+v)", expression.Left.Sprint(), expression.Operator.Value, expression.Right.Sprint())
}
