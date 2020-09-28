package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)

func pascalsTriangleAt(n int) []int {
	if n < 0 {
		return nil
	}
	x := make([]int, n+1)
	x[0] = 1
	x[n] = 1
	if n > 1 {
		y := pascalsTriangleAt(n-1)
		for i := 1; i < n; i++ {
			x[i] = y[i-1] + y[i]
		}
	}
	return x
}

func pascalsTriangleAtNonRecursive(n int) []int {
	x := make([]int, n+1)
	for i := 0; i <= n; i++ {
		x[i] = 1
		for j := i - 1; j > 0; j-- {
			x[j] = x[j] + x[j-1]
		}
	}
	return x
}

func main() {
	count := 10
	if len(os.Args) > 1 {
		countArg, _ := strconv.ParseInt(os.Args[1], 10, 64)
		count = int(countArg)
	}

	fmt.Println("Recursive: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", pascalsTriangleAt(i))
	}

	fmt.Println("Non-Recursive: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", pascalsTriangleAtNonRecursive(i))
	}

	start := time.Now()
	for i := 0; i < count; i++ {
		pascalsTriangleAt(i)
	}
	t := time.Now()
	fmt.Printf("    Recursive time elapsed running %v executions: %v\n", count, t.Sub(start))

	start = time.Now()
	for i := 0; i < count; i++ {
		pascalsTriangleAtNonRecursive(i)
	}
	t = time.Now()
	fmt.Printf("Non-Recursive time elapsed running %v executions: %v\n", count, t.Sub(start))
}
