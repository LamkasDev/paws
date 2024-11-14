package instruction

func NewInstructionEndBr64() Instruction {
	return Instruction{
		0xF3, 0x0F, 0x1E, 0xFA,
	}
}
