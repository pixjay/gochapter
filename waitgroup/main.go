package main

import (
	"fmt"
	"sync"
)

func startAndWait() {
	var wg sync.WaitGroup
	count := 10
	for i := 0; i < count; i++ {
		wg.Add(1)
		fmt.Printf("starting %v\n", i)
		go do(i, &wg)
	}
	wg.Wait()
	fmt.Printf("finished\n")
}

func dontWait() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("starting %v\n", i)
		go do(i, nil)
	}
	fmt.Printf("finished\n")
}

func do(i int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Printf("doing something %v\n", i)
}

func main() {
	// dontWait()
	startAndWait()
}
