package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionValue() *ParserExpression {
	token := parser.GetNextToken()
	switch token.Type {
	case lexer.LexerTokenNumber:
		// TODO: correct order of operations
		expression := NewParserExpressionPrimitive(token.Value)
		if nextToken := parser.PeekNextToken(); nextToken != nil && nextToken.IsOperator() {
			_ = parser.GetNextToken()
			expression = NewParserExpressionArithmetic(expression, parser.GetExpressionValue(), nextToken)
		}
		return expression
	default:
		// TODO: error handling
		return nil
	}

	return nil
}
