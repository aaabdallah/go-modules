package main

import ( "fmt" 
//W	"sync" 
)

func main() {
	fmt.Println( "MAIN: Starting..." )

	// deadlock()
	// unbufferedChannel()
	// bufferedChannel()
	directionalChannels()

	fmt.Println( "MAIN: Ending..." )
}

func directionalChannels() {

	chPayments := make( chan map[string]float32 )

//W	wg := sync.WaitGroup{}
//W	wg.Add(2)
	go func() {
		sendPayments( chPayments )
//W		wg.Done()
	}()

//W	go func() {
		receivePayments( chPayments )
//W		wg.Done()
//W	}()

//W	wg.Wait()
}

func sendPayments( chPayments chan<- map[string]float32 ) {

	var map01 map[string]float32
	for i:=1 ; i <= 1 ; i++ {
		map01 = map[string]float32{"September":float32(i*1000+100),
			"October":float32(i*1000+100.0), "November":float32(i*1000+100.0)}
		chPayments <- map01
	}

}

func receivePayments( chPayments <-chan map[string]float32 ) {

	map01 := <- chPayments

	fmt.Println( map01 )
}

func bufferedChannel() {

	chInt := make( chan int, 2 ) // allow the channel to buffer two items before blocking

	chInt <- 100
	chInt <- 101

	fmt.Println( "Read:", <-chInt )

}

func unbufferedChannel() {

	chInt := make( chan int )

	go func() {
		fmt.Println( "GOROUTINE: Starting..." )
		fmt.Println( "Read:", <-chInt )
		fmt.Println( "GOROUTINE: Ending..." )
	}()

	chInt <- 100

}

func deadlock() {

	chInt := make( chan int )

	chInt <- 100

	fmt.Println( "Read:", <-chInt )
}

