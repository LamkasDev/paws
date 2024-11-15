package main

import (
	"fmt"
	"os"

	"github.com/LamkasDev/paws/cmd/common/compiler"
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/lexer"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

func main() {
	input, _ := os.ReadFile("../../data/main.txt")

	lexerc := lexer.NewLexer()
	lexerc.Process(string(input))
	lexerc.Print()

	parserc := parser.NewParser()
	parserc.Process(lexerc)
	parserc.Print()

	compilerc := compiler.NewCompiler()
	compilerc.Process(parserc)
	compilerc.PostProcess()

	f, err := os.OpenFile("main", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := elf.NewElfWriter(f)

	elf.NewElf(*compilerc.Data).WriteTo(writer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("done :3")
}
