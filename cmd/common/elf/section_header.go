package elf

import (
	"unsafe"
)

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

const ElfSectionHeaderSize = uint16(unsafe.Sizeof(ElfSectionHeader{}))

func NewElfSectionHeaderNull() *ElfSectionHeader {
	return &ElfSectionHeader{
		AddressAlign: 1,
	}
}

func NewElfSectionHeaderProgram(name uint32, address uint64, offset uint64, size uint64) *ElfSectionHeader {
	return &ElfSectionHeader{
		Name:         name,
		Type:         1,
		Flags:        7,
		Address:      address,
		Offset:       GetAlignedAddress(offset, 16),
		Size:         size,
		AddressAlign: 16,
	}
}

func NewElfSectionHeaderSymbolTable(name uint32, offset uint64, size uint64) *ElfSectionHeader {
	return &ElfSectionHeader{
		Name:         name,
		Type:         2,
		Offset:       GetAlignedAddress(offset, 8),
		Size:         size,
		Link:         3,
		Info:         2,
		AddressAlign: 8,
		EntrySize:    24,
	}
}

func NewElfSectionHeaderStringTable(name uint32, offset uint64, size uint64) *ElfSectionHeader {
	return &ElfSectionHeader{
		Name:         name,
		Type:         3,
		Offset:       GetAlignedAddress(offset, 1),
		Size:         size,
		AddressAlign: 1,
	}
}

func (data *ElfSectionHeader) WriteTo(w *ElfWriter) {
	w.Write(data.Name)
	w.Write(data.Type)
	w.Write(data.Flags)
	w.Write(data.Address)
	w.Write(data.Offset)
	w.Write(data.Size)
	w.Write(data.Link)
	w.Write(data.Info)
	w.Write(data.AddressAlign)
	w.Write(data.EntrySize)
}
