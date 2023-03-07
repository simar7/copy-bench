package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	//ReadAllFunc()
	//CopyFunc()
}

func ReadAllFunc(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var data []byte
	if data, err = io.ReadAll(f); err != nil {
		panic(err)
	}

	return data
}

func CopyFunc(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
