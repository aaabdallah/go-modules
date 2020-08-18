package main

import ( "fmt"; "sync" )

func main() {

	goRoutines := 10

	var wg sync.WaitGroup

	wg.Add(goRoutines)

	for i:=0 ; i<goRoutines ; i++ {
		// if you use i directly, the go routines are rapidly created
		// and each one refers to the outer reference (address) of i
		// which is pointing to a value that has rapidly reached 10...
		// so practically all the goroutines will say 10

		// for this reason, we pass in a copy of i at the time of running the loop
		go func(j int) {
			fmt.Println( "Anonymous function #", j )
			wg.Done()
		}(i)
	}

	wg.Wait()
}