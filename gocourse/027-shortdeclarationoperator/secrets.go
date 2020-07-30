package main

import "fmt"

// accessible across entire package
var secret = "Secret inside secrets.go"

func printMainSecret() {
	fmt.Println(mainSecret)
}

func printLocalVariables() {

	var (
		weight int = 65
		height int = 165
		age int = 20
	)
	fmt.Println( age, weight, height )
}