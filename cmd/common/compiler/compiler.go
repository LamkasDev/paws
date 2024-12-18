package compiler

import (
	"github.com/LamkasDev/paws/cmd/common/elf"
	"github.com/LamkasDev/paws/cmd/common/instruction"
	"github.com/LamkasDev/paws/cmd/common/parser"
)

type Compiler struct {
	Data               *elf.ElfProgram
	PostProcessEntries []*PostProcessEntry
	Address            uint64

	Parser   *parser.Parser
	Position int
}

func NewCompiler() *Compiler {
	compiler := &Compiler{
		Data: &elf.ElfProgram{
			Sections: []*elf.ElfProgramSection{},
		},
		PostProcessEntries: []*PostProcessEntry{},
		Address:            elf.GetAlignedAddress(0x08048000+uint64(elf.ElfHeaderSize)+uint64(3*elf.ElfProgramHeaderSize), 16),
	}

	section := compiler.CreateSection(elf.ElfProgramSectionNone, "main.c", 1)
	compiler.AddSection(section)

	return compiler
}

func (compiler *Compiler) Process(parserc *parser.Parser) {
	compiler.Parser = parserc
	compiler.Position = 0
	for _, rawExpression := range compiler.Parser.Expressions {
		switch expression := rawExpression.Data.(type) {
		case *parser.ParserExpressionFunction:
			compiler.CompileFunction(expression)
		case *parser.ParserExpressionAssignment:
			compiler.CompileVariable(expression)
		}
	}
}

func (compiler *Compiler) CreateSection(sectionType uint8, name string, align uint64) *elf.ElfProgramSection {
	section := &elf.ElfProgramSection{
		Type:    sectionType,
		Name:    name,
		Data:    []byte{},
		Address: elf.GetAlignedAddress(compiler.Address, align),
		Align:   align,
	}
	compiler.Address = section.Address

	return section
}

func (compiler *Compiler) AddSection(section *elf.ElfProgramSection) {
	compiler.Data.Sections = append(compiler.Data.Sections, section)
	compiler.Address += uint64(len(section.Data))
}

func (compiler *Compiler) AddPostProcessEntry(section *elf.ElfProgramSection, entryType uint16, target string) {
	postProcessEntry := NewPostProcessEntry(entryType, section.Name, uint64(len(section.Data)), target)
	compiler.PostProcessEntries = append(compiler.PostProcessEntries, postProcessEntry)
}

func (compiler *Compiler) PostProcess() {
	for _, postProcessEntry := range compiler.PostProcessEntries {
		section := compiler.Data.FindSection(postProcessEntry.Section)
		targetSection := compiler.Data.FindSection(postProcessEntry.Target)
		switch postProcessEntry.Type {
		case PostProcessEntrySectionAddress:
			(instruction.Instruction)(section.Data[postProcessEntry.Offset:]).EditInstructionMovMem32ToR32(int32(targetSection.Address) - int32(section.Address+postProcessEntry.Offset+7))
		case PostProcessEntrySectionStringSize:
			(instruction.Instruction)(section.Data[postProcessEntry.Offset:]).EditInstructionMovImm32ToR32(uint32(len(targetSection.Data) - 8))
		}
	}
}
