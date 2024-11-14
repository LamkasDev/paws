package instruction

func NewInstructionReturn() Instruction {
	return Instruction{
		0xC3,
	}
}
