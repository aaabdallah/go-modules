package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortableByLast []user
func (s sortableByLast) Len() int { return len(s) }
func (s sortableByLast) Swap( i,j int ) { s[i], s[j] = s[j], s[i] }
func (s sortableByLast) Less( i,j int) bool { return s[i].Last < s[j].Last }

type sortableByAge []user
func (s sortableByAge) Len() int { return len(s) }
func (s sortableByAge) Swap( i,j int ) { s[i], s[j] = s[j], s[i] }
func (s sortableByAge) Less( i,j int) bool { return s[i].Age < s[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	fmt.Println("---")
	for _, v := range users {
		sort.Strings( v.Sayings )
	}
	fmt.Println(users)
	fmt.Println("---")
	sort.Sort( sortableByLast(users) )
	fmt.Println( users )
	fmt.Println("---")
	sort.Sort( sortableByAge(users) )
	fmt.Println( users )
}
