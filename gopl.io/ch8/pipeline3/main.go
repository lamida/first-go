package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	go func() {
		for x := 0; ; x++ {
			out <- x
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

func squarer(out chan<- int, in <-chan int) {
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	counter(naturals)
	squarer(squares, naturals)
	printer(squares)
}
