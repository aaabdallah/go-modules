package main

import "fmt"

const (
	a, b = 1, iota
	c, d
)

const (
	d1 = iota
	d2
	d3
)

func main() {

	fmt.Println("Beginning...")

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)

}
