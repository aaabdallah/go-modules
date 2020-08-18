package main

import (
	"fmt"
)

func main() {
	q := make(chan int) // used to signal all done loading
	c := gen(q) // used to place work into

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}

