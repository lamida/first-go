package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Walk2(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t != nil {
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
		}
	}
	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk2(t1, ch1)

	ch2 := make(chan int)
	go Walk2(t2, ch2)

	for {
		a, ok := <-ch1
		if !ok {
			break
		}
		b, ok := <-ch2
		if !ok {
			return false
		}
		if a != b {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	go Walk2(t, ch)
	for c := range ch {
		fmt.Println(c)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
}
