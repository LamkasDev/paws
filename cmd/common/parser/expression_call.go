package parser

import "fmt"

type ParserExpressionCall struct {
	Symbol *ParserSymbol
}

func NewParserExpressionCall(symbol *ParserSymbol) *ParserExpression {
	return &ParserExpression{
		Type: ParserExpressionTypeCall,
		Data: &ParserExpressionCall{
			Symbol: symbol,
		},
	}
}

func (expression *ParserExpressionCall) Sprint() string {
	return fmt.Sprintf("%s()", expression.Symbol.Name)
}
