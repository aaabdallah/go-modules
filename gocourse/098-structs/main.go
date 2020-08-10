package main

import ( "fmt"; "time" )

type person struct {
	firstName string
	lastName string
}

type family struct {
	members []person
}

func main() {

	type myint2 int

	fmt.Println( "Program began at", time.Now() )

	// simpleType()

	// testFamily()

	// testEmployees()

	anonStructs()
}

func anonStructs() {

	p1 := struct {
		firstName string
		lastName string
	}{ firstName:"Ahmed", lastName: "Abd-Allah" }

	fmt.Println("Anonymous structure:", p1)

}

type employee struct {
	person // an embedded type: an approximation of inheritance
	badge string
}

func testEmployees() {

	emp01 := employee{
		person : person{ firstName : "Ahmed", lastName : "Abd-Allah" },
		badge : "n72678", 
	}

	fmt.Println( emp01.firstName, emp01.lastName, emp01.badge )
}

func testFamily() {
	p1 := person{ "ahmed", "abd-Allah" }
	fmt.Println( p1 )

	p2 := person{ lastName : "Roberts" }
	fmt.Println( p2 )

	fmt.Println( p1.firstName, p2.lastName )

	fam1 := family{
		members : []person{ p1, p2 },
	}

	fmt.Println( "Family:", fam1 )
}

type myint int
func simpleType() {

	var x myint = 3
	fmt.Printf("x has type %T and value %v\n", x, x)

	y := myint(4)
	fmt.Printf("y has type %T and value %v\n", y, y)

}
