package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionFunction() *ParserExpression {
	if parser.MatchToken(lexer.LexerTokenFunction) == nil {
		return nil
	}
	name := parser.MatchToken(lexer.LexerTokenIdentifier)
	if name == nil {
		return nil
	}
	if parser.MatchToken(lexer.LexerTokenLeftCurly) == nil {
		return nil
	}

	fn := NewParserExpressionFunction(name)
	if parser.MatchToken(lexer.LexerTokenRightCurly) != nil {
		return fn
	}

	for !parser.IsDone() {
		expression := parser.GetExpressionStatement()
		if expression == nil {
			panic("couldn't create statement")
			continue
		}
		fn.Data.(*ParserExpressionFunction).Statements = append(fn.Data.(*ParserExpressionFunction).Statements, expression)
		if parser.MatchToken(lexer.LexerTokenRightCurly) != nil {
			return fn
		}
	}

	return fn
}
