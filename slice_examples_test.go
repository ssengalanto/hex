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

func ExampleMap() {
	fmt.Println(hex.Map([]int{1, 2, 3}, func(el int, i int) int {
		return el * el
	}))
	fmt.Println(hex.Map([]int{104, 101, 108, 108, 111}, func(el int, i int) string {
		return string(rune(el))
	}))
	// Output:
	// [1 4 9]
	// [h e l l o]
}

func ExamplePrepend() {
	fmt.Println(hex.Prepend([]int{1, 2, 3}, 0))
	// Output:
	// [0 1 2 3]
}
