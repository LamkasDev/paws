package elf

import (
	"encoding/binary"
)

type Symbol struct {
	Name               uint32
	Flags              uint8
	Other              uint8
	SectionHeaderIndex uint16
	Value              uint64
	Size               uint64
}

type ElfSymbolTable []Symbol

func EncodeElfSymbolTable(data ElfSymbolTable) []byte {
	buf := []byte{}
	for _, symbol := range data {
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.Name)
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.Flags)
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.Other)
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.SectionHeaderIndex)
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.Value)
		buf, _ = binary.Append(buf, binary.LittleEndian, symbol.Size)
	}

	return buf
}
