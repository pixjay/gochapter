package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		Walk(t1, c1)
		close(c1)
	}()
	go func() {
		Walk(t2, c2)
		close(c2)
	}()
	for {
		v1, e1 := <-c1
		v2, e2 := <-c2
		if v1 != v2 || e1 != e2 {
			return false
		} else if !e1 || !e2 {
			break
		}
	}
	return true
}

func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func main() {
	t := tree.New(1)
	c := make(chan int)
	go Walk(t, c)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	t1 := tree.New(1)
	// insert(t1, 20)
	t2 := tree.New(1)
	fmt.Println(Same(t1, t2))
}
