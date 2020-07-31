package main

import "fmt"

type wholenumber int

func main() {

	var x wholenumber

	fmt.Printf("%v %#v %+v %T\n", x, x, x, x)

	x = 42

	fmt.Printf("%v %#v %+v %T\n", x, x, x, x)
}