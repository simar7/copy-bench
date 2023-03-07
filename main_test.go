package main

import (
	"testing"
)

func BenchmarkReadAllFunc_SmallFile(b *testing.B) {
	b.ReportAllocs()
	ReadAllFunc("small.txt")
}

func BenchmarkReadAllFunc_BigFile(b *testing.B) {
	b.ReportAllocs()
	ReadAllFunc("big.txt")
}

func BenchmarkCopyFunc_SmallFile(b *testing.B) {
	b.ReportAllocs()
	CopyFunc("small.txt")
}

func BenchmarkCopyFunc_BigFile(b *testing.B) {
	b.ReportAllocs()
	CopyFunc("big.txt")
}

func BenchmarkSmallFile(b *testing.B) {
	b.ReportAllocs()
	//ReadAllFunc("small.txt")
	CopyFunc("small.txt")
}

func BenchmarkBigFile(b *testing.B) {
	b.ReportAllocs()
	//ReadAllFunc("big.txt")
	CopyFunc("big.txt")
}
