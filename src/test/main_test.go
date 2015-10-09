package main

import (
	"fmt"
	"sync"
	"testing"
)

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}
func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func add(x, y int) (z int) {
	z = x + y
	return
}
func main() {
	fmt.Println(add(1, 2))
}
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}
func BenchmarkTestDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testdefer()
	}
}
