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

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		a, b := <-ch1, <-ch2
		fmt.Println(a, b)
		if a != b {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	fmt.Println(t.String())
	go Walk(t, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
}
