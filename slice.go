package generic

// FilterSlice calls the filter function on each element of the input slice
// and will add it to the returned slice if the filter function returned true.
func FilterSlice[T any](input []T, filter func(i int, v T) (keep bool)) []T {
	new := make([]T, 0, len(input))
	for i, v := range input {
		if filter(i, v) {
			new = append(new, v)
		}
	}
	return new
}

// FilterSliceInPlace executes `filter` on every element of `input` and will keep only elements for which true was returned.
// This functions modifies `input`, if this is not desirable, use the FilterSlice function.
// This function doesn't preserve the order of elements
func FilterSliceInPlace[T any](inputPtr *[]T, filter func(v T) (keep bool)) {
	input := *inputPtr
	length := len(input)
	for i := 0; i < length; i++ {
		if filter(input[i]) {
			continue
		}

		for {
			length--

			// If length == i, then there are no more elements after i which are not deleted, break
			if length == i {
				break
			}

			if filter(input[length]) {
				input[i] = input[length]
				break
			}
		}
	}
	*inputPtr = input[:length]
}

// SliceContains returns true if input contains value
func SliceContains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}

	return false
}
