package main

func sumGoStyle( numberSlice []int ) int {

	total := 0
	for _, v := range numberSlice {
		total += v
	}

	return total
}

func main() {

	const tries = 100000000

	numbersSlice := []int{49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51,49,50,51}

	total := 0
	for i := 0 ; i < tries ; i++ {
		total = sumGoStyle( numbersSlice )
	}
	total += 1
}

