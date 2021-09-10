package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fibo(n)
	fmt.Printf("\rFibbonaci(%d) = %d\n", n, fibN)
}

func fibo(x int) int {
	if x < 2 {
		return x
	} else {
		return fibo(x-1) + fibo(x-2)
	}
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
