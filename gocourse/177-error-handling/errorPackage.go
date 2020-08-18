package main

import (
	"fmt"
	"errors"
)

func main() {

	fmt.Println( "MAIN: Starting..." )

	errorInfo()

	fmt.Println( "MAIN: Ending..." )

}

func errorInfo() {

	e := errors.New("Error message here")

	fmt.Printf( "%T\n%v\n%+v\n%#v\n", e, e, e, e)

	fmt.Printf( "Error string details: %s\n", e.Error() )

}
