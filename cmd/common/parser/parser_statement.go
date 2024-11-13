package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionStatement() *ParserExpression {
	token := parser.GetNextToken()
	switch token.Type {
	case lexer.LexerTokenIdentifier:
		if nextToken := parser.PeekNextToken(); nextToken == nil || nextToken.Type != lexer.LexerTokenEq {
			return nil
		}
		_ = parser.GetNextToken()

		value := parser.GetExpressionValue()
		if value == nil {
			return nil
		}

		symbol := NewParserSymbol(token.Value.(string), ParserSymbolInt)
		parser.Scope.AddSymbol(symbol)

		if nextToken := parser.PeekNextToken(); nextToken == nil || nextToken.Type != lexer.LexerTokenSemicolon {
			return nil
		}
		_ = parser.GetNextToken()

		return NewParserExpressionAssignment(symbol, value)
	default:
		// TODO: error handling
		return nil
	}

	return nil
}
