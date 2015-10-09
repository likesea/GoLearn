package main

import (
	"testing"
)

func Benchmark_getWithCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getTreeRootWithCopy()
	}
}
func Benchmark_getWithPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getTreeRootWithPointer()
	}
}
