package main

import "fmt"

type person struct {
	first string
	last string
}

func (p *person) setFirst( f string ) {

	p.first = f

}

func main() {

	p1 := person { "Ahmed", "Abd-Allah" }

	p1.setFirst( "Muhammad" )

	fmt.Println( p1 )

}
