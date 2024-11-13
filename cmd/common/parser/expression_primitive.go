package parser

import "fmt"

type ParserExpressionPrimitive struct {
	Value interface{}
}

func NewParserExpressionPrimitive(value interface{}) *ParserExpression {
	return &ParserExpression{
		Type: ParserExpressionTypePrimitive,
		Data: &ParserExpressionPrimitive{
			Value: value,
		},
	}
}

func (expression *ParserExpressionPrimitive) Sprint() string {
	return fmt.Sprintf("%+v", expression.Value)
}
