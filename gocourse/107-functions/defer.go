package main

import ( "fmt"; "time" )

func main() {

	fmt.Println( "Starting at", time.Now() )

	fmt.Println( "Integer Division: 5 / 3 = ", integerDivide( 5, 3 ) )
	fmt.Println( "Integer Division: 5 / 0 = ", integerDivide( 5, 0 ) )
}

func integerDivide( dividend, divisor int ) int {

	defer func() {
		fmt.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v\n", x)
		}
	}()

	return dividend / divisor

}

func raise( base, exponent int ) int {

	if exponent == 0 { return 1 }

	return base * raise( base, exponent - 1 )

}
