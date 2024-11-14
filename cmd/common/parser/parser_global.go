package parser

import "github.com/LamkasDev/paws/cmd/common/lexer"

func (parser *Parser) GetExpressionGlobal() *ParserExpression {
	token := parser.PeekNextToken()
	switch token.Type {
	case lexer.LexerTokenFunction:
		return parser.GetExpressionFunction()
	}

	return parser.GetExpressionStatement()
}
