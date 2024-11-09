package elf

import "bytes"

type ElfStringtable []string

func EncodeElfStringtable(data ElfStringtable) []byte {
	var buf bytes.Buffer
	for _, str := range data {
		buf.Write([]byte(str))
		buf.Write([]byte{0})
	}

	return buf.Bytes()
}
