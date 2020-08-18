package main

import "fmt"

func main() {

	numberOfGoroutines := 3
	numbersToGenerate := 5

	c := make( chan int )

	for i:=0 ; i<numberOfGoroutines ; i++ {
		go gen(c, numbersToGenerate)
	}

	rcv( c, numberOfGoroutines * numbersToGenerate )

}

func gen(c chan int, count int) {

	for i:=1 ; i<=count ; i++ {
		c <- i
	}

}

func rcv(c chan int, count int) {

	fmt.Println()
	for i:=0 ; i<count ; i++ {
		fmt.Printf("%d ", <-c)
		if (i+1) % 30 == 0 { fmt.Println() }
	}
	fmt.Println()

}
