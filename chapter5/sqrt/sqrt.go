package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {	
	if x < 0.0 {
		e := errNegativeSqrt(x)
		return x, e
	}
	z := 1.0
	for i := 1; i < 10; i++ {
		y := z - (z*z - x) / (2*z)
		z = math.Abs(y)
	}
	return z, nil
}

func main() {
	if x, e := sqrt(2); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(x)
	}
	if x, e := sqrt(-2); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(x)
	}
}
