package main

import (
	"io"
	"os"
)

func main() {
	fs, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, fs)
}
