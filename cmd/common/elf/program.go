package elf

const ElfProgramSectionNone = uint8(0)
const ElfProgramSectionFunction = uint8(1)
const ElfProgramSectionString = uint8(2)

type ElfProgram struct {
	Sections []*ElfProgramSection
}

type ElfProgramSection struct {
	Type uint8
	Name string
	Data []byte

	Address uint64
	Align   uint64
}

func (program ElfProgram) FindSection(name string) *ElfProgramSection {
	for _, section := range program.Sections {
		if section.Name == name {
			return section
		}
	}

	return nil
}

func (program ElfProgram) Encode() []byte {
	buf := []byte{}
	for _, section := range program.Sections {
		align := make([]byte, GetAlignedShift(uint64(len(buf)), section.Align))
		buf = append(buf, align...)
		buf = append(buf, section.Data...)
	}

	return buf
}
