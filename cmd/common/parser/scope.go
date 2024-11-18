package parser

type ParserScope struct {
	Name    string
	Symbols []*ParserSymbol
	Parent  *ParserScope
}

func NewParserScope(name string, parent *ParserScope) *ParserScope {
	return &ParserScope{
		Name:    name,
		Symbols: []*ParserSymbol{},
		Parent:  parent,
	}
}

func (scope *ParserScope) AddSymbol(symbol *ParserSymbol) {
	scope.Symbols = append(scope.Symbols, symbol)
}

func (scope *ParserScope) FindSymbol(name string) *ParserSymbol {
	for _, symbol := range scope.Symbols {
		if symbol.Name == name {
			return symbol
		}
	}
	if scope.Parent != nil {
		return scope.Parent.FindSymbol(name)
	}

	return nil
}
