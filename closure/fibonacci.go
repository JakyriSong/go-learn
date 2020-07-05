package main

import (
	"fmt"
)

func fabonacci() func() int {
	x, y, z := 0, 0, 1
	return func() int {
		x, y, z = y, z, y+z
		return x
	}
}

func main() {
	f := fabonacci()
	for i := 1; i < 10; i++ {
		fmt.Println(f())
	}
}
