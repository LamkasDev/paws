package elf

import (
	"unsafe"
)

type ElfProgramHeader struct {
	Type            uint32
	Flags           uint32
	Offset          uint64
	VirtualAddress  uint64
	PhysicalAddress uint64
	FileSize        uint64
	MemorySize      uint64
	Align           uint64
}

const ElfProgramHeaderSize = uint16(unsafe.Sizeof(ElfProgramHeader{}))

func NewElfProgramHeaderLoadElf() *ElfProgramHeader {
	return &ElfProgramHeader{
		Type:            1,
		Flags:           4,
		VirtualAddress:  134512640,
		PhysicalAddress: 134508544,
		Align:           4096,
	}
}

func NewElfProgramHeaderLoadCode(size uint64) *ElfProgramHeader {
	return &ElfProgramHeader{
		Type:       1,
		Flags:      7,
		FileSize:   size,
		MemorySize: size,
		Align:      4096,
	}
}

func NewElfProgramHeaderStack() *ElfProgramHeader {
	return &ElfProgramHeader{
		Type:  1685382481,
		Flags: 6,
		Align: 16,
	}
}

func (data *ElfProgramHeader) WriteTo(w *ElfWriter) {
	w.Write(data.Type)
	w.Write(data.Flags)
	w.Write(data.Offset)
	w.Write(data.VirtualAddress)
	w.Write(data.PhysicalAddress)
	w.Write(data.FileSize)
	w.Write(data.MemorySize)
	w.Write(data.Align)
}
