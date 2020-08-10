package main

import "fmt"

func main() {

	fmt.Println("Starting...")

	defer func() {
		fmt.Println("The deferred function")
	}()

	fmt.Println("Ending...")

}