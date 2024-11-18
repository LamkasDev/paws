package instruction

func NewInstructionPopR32(reg uint8) Instruction {
	ins := Instruction{
		0x58 | reg,
	}

	return ins
}
