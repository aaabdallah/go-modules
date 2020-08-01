package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
}

func ex0() {
	fmt.Println("############################################")
	fmt.Println("Exercise #0\n")


	
	fmt.Println("############################################")
	fmt.Println()
}

func ex06() {
	fmt.Println("############################################")
	fmt.Println("Exercise #06\n")

	const(
		year1 = iota + 2021
		year2
		year3
		year4
	)

	fmt.Println( year1, year2, year3, year4 )
	
	fmt.Println("############################################")
	fmt.Println()
}

func ex05() {
	fmt.Println("############################################")
	fmt.Println("Exercise #05\n")

	s := 
`This is the paragraph.
Here is another sentence.
This is hard - "That's what she said!"`

	fmt.Println(s)
	
	fmt.Println("############################################")
	fmt.Println()
}

func ex04() {
	fmt.Println("############################################")
	fmt.Println("Exercise #04\n")

	var x = 24
	fmt.Printf("%d %b %x\n", x, x, x)

	y := x << 1
	fmt.Printf("%d %b %x\n", y, y, y)
	
	fmt.Println("############################################")
	fmt.Println()
}

func ex03() {
	fmt.Println("############################################")
	fmt.Println("Exercise #03\n")
	
	const x = 24
	const xi int = 24

	fmt.Printf("%f %T %d\n", x, x, xi)

	fmt.Println("############################################")
	fmt.Println()
}

func ex02() {
	fmt.Println("############################################")
	fmt.Println("Exercise #02\n")
	
	g, h, i, j, k, l := "a" == "a", 3 <= 2, 4 >= 3, 1 != 2, 2 < 3, 4 > 10

	fmt.Printf("%v %v %v %v %v %v\n", g, h, i, j, k, l)

	fmt.Println("############################################")
	fmt.Println()
}

func ex01() {
	fmt.Println("############################################")
	fmt.Println("Exercise #01\n")
	num := 24

	fmt.Printf("%d %b %x\n", num, num, num)
	fmt.Printf("%#d %#b %#x\n", num, num, num)

	fmt.Println("############################################")
	fmt.Println()
}
