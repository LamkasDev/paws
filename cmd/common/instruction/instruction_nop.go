package instruction

func NewInstructionNop() Instruction {
	ins := Instruction{
		0x90,
	}

	return ins
}
