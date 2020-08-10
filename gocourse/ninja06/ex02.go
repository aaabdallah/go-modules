package main

// import "fmt"

func main() {

	const tries = 100000000

	numbersSlice := []int{49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51}
	//fmt.Println( sum( numbersSlice... ) )
	//fmt.Println( sumOfSlice( numbersSlice ) )

	total := 0
	for i := 0 ; i < tries ; i++ {
		// total = sumCStyle( numbersSlice )
		total = sumGoStyle( numbersSlice )
	}
	total += 1
	//fmt.Println( total )
}

func sum( numbers ...int ) int {

	sum := 0
	for i:=0; i<len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func sumOfSlice( numberSlice []int ) int {

	if numberSlice == nil {
		return 0
	}

	return sum( numberSlice... )
}

func sumCStyle( numberSlice []int ) int {

	total := 0
	for i:=0; i<len(numberSlice); i++ {
		total += numberSlice[i]
	}

	return total
}

func sumGoStyle( numberSlice []int ) int {

	total := 0
	for _, v := range numberSlice {
		total += v
	}

	return total
}
