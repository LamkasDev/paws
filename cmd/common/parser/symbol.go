package parser

const ParserSymbolVariable = uint16(0)
const ParserSymbolFunction = uint16(1)

type ParserSymbol struct {
	Name string
	Type uint16
}

func NewParserSymbol(name string, symbolType uint16) *ParserSymbol {
	return &ParserSymbol{
		Name: name,
		Type: symbolType,
	}
}
