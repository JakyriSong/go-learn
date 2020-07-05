package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt 开平方非法错误
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintln("cannot Sqrt negative number:", float64(e))
}

// Sqrt 牛顿开方
func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		z := 1.0
		dt := math.Inf(1)
		for math.Abs(dt) > 0.000000000001 {
			dt = (z*z - x) / (2 * z)
			z -= dt
			fmt.Println(z)
		}
		return z, nil
	}
	return 0, ErrNegativeSqrt(x)

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
