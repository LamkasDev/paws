package parser

import "fmt"

type ParserExpressionAssignment struct {
	Symbol *ParserSymbol
	Value  *ParserExpression
}

func NewParserExpressionAssignment(symbol *ParserSymbol, value *ParserExpression) *ParserExpression {
	return &ParserExpression{
		Type: ParserExpressionTypeAssignment,
		Data: &ParserExpressionAssignment{
			Symbol: symbol,
			Value:  value,
		},
	}
}

func (expression *ParserExpressionAssignment) Sprint() string {
	return fmt.Sprintf("%s = %+v", expression.Symbol.Name, expression.Value.Sprint())
}
