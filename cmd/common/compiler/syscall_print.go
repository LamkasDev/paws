package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/cpu"
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
)

func (compiler *Compiler) AddSyscallPrint(section *elf.ElfProgramSection, name string) {
	instruction.NewInstructionMovImm32ToR32(0x01, cpu.RegisterSyscallNumber).WriteTo(section)
	instruction.NewInstructionMovImm32ToR32(0x01, cpu.RegisterSyscallArg0).WriteTo(section)
	compiler.AddInstructionMovMem32ToR32(section, name, cpu.RegisterSyscallArg1)
	compiler.AddPostProcessEntry(section, PostProcessEntrySectionStringSize, name)
	instruction.NewInstructionMovImm32ToR32(0xD, cpu.RegisterSyscallArg2).WriteTo(section)
	instruction.NewInstructionSyscall().WriteTo(section)
}
