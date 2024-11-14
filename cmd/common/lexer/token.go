package lexer

const LexerTokenPlus = uint16(0)
const LexerTokenMinus = uint16(1)
const LexerTokenStar = uint16(2)
const LexerTokenSlash = uint16(3)
const LexerTokenIdentifier = uint16(4)
const LexerTokenEq = uint16(5)
const LexerTokenNumber = uint16(6)
const LexerTokenSemicolon = uint16(7)
const LexerTokenString = uint16(8)
const LexerTokenFunction = uint16(9)
const LexerTokenLeftCurly = uint16(10)
const LexerTokenRightCurly = uint16(11)

var LexerTokenMap = map[string]uint16{
	"+": LexerTokenPlus,
	"-": LexerTokenMinus,
	"*": LexerTokenStar,
	"/": LexerTokenSlash,
	"=": LexerTokenEq,
	";": LexerTokenSemicolon,
	"{": LexerTokenLeftCurly,
	"}": LexerTokenRightCurly,
}

var LexerTokenKeywords = map[string]uint16{
	"fn": LexerTokenFunction,
}

type LexerToken struct {
	Type  uint16
	Value interface{}
}

func (token *LexerToken) IsOperator() bool {
	return token.Type == LexerTokenPlus ||
		token.Type == LexerTokenMinus ||
		token.Type == LexerTokenSlash ||
		token.Type == LexerTokenStar
}
