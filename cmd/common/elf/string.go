package elf

import "bytes"

type ElfStringtable []string

func (data ElfStringtable) Encode() []byte {
	var buf bytes.Buffer
	for _, str := range data {
		buf.Write([]byte(str))
		buf.Write([]byte{0})
	}

	return buf.Bytes()
}
