package cpu

const RegisterRax = uint8(0)
const RegisterRcx = uint8(1)
const RegisterRdx = uint8(2)
const RegisterRbx = uint8(3)
const RegisterRsp = uint8(4)
const RegisterRbp = uint8(5)
const RegisterRsi = uint8(6)
const RegisterRdi = uint8(7)
const RegisterR8 = uint8(8)
const RegisterR9 = uint8(9)
const RegisterR10 = uint8(10)
const RegisterR11 = uint8(11)
const RegisterR12 = uint8(12)
const RegisterR13 = uint8(13)
const RegisterR14 = uint8(14)
const RegisterR15 = uint8(15)

const RegisterSyscallNumber = RegisterRax
const RegisterSyscallArg0 = RegisterRdi
const RegisterSyscallArg1 = RegisterRsi
const RegisterSyscallArg2 = RegisterRdx
