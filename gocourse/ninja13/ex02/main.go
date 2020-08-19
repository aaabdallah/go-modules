package main

import (
	"fmt"
	"syscs.com/gocourse/ninja13/ex02/word"
	"syscs.com/gocourse/ninja13/ex02/quote"
//	"gocourse/ninja13/ex02/word"
//	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/quote"
//	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/word"
	
)

func main() {
	fmt.Println( "Hello" )
	fmt.Println( word.Count("ab cd") )

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

}
