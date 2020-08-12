package main

import "fmt"

type circle struct {
	radius float64
}

func (c *circle) setRadius( r float64 ) {
	c.radius = r
}

func main() {

	x := 25
	y := 50

	fmt.Printf("x %d %d\n", x, &x)
	fmt.Printf("y %d %d\n", y, &y)

	//fmt.Println( *( &x + *int(2) ) )


	c01 := circle{ 10 }
	fmt.Println( c01.radius )
	c01.setRadius( 20 )
	fmt.Println( c01.radius )
}
