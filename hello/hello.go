package main

import (
	"fmt"

	"github.com/JakyriSong/go-learn/stringutil"
)

func defer_test() int {
	a := 10
	defer func() {
		a++
		fmt.Println(a)
	}()
	fmt.Println(a)
	return a
}

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println(stringutil.Reverse("Hello, Go!"))
	fmt.Println(defer_test())
}
