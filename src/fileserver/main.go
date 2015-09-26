package main

import (
	"fmt"
)

func main() {
	b := 10
	b = increase(b)
	fmt.Println(b)
}
func increase(a int) int {
	a++
	return a
}
