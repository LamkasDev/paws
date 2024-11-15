package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/cpu"
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
)

func (compiler *Compiler) AddSyscallSemop(section *elf.ElfProgramSection) {
	instruction.NewInstructionMovImm32ToR32(0x3c, cpu.RegisterSyscallNumber).WriteTo(section)
	instruction.NewInstructionMovImm32ToR32(0x00, cpu.RegisterSyscallArg0).WriteTo(section)
	instruction.NewInstructionSyscall().WriteTo(section)
}
