package main

import "fmt"

func main() {

	helloDr := makeGreeter("Hello", "Dr.")
	welcomeMr := makeGreeter("Welcome", "Mr.")

	fmt.Println( helloDr( "Ahmed" ) )

	fmt.Println( welcomeMr( "Khaled" ) )

}

// To my understanding, the real value of building and returning functions
// is with closure. Otherwise, the value is somewhat lost on me...
// In this case, greeting and prefix are the enclosed variables
func makeGreeter(greeting, prefix string) (func(string) string) {

	return func(name string) string {
		return greeting + " " + prefix + " " + name;
	}

}