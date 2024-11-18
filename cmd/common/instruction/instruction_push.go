package instruction

func NewInstructionPushR32(reg uint8) Instruction {
	ins := Instruction{
		0x50 | reg,
	}

	return ins
}
