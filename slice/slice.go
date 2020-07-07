package main

import (
	"fmt"
)

func addElem(b []byte) {
	//s := []byte{'a', 's', 'd'}
	fmt.Printf("b: %p\n", &b)
	//b[0] = 'A'
	b = append(b, 'c')
	fmt.Printf("b: %p ", &b)
	fmt.Println(b)
	return
}

func main() {
	c := make([]byte, 1)
	addElem(c)
	fmt.Printf("c: %p ", &c)
	fmt.Println(c)

	// r := strings.NewReader("Hello, Reader!")

	// b := make([]byte, 8)
	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
}
