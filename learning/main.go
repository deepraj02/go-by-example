package main

import (
	"fmt"
	"go-by-example/learning/loops"
)

func main() {
	fmt.Println("hello")
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("----------------------------------------------\n\n")
	variables()

	loops.LoopInitialCondition()
	loops.LoopSingleCondition()
	loops.LoopRange()

}

func variables() {
	var a = "hello"
	fmt.Println(a)
	var b int = 2
	fmt.Println(b)

	///Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	const (
		NAME    string = "DEEPRAJ"
		SURNAME string = "\tBaiyda"
	)
	fmt.Printf(NAME + SURNAME)
	fmt.Println("\n\n----------------------------------------------\n\n")

}
