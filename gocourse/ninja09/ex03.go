package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	goRoutineTotal := 10
	wg := sync.WaitGroup{}

	wg.Add( goRoutineTotal )
	counter := 0

	for i:=0 ; i<goRoutineTotal ; i++ {
		go func() {
			myCntr := counter
			myCntr++

			runtime.Gosched()

			counter = myCntr
			fmt.Printf("%d ", counter)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("\nFinal Counter: ", counter)
}