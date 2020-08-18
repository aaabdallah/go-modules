package main

import (
	"fmt"
)

func main() {

	c := make( chan int )

	go gen(c)

	rcv2(c)
}

func gen(c chan int) {

	for i:=0 ; i<100; i++ {
		c <- i
	}
	close( c )
}

func rcv2( c chan int ) {

	fmt.Println()
	forLoop: for {
		select {
		case i, ok := <-c :
			if !ok { break forLoop }
			fmt.Printf("%d ", i)
			if (i+1) % 30 == 0 { fmt.Println() }
		}
	}
	fmt.Println()

}

func rcv1( c chan int ) {

	fmt.Println()
	for i := range c {
		fmt.Printf("%d ", i)
		if (i+1) % 30 == 0 { fmt.Println() }
	}
	fmt.Println()

}
