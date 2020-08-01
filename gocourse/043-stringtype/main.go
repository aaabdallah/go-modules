package main

import "fmt"

func main() {

	es := "This is a purely English sentence."
	//as := "هذه جملة عربيةز"
	as := "ش𡁯"

	reportOnString( es )

	reportOnString( as )
}

func reportOnString(s string) {
	fmt.Println("### BEGIN ###")
	fmt.Println( s )
	fmt.Println( "Length in bytes:", len(s) )
	fmt.Println( "Length in runes:", len( []rune(s) ) )
	fmt.Printf("%q\n", s)
	fmt.Println( []byte(s) )
	fmt.Printf("%x\n", s)
	fmt.Println("About to print as UTF8 string")
	fmt.Printf("%#U\n", s)
	for i:=0 ; i<len(s); i++ {
		fmt.Printf("%U ", s[i])
	}
	fmt.Println()
	fmt.Println("### END ###")
	fmt.Println()
}