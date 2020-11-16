package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	threadCount := 10
	ch := make(chan int, 5)
	for threadID := 0; threadID < threadCount; threadID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("thread %v is blocked\n", id)
			ch <- id
			fmt.Printf("thread %v is unblocked\n", id)
			time.Sleep(1000 * time.Millisecond)
		}(threadID)
	}
	time.Sleep(5000 * time.Millisecond)
	for i := 0; i < threadCount; i++ {
		fmt.Printf("releasing thread %v\n", <-ch)
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Wait()
}