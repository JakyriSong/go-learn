package main

import (
	"fmt"
)

func main() {
	var i interface{}
	describe(i)
	s, ok := i.(int)
	fmt.Println(s, ok)

	i = 42
	describe(i)
	s, ok = i.(int)
	fmt.Println(s, ok)

	i = "hello"
	describe(i)
	f, ok := i.(string)
	fmt.Println(f, ok)
}

// describe 接口描述
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
