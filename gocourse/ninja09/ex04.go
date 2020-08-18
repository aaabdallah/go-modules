package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	goRoutineTotal := 10
	// var wg sync.WaitGroup // same thing as below
	wg := sync.WaitGroup{}

	wg.Add( goRoutineTotal )
	counter := 0

	var mtx sync.Mutex

	for i:=0 ; i<goRoutineTotal ; i++ {
		go func() {
			mtx.Lock()
			myCntr := counter
			myCntr++

			if false { runtime.Gosched() }

			counter = myCntr
			fmt.Printf("%d ", counter)
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("\nFinal Counter:", counter)
}