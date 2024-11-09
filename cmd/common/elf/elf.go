package elf

type Elf struct {
	Header         ElfHeader
	ProgramHeaders []*ElfProgramHeader
	Data           []byte
	SectionHeaders []*ElfSectionHeader

	Offset uint64
}

func NewElf(program ElfProgram) Elf {
	elf := Elf{
		Header: ElfHeader{
			Identification: ElfIdentification{
				Magic:      [4]byte{0x7F, 0x45, 0x4C, 0x46},
				Class:      2,
				Data:       1,
				Version:    1,
				OsAbi:      0,
				AbiVersion: 0,
			},
			Type:                     2,
			Machine:                  0x3E,
			Version:                  1,
			Entry:                    134512960,
			Flags:                    0,
			HeaderSize:               ElfHeaderSize,
			ProgramHeaderEntrySize:   ElfProgramHeaderSize,
			SectionHeaderEntrySize:   ElfSectionHeaderSize,
			SectionHeaderStringIndex: 4,
		},
		ProgramHeaders: []*ElfProgramHeader{},
		SectionHeaders: []*ElfSectionHeader{},
		Offset:         uint64(ElfHeaderSize),
	}

	AddElfProgramHeaders(&elf, program)
	EndElfProgramHeaders(&elf)

	AddElfSectionHeaders(&elf)

	elf.Header.ProgramHeaderOffset = GetElfProgramHeadersStart(&elf)
	elf.Header.ProgramHeaderEntries = uint16(len(elf.ProgramHeaders))

	elf.Header.SectionHeaderOffset = GetElfSectionHeadersStart(&elf)
	elf.Header.SectionHeaderEntries = uint16(len(elf.SectionHeaders))

	return elf
}

func GetElfProgramHeadersStart(elf *Elf) uint64 {
	return uint64(ElfHeaderSize)
}

func GetElfProgramHeadersEnd(elf *Elf) uint64 {
	return GetElfProgramHeadersStart(elf) + uint64(ElfProgramHeaderSize)*uint64(len(elf.ProgramHeaders))
}

func GetElfProgramStart(elf *Elf) uint64 {
	return GetAlignedAddress(GetElfProgramHeadersEnd(elf), 16)
}

func GetElfSectionHeadersStart(elf *Elf) uint64 {
	return GetAlignedAddress(GetElfProgramStart(elf)+uint64(len(elf.Data)), 8)
}

func AddElfProgramHeaders(elf *Elf, program ElfProgram) {
	AddElfProgramHeader(elf, NewElfProgramHeaderLoadElf(), []byte{})
	code := EncodeElfProgramData(program)
	AddElfProgramHeader(elf, NewElfProgramHeaderLoadCode(uint64(len(code))), code)
	AddElfProgramHeader(elf, NewElfProgramHeaderStack(), []byte{})
}

func EndElfProgramHeaders(elf *Elf) {
	elf.ProgramHeaders[0].FileSize = GetElfProgramHeadersEnd(elf)
	elf.ProgramHeaders[0].MemorySize = elf.ProgramHeaders[0].FileSize

	elf.ProgramHeaders[1].Offset = GetElfProgramStart(elf)
	elf.ProgramHeaders[1].VirtualAddress = elf.ProgramHeaders[0].VirtualAddress + elf.ProgramHeaders[1].Offset
	elf.ProgramHeaders[1].PhysicalAddress = elf.ProgramHeaders[1].VirtualAddress
}

func AddElfSectionHeaders(elf *Elf) {
	AddElfSectionHeader(elf, NewElfSectionHeaderNull(), []byte{})
	AddElfSectionHeader(elf, NewElfSectionHeaderProgram(27, elf.ProgramHeaders[1].VirtualAddress, elf.ProgramHeaders[1].Offset, elf.ProgramHeaders[1].MemorySize), []byte{})
	symbolTable := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0xF1, 0xFF,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09, 0x00, 0x00, 0x00, 0x12, 0x00, 0x01, 0x00, 0xF0, 0x80, 0x04, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x26, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x12, 0x00, 0x01, 0x00,
		0x20, 0x81, 0x04, 0x08, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x18, 0x00, 0x00, 0x00, 0x12, 0x00, 0x01, 0x00, 0x40, 0x81, 0x04, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1F, 0x00, 0x00, 0x00, 0x11, 0x00, 0x01, 0x00,
		0x78, 0x81, 0x04, 0x08, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	AddElfSectionHeader(elf, NewElfSectionHeaderSymbolTable(1, elf.SectionHeaders[1].Offset+elf.SectionHeaders[1].Size, uint64(len(symbolTable))), symbolTable)
	stringTable := EncodeElfStringtable(ElfStringtable{
		"", "world.c", "myprint", "myexit", "nomain", "str",
	})
	AddElfSectionHeader(elf, NewElfSectionHeaderStringTable(9, elf.SectionHeaders[2].Offset+elf.SectionHeaders[2].Size, uint64(len(stringTable))), stringTable)
	stringTable = EncodeElfStringtable(ElfStringtable{
		"", ".symtab", ".strtab", ".shstrtab", "tiny",
	})
	AddElfSectionHeader(elf, NewElfSectionHeaderStringTable(17, elf.SectionHeaders[3].Offset+elf.SectionHeaders[3].Size, uint64(len(stringTable))), stringTable)
}

func AddElfProgramHeader(elf *Elf, header *ElfProgramHeader, data []byte) {
	align := make([]byte, GetAlignedShift(elf.Offset, 8))
	elf.Data = append(elf.Data, align...)
	elf.Offset += uint64(len(align))

	elf.ProgramHeaders = append(elf.ProgramHeaders, header)
	elf.Data = append(elf.Data, data...)
	elf.Offset += uint64(len(data))
}

func AddElfSectionHeader(elf *Elf, header *ElfSectionHeader, data []byte) {
	align := make([]byte, GetAlignedShift(elf.Offset, header.AddressAlign))
	elf.Data = append(elf.Data, align...)
	elf.Offset += uint64(len(align))

	elf.SectionHeaders = append(elf.SectionHeaders, header)
	elf.Data = append(elf.Data, data...)
	elf.Offset += uint64(len(data))
}

func EncodeElf(w *ElfWriter, data Elf) {
	EncodeElfHeader(w, data.Header)
	for _, header := range data.ProgramHeaders {
		EncodeElfProgramHeader(w, header)
	}
	w.Align(16)
	for _, data := range data.Data {
		w.Write(data)
	}
	w.Align(8)
	for _, header := range data.SectionHeaders {
		EncodeElfSectionHeader(w, header)
	}
}
