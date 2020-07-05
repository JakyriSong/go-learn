package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 9; i++ {
		fmt.Println("random num(0~100): ", rand.Intn(10))
	}
}
