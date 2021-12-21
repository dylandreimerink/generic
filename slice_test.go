package generic_test

import (
	"fmt"

	"github.com/dylandreimerink/generic"
)

// generics.FilterSlice will leaves the original slice intact
func ExampleFilterSlice() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := generic.FilterSlice(input, func(i int, v int) bool {
		return v >= 5
	})

	fmt.Println("input:", input)
	fmt.Println("output:", output)

	// Output:
	// input: [1 2 3 4 5 6 7 8 9 10]
	// output: [5 6 7 8 9 10]
}

// generics.FilterSliceInPlace modifies the original slice, thus no memory allocation takes place.
// This function doesn't preserve the order of elements inside the slice
func ExampleFilterSliceInPlace() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input:", ints)

	generic.FilterSliceInPlace(&ints, func(v int) bool {
		return v >= 5
	})

	fmt.Println("output:", ints)

	// Output:
	// input: [1 2 3 4 5 6 7 8 9 10]
	// output: [10 9 8 7 5 6]
}
