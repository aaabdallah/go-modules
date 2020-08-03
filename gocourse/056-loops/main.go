package main

import "fmt"

func main() {

	loop06()

}

func loop06() {

	printable := false

	for i, c := range "\u0007\u001Fabc3def1AB4CDEF\u0032\u007C" {

		fmt.Printf("[%2d]: ", i)
		switch {
			case c >= 32 && c < 127:
				fmt.Print("Printable ")
				printable = true
			default:
				fmt.Print("Unprintable ")
				printable = false
		}

		switch c {
			case '1','2','3','4','5','6','7','8','9','0':
				fmt.Print("digit ")
			default:
				fmt.Print("non-digit ")
		}

		if printable { 
			fmt.Printf("-- %c\n", c) 
		} else {
			fmt.Printf("-- %U\n", c)
		}

	}

}

func loop05(){

	for x := 0 ; x < 10 ; x++ {
		
		if x++ ; x >= 3 {
			fmt.Println( "Condition met" )
		} else {
			fmt.Println("Condition NOT met")
		}
		fmt.Println("x:", x)
	}

}

func loop04() {
	for ;; { // in this case the semicolons are not needed
		fmt.Println("Infinite Loop")
	}
}

func loop03() {

	a := 0
	b := 5
	for a < b {
		fmt.Println("a is less than b:", a, b)
		a++
	}

	a = 0
	for ; a < b ; {
		fmt.Println("a is less than b:", a, b)
		a++
	}
}

func loop02() {

	for i:=0; i<5; i++ {
		fmt.Printf("%d\n", i)

		for j:=0; j<3; j++ {
			fmt.Printf("\t%d --> %d\n", i, j)
		}
	}

}

func loop01() {

	for i:=0 ; i<10 ; i++ {
		fmt.Printf("%d\n", i)
	}

}

func loop00() {

	i := 0
	for {
		// Doesn't compile in Go. ++ is "punctuation" that forms a statement
		// and therefore can't be used in that location since a variable is
		// expected. Anyway, remember that Go removes expressivity on purpose.
		//fmt.Printf("%d - Infinite Loop\n", i++)

		fmt.Printf("%d - Infinite Loop\n", i)
		i++
	}

}