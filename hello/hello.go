package main

import "rsc.io/quote"
import "fmt"

func Hello() string {
	//fmt.Printf( quote.Hello() )
	return quote.Hello()
}

func main() {
	fmt.Printf( "Hello, world." )
	fmt.Println( "Hello, world." )

	foo()

	for i := 0; i < 10;  i++ {
		foo()
		if i % 2 == 0 {
			fmt.Println( i )
		}
	}
}

func foo() {
	fmt.Println( "I'm in foo" )
}

