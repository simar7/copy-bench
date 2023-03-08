package main

import (
	"bytes"
	"encoding/json"
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

func CopyAndUnmarshal(filename string) interface{} {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		panic(err)
	}
	data := buf.Bytes()

	var contents interface{}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		panic(err)
	}

	return contents
}

func JSONDecode(filename string) interface{} {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var contents interface{}
	err = json.NewDecoder(f).Decode(&contents)
	if err != nil {
		panic(err)
	}

	return contents
}
