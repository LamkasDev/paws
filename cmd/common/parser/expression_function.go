package parser

import (
	"fmt"

	"github.com/LamkasDev/paws/cmd/common/lexer"
)

type ParserExpressionFunction struct {
	Name       *lexer.LexerToken
	Statements []*ParserExpression
}

func NewParserExpressionFunction(name *lexer.LexerToken) *ParserExpression {
	return &ParserExpression{
		Type: ParserExpressionTypePrimitive,
		Data: &ParserExpressionFunction{
			Name:       name,
			Statements: []*ParserExpression{},
		},
	}
}

func (expression *ParserExpressionFunction) Sprint() string {
	str := fmt.Sprintf("fn %s()", expression.Name.Value)
	for _, statement := range expression.Statements {
		str = fmt.Sprintf("%s\n%s", str, statement.Sprint())
	}

	return str
}
