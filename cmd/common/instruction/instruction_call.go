package instruction

import (
	"encoding/binary"
)

func NewInstructionCall(address uint32) Instruction {
	ins := Instruction{
		0xE8,
	}
	ins, _ = binary.Append(ins, binary.LittleEndian, address)

	return ins
}
