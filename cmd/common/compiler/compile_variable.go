package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

func (compiler *Compiler) CompileVariable(expression *parser.ParserExpressionAssignment) {
	section := compiler.CreateSection(elf.ElfProgramSectionString, expression.Symbol.Name, 8)
	address := section.Address + 8
	section.Data = append(section.Data, byte(address))
	section.Data = append(section.Data, byte(address>>8))
	section.Data = append(section.Data, byte(address>>16))
	section.Data = append(section.Data, byte(address>>24))
	section.Data = append(section.Data, []byte{
		0x00, 0x00, 0x00, 0x00,
	}...)
	section.Data = append(section.Data, []byte(expression.Value.Data.(*parser.ParserExpressionPrimitive).Value.(string))...)
	section.Data = append(section.Data, []byte("\n")...)
	section.Data = append(section.Data, 0x00)
	compiler.AddSection(section)
}
