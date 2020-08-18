package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething( h human ) {
	h.speak()
}

func personSaySomething( p person ) {
	p.speak()
}

func main() {
	p01 := person{ "Ahmed", "Abd-Allah", 52 }

	fmt.Println( p01 )

	saySomething( &p01 )

	personSaySomething( &p01 )
}
