package main

import (
	"fmt"
)

func main() {
	fmt.Println(LoopSingleCondition("B"))
	fmt.Println(LoopInitialCondition(10))
	fmt.Println(LoopRange(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func LoopSingleCondition(ch string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += ch
	}
	return repeated
}

func LoopInitialCondition(number int) int {
	total := 0
	for j := 0; j <= number; j++ {
		total += j
	}
	return total
}

func LoopRange(nums ...int) []int {
	totalNums := []int{}
	/// used with slices and maps
	for _, number := range nums {
		totalNums = append(totalNums, number)
	}
	return totalNums
}
