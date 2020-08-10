package main

import ( "fmt"; "strconv" )

func main() {

	add := func(x, y int) int {
		return x + y
	}

	fmt.Printf("%T\n", add)

	fmt.Println( add(3, 4) )
	fmt.Println( add( add(3, 4), add(5, 6) ) )

	mult := func(x, y int) int {
		product := 0
		for i:=0 ; i<x ; i++ {
			product = add( product, y )
		}
		return product
	}

	fmt.Println( mult(7,4) )

	callback005 := makeCallback( 5 )
//	callback017 := makeCallback( 17 )

	fmt.Println( makeCallback(5)() )
	fmt.Println( makeCallback(17)() )

	fmt.Println( callback005() )
	fmt.Println( callback005() )
	fmt.Println( callback005() )
	fmt.Println( callback005() )
}

func makeCallback( id int ) func() string {

	return func() string {
		// fmt.Println( "Callback for ID: ", id, string(id) )
		id++
		return "Callback for ID: " + strconv.Itoa( id )
	}

}