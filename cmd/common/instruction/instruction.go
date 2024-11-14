package instruction

import (
	"github.com/LamkasDev/paws/cmd/common/elf"
)

type Instruction []byte

func (instruction Instruction) WriteTo(section *elf.ElfProgramSection) {
	section.Data = append(section.Data, instruction...)
}
