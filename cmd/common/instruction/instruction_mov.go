package instruction

import (
	"encoding/binary"
)

func NewInstructionMovR32ToR32(src uint8, dst uint8) Instruction {
	ins := Instruction{
		0x48, 0x89,
	}
	ins = append(ins, 0xC0|(src<<3)|dst)

	return ins
}

func NewInstructionMovImm32ToR32(imm uint32, dst uint8) Instruction {
	ins := Instruction{
		0x48, 0xC7,
	}
	ins = append(ins, 0xC0|dst)
	ins, _ = binary.Append(ins, binary.LittleEndian, imm)

	return ins
}

func NewInstructionMovMem32ToR32(address uint32, dst uint8) Instruction {
	// TODO: actual adresssing (RIP-relative, make sure SIB byte doesn't fuck shit up)
	ins := Instruction{
		0x48, 0x8B,
	}
	ins = append(ins, 0x05|dst<<3)
	ins, _ = binary.Append(ins, binary.LittleEndian, address)

	return ins
}
