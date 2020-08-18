package main

import (
	"fmt"
	"time"
)

type customErr struct {
	timestamp time.Time
	details string
}

func (err customErr) Error() string {
	return err.timestamp.String() + ": " + err.details
}

func main() {
	fmt.Println("MAIN: 000")

	foo( customErr{} )

	fmt.Println("MAIN: 999")
}

func foo( errIn error ) { // don't use customErr as the type; use the general interface error

	fmt.Println( "foo:", errIn.Error() )

}
