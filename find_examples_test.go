package hex_test

import (
	"fmt"

	"github.com/ssengalanto/hex"
)

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
