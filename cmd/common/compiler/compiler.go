package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/cpu"
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

type Compiler struct {
	Data *elf.ElfProgram

	Parser   *parser.Parser
	Position int
}

func NewCompiler() *Compiler {
	return &Compiler{
		Data: &elf.ElfProgram{
			Sections: []*elf.ElfProgramSection{
				{
					Name:  "world.c",
					Align: 1,
				},
				{
					Name: "myprint",
					Data: []byte{
						0xF3, 0x0F, 0x1E, 0xFA, 0x48, 0x8B, 0x0D, 0x7D, 0x00, 0x00, 0x00, 0x48, 0xC7, 0xC0, 0x01, 0x00,
						0x00, 0x00, 0x48, 0xC7, 0xC7, 0x01, 0x00, 0x00, 0x00, 0x48, 0x89, 0xCE, 0x48, 0xC7, 0xC2, 0x0D,
						0x00, 0x00, 0x00, 0x0F, 0x05, 0xC3,
					},
					Align: 16,
				},
				{
					Name: "myexit",
					Data: []byte{
						0xF3, 0x0F, 0x1E, 0xFA, 0x48, 0xC7, 0xC0, 0x3C, 0x00, 0x00, 0x00, 0x48, 0x31, 0xFF, 0x0F, 0x05,
						0xC3,
					},
					Align: 16,
				},
			},
		},
	}
}

func (compiler *Compiler) Process(parserc *parser.Parser) {
	compiler.Parser = parserc
	compiler.Position = 0
	for _, rawExpression := range compiler.Parser.Expressions {
		switch expression := rawExpression.Data.(type) {
		case *parser.ParserExpressionFunction:
			section := &elf.ElfProgramSection{
				Name:  expression.Name.Value.(string),
				Data:  []byte{},
				Align: 16,
			}
			instruction.NewInstructionEndBr64().WriteTo(section)
			instruction.NewInstructionMovMem32ToR32(0x2D, cpu.RegisterRcx).WriteTo(section)
			instruction.NewInstructionMovImm32ToR32(0x1, cpu.RegisterRax).WriteTo(section)
			instruction.NewInstructionMovImm32ToR32(0x1, cpu.RegisterRdi).WriteTo(section)
			instruction.NewInstructionMovR32ToR32(cpu.RegisterRcx, cpu.RegisterRsi).WriteTo(section)
			instruction.NewInstructionMovImm32ToR32(0xD, cpu.RegisterRdx).WriteTo(section)
			instruction.NewInstructionSyscall().WriteTo(section)
			instruction.NewInstructionMovImm32ToR32(0x3C, cpu.RegisterRax).WriteTo(section)
			instruction.NewInstructionXorR32ToR32(cpu.RegisterRdi, cpu.RegisterRdi).WriteTo(section)
			instruction.NewInstructionSyscall().WriteTo(section)
			instruction.NewInstructionReturn().WriteTo(section)
			compiler.Data.Sections = append(compiler.Data.Sections, section)
		case *parser.ParserExpressionAssignment:
			section := &elf.ElfProgramSection{
				Name: expression.Symbol.Name,
				Data: []byte{
					0x80, 0x81, 0x04, 0x08, 0x00, 0x00, 0x00, 0x00,
				},
				Align: 8,
			}
			section.Data = append(section.Data, []byte(expression.Value.Data.(*parser.ParserExpressionPrimitive).Value.(string))...)
			section.Data = append(section.Data, []byte("\n")...)
			section.Data = append(section.Data, 0x00)

			compiler.Data.Sections = append(compiler.Data.Sections, section)
		}
	}
}
