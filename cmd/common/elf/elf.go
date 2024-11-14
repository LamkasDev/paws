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
			Entry:                    0x8048140,
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

	AddElfSectionHeaders(&elf, program)

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
	code := program.Encode()
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

func AddElfSectionHeaders(elf *Elf, program ElfProgram) {
	AddElfSectionHeader(elf, NewElfSectionHeaderNull(), []byte{})
	AddElfSectionHeader(elf, NewElfSectionHeaderProgram(27, elf.ProgramHeaders[1].VirtualAddress, elf.ProgramHeaders[1].Offset, elf.ProgramHeaders[1].MemorySize), []byte{})
	symbolTable := ElfSymbolTable{
		Symbol{},
		Symbol{
			Name:               1,
			Flags:              0x4,
			SectionHeaderIndex: 65521,
		},
		Symbol{
			Name:               9,
			Flags:              0x12,
			SectionHeaderIndex: 1,
			Value:              134512880,
			Size:               38,
		},
		Symbol{
			Name:               17,
			Flags:              0x12,
			SectionHeaderIndex: 1,
			Value:              134512928,
			Size:               17,
		},
		Symbol{
			Name:               24,
			Flags:              0x12,
			SectionHeaderIndex: 1,
			Value:              134512960,
			Size:               50,
		},
		Symbol{
			Name:               31,
			Flags:              0x11,
			SectionHeaderIndex: 1,
			Value:              134513016,
			Size:               8,
		},
	}
	symbolTableData := symbolTable.Encode()
	AddElfSectionHeader(elf, NewElfSectionHeaderSymbolTable(1, elf.SectionHeaders[1].Offset+elf.SectionHeaders[1].Size, uint64(len(symbolTableData))), symbolTableData)
	stringTable := ElfStringtable{
		"",
	}
	for _, section := range program.Sections {
		stringTable = append(stringTable, section.Name)
	}
	stringTableData := stringTable.Encode()
	AddElfSectionHeader(elf, NewElfSectionHeaderStringTable(9, elf.SectionHeaders[2].Offset+elf.SectionHeaders[2].Size, uint64(len(stringTableData))), stringTableData)
	stringTable = ElfStringtable{
		"", ".symtab", ".strtab", ".shstrtab", "tiny",
	}
	stringTableData = stringTable.Encode()
	AddElfSectionHeader(elf, NewElfSectionHeaderStringTable(17, elf.SectionHeaders[3].Offset+elf.SectionHeaders[3].Size, uint64(len(stringTableData))), stringTableData)
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

func (data Elf) WriteTo(w *ElfWriter) {
	data.Header.WriteTo(w)
	for _, header := range data.ProgramHeaders {
		header.WriteTo(w)
	}
	w.Align(16)
	for _, data := range data.Data {
		w.Write(data)
	}
	w.Align(8)
	for _, header := range data.SectionHeaders {
		header.WriteTo(w)
	}
}
