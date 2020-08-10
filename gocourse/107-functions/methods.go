package main

import ( "fmt"; "time" )

type person struct{
	first string
	last string
}

// An employee is a person
type employee struct{
	person
	salaryPayments map[time.Month]float32
}

// A contractor is a person
type contractor struct{
	person
	invoicePayments map[time.Time]float32
}

type staffMember interface{
	pay( payment float32 )
}

func (e *employee) pay( payment float32 ) {
//	fmt.Printf("Inside employee.pay with type %T\n", e)
//	fmt.Println("Before payment of", payment, ":", e.salaryPayments)

	if e.salaryPayments == nil {
		e.salaryPayments = make( map[time.Month]float32 )
	}

	e.salaryPayments[ time.Now().Month() ] = payment

//	fmt.Println("After payment of", payment, ":", e.salaryPayments)
}

func (c *contractor) pay( payment float32 ) {
//	fmt.Printf("Inside contractor.pay with type %T\n", c)
//	fmt.Println("Before payment of", payment, ":", c.invoicePayments)

	if c.invoicePayments == nil {
		c.invoicePayments = make( map[time.Time]float32 )
	}

	c.invoicePayments[ time.Now() ] = payment

//	fmt.Println("After payment of", payment, ":", c.invoicePayments)
}

func deservesBonus( sm staffMember ) bool {

	switch sm.(type) {
		case *employee:
			return true
		case *contractor:
			return false
		default:
			return false
	}
}

func main() {

	p1 := person{ first : "Ahmed", last : "Abd-Allah" }

	e1 := employee{ person : p1 }
	c1 := contractor{
		person : person{ first : "Abdulrahman", last : "Mohammed" },
	}

	fmt.Println( e1 )
	fmt.Println( c1 )

	e1.pay( 100 )
	e1.pay( 200 )
	e1.pay( 300 )

	c1.pay( 100 )
	c1.pay( 200 )
	c1.pay( 300 )

	fmt.Println( e1 )
	fmt.Println( c1 )

	staff := []staffMember{ &e1, &c1 }

	// Pause to make sure the time changes enough to avoid overwriting the contractor's
	// payments (yes, Go is fast!)
	time.Sleep( 100 * time.Millisecond )

	fmt.Println("\nProceeding to pay everyone 999")
	for _, v := range staff {
		v.pay( 999 )

		fmt.Println( v )
		if deservesBonus( v ) {
			fmt.Println("\t ... and s/he deserves a bonus!")
		}

	}
}
