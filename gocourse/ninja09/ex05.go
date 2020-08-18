package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	goRoutineTotal := 10
	// var wg sync.WaitGroup // same thing as below
	wg := sync.WaitGroup{}

	wg.Add( goRoutineTotal )
	var counter int32

	for i:=0 ; i<goRoutineTotal ; i++ {
		go func() {
			atomic.AddInt32( &counter, 1 )
			runtime.Gosched()

			// since the goroutine gave up its running
			// with the above line, the next line MAY
			// print the "wrong" value after coming back
			// here due to other goroutines incrementing
			// it. HOWEVER, due to the atomic operation
			// all goroutines increment the counter
			// exactly once.
			fmt.Printf("%d ", counter)

			wg.Done()
		}()
	}

	wg.Wait()

	// Since the actual increment is atomic, this
	// will always print the right value regardless
	// of what was printed in the goroutine.
	fmt.Println("\nFinal Counter:", counter)
}