package elf

import (
	"encoding/binary"
	"io"
)

type Elf struct {
	Identification ElfIdentification
	Header         ElfHeader
	ProgramHeader  ElfProgramHeader
	SectionHeader  *ElfSectionHeader
}

type ElfIdentification struct {
	Magic      string
	Class      uint8
	Data       uint8
	Version    uint8
	OsAbi      uint8
	AbiVersion uint8
}

type ElfHeader struct {
	Type                     uint16
	Machine                  uint16
	Version                  uint32
	Entry                    uint64
	ProgramHeaderOffset      uint64
	SectionHeaderOffset      uint64
	Flags                    uint32
	HeaderSize               uint16
	ProgramHeaderEntrySize   uint16
	ProgramHeaderEntries     uint16
	SectionHeaderEntrySize   uint16
	SectionHeaderEntries     uint16
	SectionHeaderStringIndex uint16
}

type ElfProgramHeader struct {
	Type            uint32
	Flags           uint32
	Offset          uint64
	VirtualAddress  uint64
	PhysicalAddress uint64
	FileSize        uint64
	MemorySize      uint64
	Align           uint64
	Data            []byte
}

type ElfSectionHeader struct {
	Name         uint32
	Type         uint32
	Flags        uint64
	Address      uint64
	Offset       uint64
	Size         uint64
	Link         uint32
	Info         uint32
	AddressAlign uint64
	EntrySize    uint64
}

func NewElf() Elf {
	program := []byte{
		0xB8, 0x01, 0x00, 0x00, 0x00, 0xBF, 0x01, 0x00, 0x00, 0x00, 0x48, 0xBE, 0x9D, 0x00, 0x40, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xBA, 0x0D, 0x00, 0x00, 0x00, 0x0F, 0x05, 0xB8, 0x3C, 0x00, 0x00, 0x00,
		0x48, 0x31, 0xFF, 0x0F, 0x05,
	}
	text := "hall\n"

	elf := Elf{
		Identification: ElfIdentification{
			Magic:      "\x7FELF",
			Class:      2,
			Data:       1,
			Version:    1,
			OsAbi:      0,
			AbiVersion: 0,
		},
		Header: ElfHeader{
			Type:                     2,
			Machine:                  0x3E,
			Version:                  1,
			Entry:                    4194424,
			ProgramHeaderOffset:      64,
			SectionHeaderOffset:      0,
			Flags:                    0,
			HeaderSize:               64,
			ProgramHeaderEntrySize:   56,
			ProgramHeaderEntries:     1,
			SectionHeaderEntrySize:   0,
			SectionHeaderEntries:     0,
			SectionHeaderStringIndex: 0,
		},
		ProgramHeader: ElfProgramHeader{
			Type:            1,
			Flags:           5,
			Offset:          0,
			VirtualAddress:  4194304,
			PhysicalAddress: 4194304,
			Align:           4096,
			Data:            program,
		},
	}
	elf.ProgramHeader.Data = append(elf.ProgramHeader.Data, []byte(text)...)
	elf.ProgramHeader.FileSize = 121 + uint64(len(elf.ProgramHeader.Data))
	elf.ProgramHeader.MemorySize = elf.ProgramHeader.FileSize

	return elf
}

func EncodeElfIdentification(w io.Writer, data ElfIdentification) {
	binary.Write(w, binary.LittleEndian, []byte(data.Magic))
	binary.Write(w, binary.LittleEndian, data.Class)
	binary.Write(w, binary.LittleEndian, data.Data)
	binary.Write(w, binary.LittleEndian, data.Version)
	binary.Write(w, binary.LittleEndian, data.OsAbi)
	binary.Write(w, binary.LittleEndian, data.AbiVersion)
	binary.Write(w, binary.LittleEndian, [7]byte{})
}

func EncodeElfHeader(w io.Writer, data ElfHeader) {
	binary.Write(w, binary.LittleEndian, data.Type)
	binary.Write(w, binary.LittleEndian, data.Machine)
	binary.Write(w, binary.LittleEndian, data.Version)
	binary.Write(w, binary.LittleEndian, data.Entry)
	binary.Write(w, binary.LittleEndian, data.ProgramHeaderOffset)
	binary.Write(w, binary.LittleEndian, data.SectionHeaderOffset)
	binary.Write(w, binary.LittleEndian, data.Flags)
	binary.Write(w, binary.LittleEndian, data.HeaderSize)
	binary.Write(w, binary.LittleEndian, data.ProgramHeaderEntrySize)
	binary.Write(w, binary.LittleEndian, data.ProgramHeaderEntries)
	binary.Write(w, binary.LittleEndian, data.SectionHeaderEntrySize)
	binary.Write(w, binary.LittleEndian, data.SectionHeaderEntries)
	binary.Write(w, binary.LittleEndian, data.SectionHeaderStringIndex)
}

func EncodeElfProgramHeader(w io.Writer, data ElfProgramHeader) {
	binary.Write(w, binary.LittleEndian, data.Type)
	binary.Write(w, binary.LittleEndian, data.Flags)
	binary.Write(w, binary.LittleEndian, data.Offset)
	binary.Write(w, binary.LittleEndian, data.VirtualAddress)
	binary.Write(w, binary.LittleEndian, data.PhysicalAddress)
	binary.Write(w, binary.LittleEndian, data.FileSize)
	binary.Write(w, binary.LittleEndian, data.MemorySize)
	binary.Write(w, binary.LittleEndian, data.Align)
	binary.Write(w, binary.LittleEndian, data.Data)
}

func EncodeElf(w io.Writer, data Elf) {
	EncodeElfIdentification(w, data.Identification)
	EncodeElfHeader(w, data.Header)
	EncodeElfProgramHeader(w, data.ProgramHeader)
}
