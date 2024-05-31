package loops

import (
	"fmt"
)

func LoopSingleCondition() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}

func LoopInitialCondition() {
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
}	

func LoopRange() {

	/// used with slices and maps
	for _, num := range []int{1, 2, 3} {
		fmt.Println(num)
	}
}