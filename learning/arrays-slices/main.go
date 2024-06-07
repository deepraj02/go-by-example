package main

import (
	"fmt"
)

func main() {
	theArrays()
	
}

func theArrays() {

	/// Can't change the size of an array once it's declared.
	/// dont have access to insert, update functions
	var myArr [10]string
	myArr[0] = "Hello"
	fmt.Println(myArr[0])

	
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100.0
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	/// You can also have the compiler count the number of elements for you with ...

	var b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	/// If you specify the index with :, the elements in between will be zeroed.
	var c = [...]int{100, 3: 400, 500, 600, 700, 9: 1000}
	fmt.Println("idx:", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
