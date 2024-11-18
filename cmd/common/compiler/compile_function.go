package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/cpu"
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

func (compiler *Compiler) CompileFunction(expression *parser.ParserExpressionFunction) {
	section := compiler.CreateSection(elf.ElfProgramSectionFunction, expression.Name.Value.(string), 8)
	instruction.NewInstructionEndBr64().WriteTo(section)
	instruction.NewInstructionPushR32(cpu.RegisterRbp).WriteTo(section)
	instruction.NewInstructionMovR32ToR32(cpu.RegisterRsp, cpu.RegisterRbp).WriteTo(section)
	for _, rawStatement := range expression.Statements {
		switch statement := rawStatement.Data.(type) {
		case *parser.ParserExpressionCall:
			compiler.CompileCall(statement, section)
		}
	}
	if expression.Name.Value.(string) == "main" {
		compiler.AddSyscallSemop(section)
	}
	instruction.NewInstructionMovImm32ToR32(0x00, cpu.RegisterRax).WriteTo(section)
	instruction.NewInstructionPopR32(cpu.RegisterRbp).WriteTo(section)
	instruction.NewInstructionReturn().WriteTo(section)
	compiler.AddSection(section)
}
