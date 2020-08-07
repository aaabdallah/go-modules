package main

import "fmt"

func main() {

	//printArgTypes( 3, 4, 5.2, 30 )
	//makeSlices()
	//arrays01()
	//arrays02()
	//arrays03()
	playWithMaps()
}

func playWithMaps() {

	map01 := map[string]int{ "one":1, "two":2, "four":4 }

	fmt.Println( map01 )

	_, present := map01["three"]
	fmt.Println( "Is three present: ", present )
	if v, ok := map01["three"]; ok {
		fmt.Println("--three is in the map", v)
	} else {
		fmt.Println("--three is NOT in the map", v)
	}

	map01["three"] = 3

	_, present = map01["three"]
	fmt.Println( "Is three present: ", present )
	if v, ok := map01["three"]; ok {
		fmt.Println("--three is in the map", v)
	} else {
		fmt.Println("--three is NOT in the map", v)
	}

	delete( map01, "three" )

	_, present = map01["three"]
	fmt.Println( "Is three present: ", present )
	if v, ok := map01["three"]; ok {
		fmt.Println("--three is in the map", v)
	} else {
		fmt.Println("--three is NOT in the map", v)
	}
	// ===================================

	map02 := make( map[string]float64 )
	fmt.Println( map02 )

	map02["Gold"] = 32.15
	fmt.Println( map02 )

}

func makeSlices() {
	arr := []byte{0, 1, 2}
	fmt.Println( arr )
	arr2 := make([]byte, 6, 12)
	fmt.Println( arr2 )

	matrix01 := [][]int{ {0, 1, 3, 5}, {20, 21} }
	fmt.Println( matrix01 )
}

// Build a slice on an array, then append to the slice: will it overflow the underlying array?
// Answer: No: since the underlying array is fixed in size, AND THE SLICE IS ALREADY SET TO ITS
// MAXIMUM CAPACITY (THE LENGTH OF THE ARRAY), the slice is reallocated (without
// touching the original array). Changes to the resulting new slice (append returns that newly
// allocated slice) do NOT change the original array.
func arrays02() {
	fmt.Println( "\nTesting a slice that is the length of the array" )
	array01 := [3]int{ 10, 20, 30 }
	slice01 := array01[:]

	array01[1] = 25 // set the second element of the array
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 )
	fmt.Println("S:", slice01 )

	slice01 = append( slice01, 40 )
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 )
	fmt.Println("S:", slice01 ) // The 2nd element of the slice should be 25

	slice01[1] = 27
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 ) // The 2nd element of the array should NOT be 27 and should still be 25: it was never changed
	fmt.Println("S:", slice01 )

	// The nail in the coffin: change the first element of the slice. Check both the slice and the array
	fmt.Println("Changing the first element of the slice to -1")
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	slice01[0] = -1
	fmt.Println("A:", array01 )
	fmt.Println("S:", slice01 )
}

// What happens when the slice is smaller than the original array? If the appends to it keep it less than the
// size of the array, then the appends DO AFFECT THE UNDERLYING ARRAY (the appends and any changes to existing
// elements in the slice).
// YOU HAVE BEEN WARNED.
// YOU HAVE BEEN WARNED.
// YOU HAVE BEEN WARNED.
// YOU HAVE BEEN WARNED.
func arrays03() {
	fmt.Println( "\nTesting a slice that is the LESS than the length of the array" )
	array01 := [3]int{ 10, 20, 30 }
	slice01 := array01[:2]

	array01[1] = 25 // set the second element of the array
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 )
	fmt.Println("S:", slice01 )

	slice01 = append( slice01, 40 )
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 ) // The 3rd element of the array should be 40
	fmt.Println("S:", slice01 ) 

	slice01[1] = 27
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 ) // The 2nd element of the array should be 27: the slice still fits in it
	fmt.Println("S:", slice01 )

	// The nail in the coffin: change the first element of the slice. Check both the slice and the array
	fmt.Println("Changing the first element of the slice to -1")
	slice01[0] = -1
	fmt.Println("Length of array:", len(array01), "Length of slice:", len(slice01), ", Capacity of slice:", cap(slice01))
	fmt.Println("A:", array01 )
	fmt.Println("S:", slice01 )
}

func arrays01() {
	var arr0 [3]int
	var arr1 [5]int
	var arr2 [10]int
	var arr3 [3]int

	fmt.Println( arr0 )
	fmt.Println( arr1 )
	fmt.Println( arr2 )

	arr0[1] = 111
	fmt.Println( arr0 )
	fmt.Println( len(arr0) )

	arr3 = arr0
	fmt.Println( arr3 )
	fmt.Println( len(arr3) )

	fmt.Println()
	printSlice("arr0[0:]", arr0[0:] )
	setValueInSlice( arr0[0:], 1, 555 )
	printSlice("arr0[0:]", arr0[0:] )
	fmt.Println( arr0 )

	fmt.Println()
	printSlice("arr1[0:]", arr1[0:] )
	setValueInSlice( arr1[0:], 3, 777 )
	printSlice("arr1[0:]", arr1[0:] )
	fmt.Println( arr1 )

	fmt.Println("Starting wonky slice assignments")
	slc01 := arr0[0:]
	slc02 := arr1[0:]
	printSlice("slc01", slc01 )
	printSlice("slc02", slc02 )

	slc02 = slc01
	printSlice("slc01", slc01 )
	printSlice("slc02", slc02 )
}

func printSlice( name string, slice []int ) {

	fmt.Printf( "Slice %s [L=%d]: %v\n", name, len(slice), slice )

}

func setValueInSlice(input []int, index int, value int) (err error) {
	input[ index ] = value

	return nil
}

// If the args share the same type, they can be expressed as a list
// in a var declaration so the type comes at the end
func printArgTypes(a, b int, x, y float64) {
	fmt.Printf("%T %T %T %T\n", a, b, x, y)
}
