package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	myFuncWithVariadicParameter("a")

	printAllIntegers(1, 3, 5, 7, 9, 11)

	printArrayOfAnything("abc", 123, true)
}

func myFuncWithVariadicParameter(a ...interface{}) {
}

func printArrayOfAnything(arrayOfAnything ...interface{}) {

	for i := 0 ; i < len(arrayOfAnything) ; i++ {
		fmt.Println("Array element", i, "is", arrayOfAnything[i])
	}
}

func printAllIntegers(intArray ...int) {
	fmt.Println( "Length of passed in array: ", len(intArray) )

	for i := 0; i < len(intArray) ; i++ {
		fmt.Printf("At position %d, value = %d\n", i, intArray[i])
	}

	n, err := fmt.Println( intArray )

	_, err2 := fmt.Println("n:", n)
	_ = err2
	fmt.Println("err:", err)
}
