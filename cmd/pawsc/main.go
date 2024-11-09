package main

import (
	"fmt"
	"os"

	"github.com/LamkasDev/paws/cmd/common/elf"
)

func main() {
	f, err := os.OpenFile("main", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	elf.EncodeElf(f, elf.NewElf())
	if err != nil {
		panic(err)
	}

	fmt.Printf("done :3")
}
