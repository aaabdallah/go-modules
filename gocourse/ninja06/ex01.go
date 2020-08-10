package main

import "fmt"

func main() {

	fmt.Println( foo() )

	fmt.Println( bar() )
}

func foo() int {
	return 100
}

func bar() (int, string) {
	return 200, "two hundred"
}
