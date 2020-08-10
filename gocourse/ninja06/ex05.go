package main

import "fmt"

func main() {

	sq01 := square { side:25 }
	cc01 := circle { radius: 2 }

	info( sq01 )

	info( cc01 )

}

func info( sh shape ) {

	var shapeName string
	switch sh.(type) {
		case circle:
			shapeName = "circle"
		case square:
			shapeName = "square"
		default:
			shapeName = "unknown"
	}
	fmt.Println("Shape - ", shapeName, " - area:", sh.area() )

}
