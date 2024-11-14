package parser

import (
	"fmt"

	"github.com/LamkasDev/paws/cmd/common/lexer"
)

type Parser struct {
	Expressions []*ParserExpression
	Scope       *ParserScope

	Lexer    *lexer.Lexer
	Position int
}

func NewParser() *Parser {
	return &Parser{
		Expressions: []*ParserExpression{},
		Scope:       NewParserScope(nil),
	}
}

func (parser *Parser) Process(lexer *lexer.Lexer) {
	parser.Lexer = lexer
	parser.Position = 0
	for !parser.IsDone() {
		expression := parser.GetExpressionGlobal()
		if expression == nil {
			panic("couldn't create statement")
			continue
		}
		parser.Expressions = append(parser.Expressions, expression)
	}
}

func (parser *Parser) MatchToken(tokenType uint16) *lexer.LexerToken {
	if nextToken := parser.PeekNextToken(); nextToken == nil || nextToken.Type != tokenType {
		return nil
	}

	return parser.GetNextToken()
}

func (parser *Parser) PeekNextToken() *lexer.LexerToken {
	if parser.IsDone() {
		return nil
	}

	return parser.Lexer.Tokens[parser.Position]
}

func (parser *Parser) GetNextToken() *lexer.LexerToken {
	token := parser.Lexer.Tokens[parser.Position]
	parser.Position++

	return token
}

func (parser *Parser) Print() {
	for _, expression := range parser.Expressions {
		fmt.Printf("[%s]\n", expression.Sprint())
	}
	fmt.Printf("\n")
}

func (parser *Parser) IsDone() bool {
	return parser.Position >= len(parser.Lexer.Tokens)
}
