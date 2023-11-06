package main

import "fmt"

func main() {
	testMessArr := []string{
		"testMess01",
		"testMess02",
		"testMess03",
	}
	numLengthMessArr := float64(len(testMessArr))
	testNum01 := 12.0

	totalCost := numLengthMessArr + testNum01
	fmt.Println("the toast value testing: ", totalCost)
}
