package parser

type ParserScope struct {
	Symbols  []*ParserSymbol
	Children []*ParserScope
	Parent   *ParserScope
}

func NewParserScope(parent *ParserScope) *ParserScope {
	return &ParserScope{
		Symbols:  []*ParserSymbol{},
		Children: []*ParserScope{},
		Parent:   parent,
	}
}

func (scope *ParserScope) AddScope(child *ParserScope) {
	scope.Children = append(scope.Children, child)
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
