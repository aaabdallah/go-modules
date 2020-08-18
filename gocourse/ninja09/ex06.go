package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println( "Arch:", runtime.GOARCH, "<|> O/S:", runtime.GOOS )

}
