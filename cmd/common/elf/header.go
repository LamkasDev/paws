package elf

import (
	"unsafe"
)

type ElfHeader struct {
	Identification           ElfIdentification
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

const ElfHeaderSize = uint16(unsafe.Sizeof(ElfHeader{}))

func EncodeElfHeader(w *ElfWriter, data ElfHeader) {
	EncodeElfIdentification(w, data.Identification)
	w.Write(data.Type)
	w.Write(data.Machine)
	w.Write(data.Version)
	w.Write(data.Entry)
	w.Write(data.ProgramHeaderOffset)
	w.Write(data.SectionHeaderOffset)
	w.Write(data.Flags)
	w.Write(data.HeaderSize)
	w.Write(data.ProgramHeaderEntrySize)
	w.Write(data.ProgramHeaderEntries)
	w.Write(data.SectionHeaderEntrySize)
	w.Write(data.SectionHeaderEntries)
	w.Write(data.SectionHeaderStringIndex)
}
