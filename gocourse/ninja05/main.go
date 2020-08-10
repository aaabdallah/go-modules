package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println( "Starting at:", time.Now() )

	exercise()

}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
	
type sedan struct {
	vehicle
	luxury bool
}

func exercise() {
	t1 := truck{
		vehicle : vehicle{
			doors : 2,
			color : "blue",
		},
		fourWheel : true,
	}

	s1 := sedan {
		vehicle : vehicle{
			doors : 4,
			color : "white",
		},
		luxury : true,
	}

	fmt.Printf( "%#v\n", t1 )
	fmt.Printf( "%#v\n", s1 )

	fmt.Println( t1.doors )
	fmt.Println( s1.color )

	s2 := struct {
		vehicle
		luxury bool 
	}{
		vehicle : vehicle{
			doors : 2,
			color : "gray",
		},
		luxury : false,
	}

	fmt.Printf( "%#v\n", s2 )
}

type person struct {
	firstName string
	lastName string
	iceCreamFavorites []string
}

func ex01and2() {

	p1 := person{
		firstName: "Ahmed",
		lastName: "Abd-Allah",
		iceCreamFavorites: []string{ "Peanut Butter Chocolate", "Love Potion"},
	}

	p2 := p1

	p2.firstName = "Amatullah"
	p2.iceCreamFavorites = append(p2.iceCreamFavorites, "Salted Caramel") // *MAY* reallocate a new slice (see previous assignments)

	p3 := p1

	p3.firstName = "Abd-Allah"
	p3.iceCreamFavorites = []string{}

	fmt.Printf( "%#v\n", p1 )
	fmt.Printf( "%#v\n", p2 )

//	printPerson( p1 )
//	printPerson( p2 )
//	printPerson( p3 )

	map01 := make( map[string]person )
	fmt.Println( map01 )

	map01[ p1.firstName+p1.lastName ] = p1
	map01[ p2.firstName+p2.lastName ] = p2
	map01[ p3.firstName+p3.lastName ] = p3

	for _, v := range map01 {
		printPerson( v )
	}	
}

func printPerson( p person ) {
	fmt.Println("Person")
	fmt.Printf("\tFirst Name: %s\n", p.firstName)
	fmt.Printf("\tLast Name: %s\n", p.lastName)
	fmt.Printf("\tIce Cream Favorites:")
	if len(p.iceCreamFavorites) > 0 {
		fmt.Printf("\n\t\t")
		for i, v := range p.iceCreamFavorites {
			fmt.Printf( "%v", v )
			if i < len(p.iceCreamFavorites)-1 {
				fmt.Printf( "," )
			}
		}
		fmt.Println()
	} else {
		fmt.Printf(" None Specified\n")
	}
}

	