package compiler

import (
	"fmt"

	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

func (compiler *Compiler) CompileCall(expression *parser.ParserExpressionCall, section *elf.ElfProgramSection) {
	switch expression.Symbol.Name {
	case "print":
		compiler.AddSyscallPrint(section, "str")
	default:
		functionSection := compiler.Data.FindSection(expression.Symbol.Name)
		if functionSection == nil {
			panic(fmt.Sprintf("unknown function %s\n", expression.Symbol.Name))
		}
		instruction.NewInstructionCall(uint32(functionSection.Address - (section.Address + uint64(len(section.Data)+5)))).WriteTo(section)
	}
}
