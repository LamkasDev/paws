package parser

const ParserExpressionTypeAssignment = uint16(0)
const ParserExpressionTypePrimitive = uint16(1)
const ParserExpressionTypeArithmetic = uint16(2)
const ParserExpressionTypeFunction = uint16(2)

type ParserExpression struct {
	Type uint16
	Data interface{}
}

func (rawExpression *ParserExpression) Sprint() string {
	switch expression := rawExpression.Data.(type) {
	case *ParserExpressionAssignment:
		return expression.Sprint()
	case *ParserExpressionPrimitive:
		return expression.Sprint()
	case *ParserExpressionArithmetic:
		return expression.Sprint()
	case *ParserExpressionFunction:
		return expression.Sprint()
	}

	return ""
}
