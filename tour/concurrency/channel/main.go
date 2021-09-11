package main

import "fmt"

func sum(inn []int, c chan<- int) {
	sum := 0
	for _, v := range inn {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	inn := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(inn[:len(inn)/2], c)
	go sum(inn[len(inn)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
