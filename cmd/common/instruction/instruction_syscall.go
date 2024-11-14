package instruction

func NewInstructionSyscall() Instruction {
	return Instruction{
		0x0F, 0x05,
	}
}
