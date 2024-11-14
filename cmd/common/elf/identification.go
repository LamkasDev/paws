package elf

type ElfIdentification struct {
	Magic      [4]byte
	Class      uint8
	Data       uint8
	Version    uint8
	OsAbi      uint8
	AbiVersion uint8
}

func (data ElfIdentification) WriteTo(w *ElfWriter) {
	w.Write(data.Magic)
	w.Write(data.Class)
	w.Write(data.Data)
	w.Write(data.Version)
	w.Write(data.OsAbi)
	w.Write(data.AbiVersion)
	w.Write([7]byte{})
}
