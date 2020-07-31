package main

import "fmt"

type wholenumber int

var y int = 43

func main() {

	var x wholenumber

	fmt.Printf("%v %#v %+v %T\n", x, x, x, x)

	x = 42

	fmt.Printf("%v %#v %+v %T\n", x, x, x, x)

	y = int(x)

	fmt.Printf("%v %#v %+v %T\n", y, y, y, y)
}