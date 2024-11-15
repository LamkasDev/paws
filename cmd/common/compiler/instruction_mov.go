package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
)

func (compiler *Compiler) AddInstructionMovMem32ToR32(section *elf.ElfProgramSection, name string, dst uint8) {
	compiler.AddPostProcessEntry(section, PostProcessEntrySectionAddress, name)
	instruction.NewInstructionMovMem32ToR32(0x00, dst).WriteTo(section)
}
