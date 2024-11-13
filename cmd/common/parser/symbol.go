package parser

const ParserSymbolInt = uint16(0)

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
