package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	jtree "github.com/pixjay/gochapter/chapter6/tree"
)

func main() {
	t := tree.New(1)
	c := make(chan int)
	go jtree.Walk(t, c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	t1 := tree.New(1)
	// insert(t1, 20)
	t2 := tree.New(1)
	fmt.Println(jtree.Same(t1, t2))
}
