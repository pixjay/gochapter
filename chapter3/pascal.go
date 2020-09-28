package main

import (
	"fmt"
)

func pascalsTriangleAt(n int) []int {
	if n < 0 {
		return nil
	}
	x := make([]int, n+1)
	x[0] = 1
	if n > 0 {
		x[n] = 1
	}
	if n > 1 {
		y := pascalsTriangleAt(n-1)
		for i := 1; i < n; i++ {
			x[i] = y[i-1] + y[i]
		}
	}
	return x
}

func main() {
	for i := 0; i < 25; i++ {
		fmt.Printf("%v\n", pascalsTriangleAt(i))
	}
}
