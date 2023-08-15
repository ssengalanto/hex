package hex_test

import (
	"fmt"

	"github.com/ssengalanto/hex"
)

func ExampleDeref() {
	var strPtr *string = nil
	var val int = 42
	ptr := &val

	fmt.Println(hex.Deref(ptr))
	fmt.Println(hex.Deref(strPtr))
	// Output:
	// 42
	//
}
