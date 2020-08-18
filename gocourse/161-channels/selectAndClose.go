package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v, ok := <-e:
			fmt.Println("from the eve channel:", v, ok)
		case v, ok := <-o:
			fmt.Println("from the odd channel:", v, ok)
		case v, ok := <-q:
			fmt.Println("from the quit channel:", v, ok)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
