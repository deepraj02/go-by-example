package main

import "fmt"

func main() {
	///Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
	/// args and kwargs in Python

	doSomething(1, 2)
	doSomething(1, 2, 3)
	nums := []int{1, 2, 3, 4}
    doSomething(nums...)
}

func doSomething(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
