package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	x := 25
	y := x / 2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(testFunc(x, y))
}

func testFunc(x int, y int) int {
	z := x + y
	return z
}
