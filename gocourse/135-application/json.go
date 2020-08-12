package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
}

type employee struct {
	person
	Salary float32
}

func main() {

	fmt.Println( "Begin JSON experimentation..." )

	testJSONStructs()
}

func testJSONStructs() {

	p01 := person{"Ahmed", "Abd-Allah", 52 }
	p01Bytes, err := json.Marshal( p01 )
	if err != nil {
		fmt.Println( "Unable to marshal data" )
		return
	}

	fmt.Printf("p01Bytes of type %T\n\tbytes: %v\n\tstring: %v\n", p01Bytes, p01Bytes, string(p01Bytes))

	fmt.Println()

	var p02 person
	err = json.Unmarshal( p01Bytes, &p02 )
	if err != nil {
		fmt.Println( "Unable to unmarshal data:\n", err )
		return
	}
	fmt.Printf("p02: %v\n", p02)

	emp01 := employee{ p02, 11000 }
	emp01Bytes, err := json.Marshal( emp01 )

	if err != nil {
		fmt.Println( "Unable to marshal data" )
		return
	}

	fmt.Printf("emp01Bytes of type %T\n\tbytes: %v\n\tstring: %v\n", emp01Bytes, emp01Bytes, string(emp01Bytes))
}

func testMarshalPrimitives() {
	x := 100
	
	fmt.Println( json.Marshal( x ) )
}

func testByteSliceStringConversions() {

	byteSlice := []byte{ 49, 49, 49 }

	fmt.Println( byteSlice )

	str01 := string( byteSlice )

	fmt.Println( str01 )

	fmt.Println( []byte( str01 ) )

}
