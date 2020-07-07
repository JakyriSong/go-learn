package main

import (
	"fmt"
	"io"
	"strings"
)

func addElem(b *[]byte) {
	*b = append(*b, 'a')
	return
}

func main() {
	c := make([]byte, 0, 5)
	addElem(&c)
	fmt.Printf("%c\n", c[0])

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
