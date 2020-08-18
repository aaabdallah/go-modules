package main

import ( "fmt"; "time"; "runtime" )

func receiveData(ch <-chan int) {
	goInfo()
	fmt.Println( "\t\treceiveData: Starting" )
	fmt.Println( "\t\treceiveData: Received:", <-ch )
	fmt.Println( "\t\treceiveData: Ending" )
}

func sendData(ch chan<- int) {
	goInfo()
	fmt.Println( "\tsendData: Starting" )
	ch <- 10
	fmt.Println( "\tsendData: Sent 10" )
	ch <- 20
	fmt.Println( "\tsendData: Sent 20" )
	fmt.Println( "\tsendData: Ending" )
}

func goInfo() {
	fmt.Print( "goInfo GOARCH:", runtime.GOARCH )
	fmt.Print( " | " )
	fmt.Print( "goInfo GOOS:", runtime.GOOS )
	fmt.Print( " | " )
	fmt.Print( "goInfo NumGoroutine:", runtime.NumGoroutine() )
	fmt.Print( " | " )
	fmt.Print( "goInfo NumCPU:", runtime.NumCPU() )
	fmt.Println()
}

func test2() {

	ch := make( chan int )

	go sendData(ch)

	go receiveData(ch)
	fmt.Println( "Main: received", <-ch)

}

func test1() {
	ch := make(chan int)

	go receiveData(ch)
	go sendData(ch)
	
	fmt.Println( "Main: Sleeping now" )
	time.Sleep( 3 * time.Second )
	fmt.Println( "Main: Waking up" )

	fmt.Println( "Main: received", <-ch)

	fmt.Println( "Main: Sleeping now" )
	time.Sleep( 3 * time.Second )
	fmt.Println( "Main: Waking up" )
}

func main() {
	fmt.Println( "Main: Starting" )

	test2()	
	
	goInfo()

	fmt.Println( "Main: Ending" )
}

