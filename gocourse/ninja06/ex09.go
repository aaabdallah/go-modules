package main

import "fmt"

func main() {

	fmt.Println( "Starting..." )

	success := func() string { return "success" }
	failure := func() string { return "failure" }

	foo( success )

	foo( failure )
}

func foo( f func() string ) {

	fmt.Println( "Inside foo with ", f() )
}

