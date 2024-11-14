package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionStatement() *ParserExpression {
	token := parser.GetNextToken()
	switch token.Type {
	case lexer.LexerTokenIdentifier:
		if parser.MatchToken(lexer.LexerTokenEq) == nil {
			return nil
		}

		value := parser.GetExpressionValue()
		if value == nil {
			return nil
		}

		symbol := NewParserSymbol(token.Value.(string), ParserSymbolInt)
		parser.Scope.AddSymbol(symbol)

		if parser.MatchToken(lexer.LexerTokenSemicolon) == nil {
			return nil
		}

		return NewParserExpressionAssignment(symbol, value)
	default:
		// TODO: error handling
		return nil
	}

	return nil
}
