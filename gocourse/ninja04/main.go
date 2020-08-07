package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Starting at", now)

	// exploreTime( now )

	exercise()
}

func exercise() {

	//var slice01 []string
	//slice01[0] = "Test" // No: no space allocated yet

	// map of People to Favorite Things
	// var mapPFT map[string][]string // doesn't allocate space
	mapPFT := make( map[string][]string )
	// mapPFT["Ahmed"] = []string{"computers", "books"}
	mapPFT[`bond_james`] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mapPFT[`moneypenny_miss`] = []string{`James Bond`, `Literature`, `Computer Science`}
	mapPFT[`no_dr`] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	fmt.Printf("Type: %T, Length: %d\n", mapPFT, len(mapPFT))
	fmt.Println(mapPFT)

	for rc, rv := range mapPFT {
		for ec, ev := range rv {
			fmt.Printf("Row %s, Elem %d = %s\n", rc, ec, ev)
		}
	}

	mapPFT["abdallah_ahmed"] = []string{"computers", "books", "basketball"}

	for rc, rv := range mapPFT {
		for ec, ev := range rv {
			fmt.Printf("Row %s, Elem %d = %s\n", rc, ec, ev)
		}
	}

	delete(mapPFT, "abdallah_ahmed")

	for rc, rv := range mapPFT {
		for ec, ev := range rv {
			fmt.Printf("Row %s, Elem %d = %s\n", rc, ec, ev)
		}
	}

	// strings to arrays of integers to practice composite literals across 2 dimensions
	mapSTA := map[string][]int{
		"a" : []int{1, 2, 3},
		"b" : []int{10, 20, 30},
		"c" : []int{100, 200, 300} }
	fmt.Println( mapSTA )
}

func exercise07() {

	rows := make( [][]string, 2 )

	rows[0] = []string{ "James", "Bond", "Shaken, not stirred" }
	rows[1] = []string{ "Miss", "Moneypenny", "Helloooooo, James." }

	for rcntr, rowvalue := range rows {
		fmt.Println("Row:", rcntr)
		for ecntr, elemvalue := range rowvalue {
			fmt.Println("\t", "Element:", ecntr, elemvalue)
		}
	}

}

func exercise06b() {

	states := make( []string, 50, 50 )
	printStrSlice( "states", states )
	fmt.Printf("Len: %d -- Cap: %d\n", len(states), cap(states))

	states = append( states, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)

	printStrSlice( "states", states )
	fmt.Printf("Len: %d -- Cap: %d\n", len(states), cap(states))

	for i:=0; i<len(states); i++ {
		fmt.Println(i, states[i])
	}
}

func exercise06a() {

	states := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`,}

	printStrSlice( "states", states )

	fmt.Printf("Len: %d -- Cap: %d\n", len(states), cap(states))

	for i:=0; i<len(states); i++ {
		fmt.Println(i, states[i])
	}
}

func printStrSlice( name string, slice []string ) {
	fmt.Println( name, slice )
}

func exercise05() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append( x[:3], x[6:]... )

	printSlice( "y", y )
}

func exercise04() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	printSlice( "x", x )

	x = append(x, 53, 54, 55)
	printSlice( "x", x )
	
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	printSlice( "x", x )
}

func exercise03() {

	slice01 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	printSlice( "slice01", slice01 )

	slice02 := slice01[:5]
	printSlice( "slice02", slice02 )

	slice03 := slice01[5:]	
	printSlice( "slice03", slice03 )

	slice04 := slice01[2:7]
	printSlice( "slice04", slice04 )

	slice05 := slice01[1:6]
	printSlice( "slice05", slice05 )
}

func printSlice( name string, slice []int ) {
	fmt.Println( name, slice )
//	for i, v := range slice {
//		fmt.Printf("%s[%d] = %d\n", name, i, v)
//	}
}

func exercise02() {

	slice01 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range slice01 {
		fmt.Printf("slice01[%d] = %d\n", i, v)
	}
	fmt.Printf("%T\n", slice01)
}

func exercise01() {
	
	arr01 := [5]int{ 1, 2, 3, 4, 5 }

	for i, v := range arr01 {
		fmt.Printf("arr01[%d] = %d\n", i, v)
	}

	fmt.Printf("%T\n", arr01)
}

func exploreTime( time time.Time ) {

	fmt.Printf("Type of time: %T\n", time)

}
