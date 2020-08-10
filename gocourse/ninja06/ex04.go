package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {

	fmt.Println("My name is", p.first, "and I am", p.age, "years old.")

}

func main() {

	p1 := person { first:"Ahmed", last:"Abd-Allah", age:52 }

	p1.speak()

}
