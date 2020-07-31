package main

import (
	"fmt"
	"time"
)

type temperature int

var (
	pkgx = 23
	pkgy = "def"
)


func experimentWithTypes() {
	var x int = 23
	y := "abc"

	fmt.Println("Starting experimentWithTypes()", time.Now())
	fmt.Printf("%T %T %T %T\n", x, y, pkgx, pkgy)
	fmt.Println(pkgy)

	var xptr *int
	fmt.Println(xptr)

	fmt.Printf("%v | %+v | %#v \n", x, x, x)
	fmt.Printf("%v | %+v | %#v \n", y, y, y)

	var daytimeTemperature temperature = 32
	daytimeTemperature = 15
	x = 15
	x = int(daytimeTemperature)
	daytimeTemperature = temperature(x)

	fmt.Println( daytimeTemperature )
}
