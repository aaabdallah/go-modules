package main

import "fmt"

func main() {

	multiply := func(x, y float32) float32 {
		return x * y
	}

	add := func(x, y float32) float32 {
		return x + y
	}

	fmt.Println( add(100, add(5, multiply( 3, 4 ))) )
}
