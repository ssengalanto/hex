package hex_test

import (
	"fmt"

	"github.com/ssengalanto/hex"
)

func ExampleFind() {
	fn := func(n int) bool {
		return n < 3
	}
	fmt.Println(hex.Find([]int{3, 2, 1}, fn))
	fmt.Println(hex.Find([]int{3, 4, 5}, fn))
	// Output:
	// 2 true
	// 0 false
}

func ExampleFindIndex() {
	fn := func(n int) bool {
		return n == 3
	}
	fmt.Println(hex.FindIndex([]int{1, 2, 3}, fn))
	fmt.Println(hex.FindIndex([]int{0, 1, 2}, fn))
	// Output:
	// 2
	// -1
}
