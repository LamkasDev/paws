package lexer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Lexer struct {
	Tokens   []*LexerToken
	File     []rune
	Position int
}

func NewLexer() *Lexer {
	return &Lexer{
		Tokens: []*LexerToken{},
	}
}

func (lexer *Lexer) Process(file string) {
	lexer.File = []rune(strings.ReplaceAll(string(file), "\n", ""))
	lexer.Position = 0
	for !lexer.IsDone() {
		token := lexer.GetToken()
		if token == nil {
			continue
		}
		lexer.Tokens = append(lexer.Tokens, token)
	}
}

func (lexer *Lexer) GetToken() *LexerToken {
	c := lexer.GetNextCharacter()
	switch c {
	case '+', '-', '*', '/', '=', ';', '{', '}':
		return &LexerToken{
			Type:  LexerTokenMap[string(c)],
			Value: string(c),
		}
	case ' ':
		return nil
	case '"':
		str := lexer.GetNextString()
		_ = lexer.GetNextCharacter()
		return &LexerToken{
			Type:  LexerTokenString,
			Value: str,
		}
	default:
		str := lexer.GetNextIdentifier(c)
		if keywordType, ok := LexerTokenKeywords[str]; ok {
			return &LexerToken{
				Type:  keywordType,
				Value: str,
			}
		}

		if num, err := strconv.Atoi(str); err == nil {
			return &LexerToken{
				Type:  LexerTokenNumber,
				Value: num,
			}
		}

		return &LexerToken{
			Type:  LexerTokenIdentifier,
			Value: str,
		}
	}

	return nil
}

func (lexer *Lexer) Print() {
	for _, token := range lexer.Tokens {
		fmt.Printf("[%s] ", fmt.Sprint(token.Value))
	}
	fmt.Printf("\n")
}

func (lexer *Lexer) GetNextString() string {
	s := ""
	for !lexer.IsDone() && lexer.File[lexer.Position] != '"' {
		s += string(lexer.GetNextCharacter())
	}

	return s
}

func (lexer *Lexer) GetNextIdentifier(c rune) string {
	s := string(c)
	for !lexer.IsDone() && unicode.IsLetter(lexer.File[lexer.Position]) {
		s += string(lexer.GetNextCharacter())
	}

	return s
}

func (lexer *Lexer) GetNextCharacter() rune {
	c := lexer.File[lexer.Position]
	lexer.Position++

	return c
}

func (lexer *Lexer) IsDone() bool {
	return lexer.Position >= len(lexer.File)
}
