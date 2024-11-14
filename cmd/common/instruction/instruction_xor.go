package instruction

func NewInstructionXorR32ToR32(src uint8, dst uint8) Instruction {
	ins := Instruction{
		0x48, 0x31,
	}
	ins = append(ins, 0xC0|(src<<3)|dst)

	return ins
}
