package main

import (
	"time"
)

func main() {

	jobChan := make( chan int )

	go createJobs( jobChan )

	processJobs( jobChan )
}

func createJobs(jobChan chan<- int) {

	for i:=0 ; i<10 ; i++ {
		jobChan <- i
	}

	close(jobChan)

}

func processJobs(jobChan <-chan int) {

	for i:=0 ; i<workers ; i++ {
		go processJob(i, jobChan)
	}

}

func processJob(workerNumber int, jobChan <-chan int) {

	for job, ok := range jobChan {
		if ok == false { return }
		// time.Sleep( 500 * time.Millisecond )
		fmt.Println("Worker", workerNumber, "processed job", job)
	}

}