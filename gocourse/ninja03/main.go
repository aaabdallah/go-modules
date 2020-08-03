package main

import "fmt"

func main() {

	exercise()

}

func exercise() {
	fmt.Println("###################################")
	fmt.Println("### Ex 10 #########################")

	fmt.Println(true && true) 
	fmt.Println(true && false) 
	fmt.Println(true || true) 
	fmt.Println(true || false) 
	fmt.Println(!true)

	fmt.Println("\n")
}

func ex09() {
	fmt.Println("###################################")
	fmt.Println("### Ex 09 #########################")

	reactToFavoriteSport( "basketball" )
	reactToFavoriteSport( "football" )
	reactToFavoriteSport( "soccer" )
	reactToFavoriteSport( "surfing" )

	fmt.Println("\n")
}

func reactToFavoriteSport(favSport string) {

	fmt.Print( favSport, ": " )
	switch (favSport) {
		case "basketball":
			fmt.Println("A sport for athletic graceful giants")
		case "football":
			fmt.Println("A sport for brute force but using great strategy")
		case "soccer":
			fmt.Println("A sport for the entire world")
		default:
			fmt.Println("Sorry, I don't recognize that sport!")
	}
}

func ex08() {
	fmt.Println("###################################")
	fmt.Println("### Ex 08 #########################")

	for r := '\u0000' ; r <= '\u0100' ; r++ {
		classifyASCII( r )
	}

	fmt.Println("\n")
}

func classifyASCII(c rune) {

	switch {
		case c < 0:
			fmt.Printf("%d is less than zero, and not ASCII\n", c)
		case (c >=0 && c <= 31) || c == 127:
			fmt.Printf("%d is an unprintable ASCII character\n", c)
		case c > 127:
			fmt.Printf("%d is a non-ASCII character in the extended ASCII or Unicode encoding\n", c)
		default:
			fmt.Printf("%d is an ASCII character: %#U\n", c, c)
	}

}

func ex07() {

	ex06()

}

func ex06() {
	fmt.Println("###################################")
	fmt.Println("### Ex 06 #########################")

	useIf( 29 )
	useIf( 30 )
	useIf( 31 )

	fmt.Println("\n")
}

func useIf( num int ) {

	if num < 30 {
		fmt.Println(num, "is less than 30")
	} else if num == 30 {
		fmt.Println(num, "is equal to 30")
	} else {
		fmt.Println(num, "is greater than 30")
	}

}

func ex05() {
	fmt.Println("###################################")
	fmt.Println("### Ex 05 #########################")

	for i:=10 ; i<=100 ; i++ {
		fmt.Println(i, i%4)
	}

	fmt.Println("\n")
}

func ex04() {
	fmt.Println("###################################")
	fmt.Println("### Ex 04 #########################")

	i := 52
	for {
		fmt.Println( 2020 - i )
		i--
		if i < 0 { break }
	}

	fmt.Println("\n")
}

func ex03() {
	fmt.Println("###################################")
	fmt.Println("### Ex 03 #########################")

	i := 52
	for i >= 0 {
		fmt.Println( 2020 - i )
		i--
	}

	fmt.Println("\n")
}

func ex02() {
	fmt.Println("###################################")
	fmt.Println("### Ex 02 #########################")

	for r := 'A' ; r <= 'Z' ; r++ {
		fmt.Printf("%d\n", r)
		for i := 0 ; i < 3 ; i++ {
			fmt.Printf("\t%#U\n", r)
		}
	}

	fmt.Println("\n")
}

func ex01() {
	fmt.Println("###################################")
	fmt.Println("### Ex 01 #########################")

	for i:=1; i<=10000; i++ {
		// fmt.Print(i, " ")
		fmt.Println( i )
	}

	fmt.Println("\n")
}
