package main

import (
	"fmt"
	"slices"
)

func main() {
	theArrays()
	theSlices()

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

func theSlices() {
	fmt.Println("-----------------\n\nSlices\n\n-----------------")

	///Unlike arrays, slices are typed only by the elements they contain (not the number of elements). An uninitialized slice equals to nil and has length 0.

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 4)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	fmt.Println("set:", s, "get:", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	for i, v := range s {
		fmt.Println("idx:", i, "val:", v)
	}
	/// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	if slices.Equal(c, s) {
		fmt.Println("c and s are equal")
	}
}
