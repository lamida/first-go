package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		c <- 5
		c <- 6
		time.Sleep(2 * time.Second)
	}()

	for xx := range c {
		fmt.Println(xx)
	}
}
