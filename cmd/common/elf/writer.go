package elf

import (
	"encoding/binary"
	"io"
)

type ElfWriter struct {
	Writer io.Writer
	Length uint64
}

func GetAlignedAddress(address uint64, boundary uint64) uint64 {
	return address + GetAlignedShift(address, boundary)
}

func GetAlignedShift(address uint64, boundary uint64) uint64 {
	if address%boundary == 0 {
		return 0
	}

	return boundary - (address % boundary)
}

func NewElfWriter(writer io.Writer) *ElfWriter {
	return &ElfWriter{
		Writer: writer,
		Length: 0,
	}
}

func (writer *ElfWriter) Write(data any) {
	buf, _ := binary.Append([]byte{}, binary.LittleEndian, data)
	writer.Writer.Write(buf)
	writer.Length += uint64(len(buf))
}

func (writer *ElfWriter) Align(boundary uint64) {
	buf := make([]byte, GetAlignedShift(writer.Length, boundary))
	writer.Writer.Write(buf)
	writer.Length += uint64(len(buf))
}
