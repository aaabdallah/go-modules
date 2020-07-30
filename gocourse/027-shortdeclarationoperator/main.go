// must have package main to 'go run'
package main

import "fmt"
import "math/rand" // part of Golang's standard packages; not reflected in go.mod
import "rsc.io/quote"

// package variables
var mainSecret = "Secret inside main.go"

func main() {
	x := 20
	_ = x // illustrate how to get away with 'unused' variable. very annoying.

	fmt.Println(rand.Int())

	fmt.Println(quote.Hello())

	fmt.Println(secret)

	printMainSecret()

	fmt.Println( "أحمد" )

	printLocalVariables()
}
