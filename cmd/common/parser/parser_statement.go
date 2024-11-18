package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionStatement() *ParserExpression {
	token := parser.GetNextToken()
	switch token.Type {
	case lexer.LexerTokenIdentifier:
		symbol := parser.Scope.FindSymbol(token.Value.(string))
		if symbol == nil {
			if parser.MatchToken(lexer.LexerTokenEq) == nil {
				return nil
			}

			value := parser.GetExpressionValue()
			if value == nil {
				return nil
			}
			valueSymbol := NewParserSymbol(token.Value.(string), ParserSymbolVariable)
			parser.Scope.AddSymbol(valueSymbol)

			if parser.MatchToken(lexer.LexerTokenSemicolon) == nil {
				return nil
			}

			return NewParserExpressionAssignment(valueSymbol, value)
		}
		if parser.MatchToken(lexer.LexerTokenLeftBracket) == nil {
			return nil
		}
		if parser.MatchToken(lexer.LexerTokenRightBracket) == nil {
			return nil
		}
		if parser.MatchToken(lexer.LexerTokenSemicolon) == nil {
			return nil
		}

		return NewParserExpressionCall(symbol)
	default:
		// TODO: error handling
		return nil
	}

	return nil
}
