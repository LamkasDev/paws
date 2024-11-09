package elf

type ElfProgram struct {
	Sections []ElfProgramSection
}

type ElfProgramSection struct {
	Name  string
	Data  []byte
	Align uint64
}

func EncodeElfProgramData(program ElfProgram) []byte {
	buf := []byte{}
	for _, section := range program.Sections {
		align := make([]byte, GetAlignedShift(uint64(len(buf)), section.Align))
		buf = append(buf, align...)
		buf = append(buf, section.Data...)
	}

	return buf
}
