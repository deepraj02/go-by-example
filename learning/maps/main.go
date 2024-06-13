package main

import (
	"fmt"
	"maps"
)

func main() {
	/// To create an empty map, use the builtin make: make(map[key-type]val-type).

	s := make(map[string]int)
	fmt.Println("emp:", s)
	s["Deepraj"] = 1

	fmt.Println("map:", s)

	///Get a value for a key with name[key].

	v1 := s["Deepraj"]
	fmt.Println("v1:", v1)

	///If the key doesn’t exist, the zero value of the value type is returned.

	v3 := s["k3"]
	fmt.Println("v3:", v3)

	///The builtin delete removes key/value pairs from a map.
	delete(s, "Deepraj")
	fmt.Println("map:", s)

	///To remove all key/value pairs from a map, use the clear builtin.
	s["Deepraj"] = 1
	s["k2"] = 2
	s["k3"] = 3
	fmt.Println("map:", s)
	clear(s)
	fmt.Println("map:", s)

	_, prs := s["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	///The maps package contains a number of useful utility functions for maps.

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
