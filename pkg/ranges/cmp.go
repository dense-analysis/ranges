package ranges

// Cmp iterates `r1` and `r2` in lockstep and compares values with the `cmp` function.
//
// The first non-zero result is returned.
//
// -1 is returned if the first range is smaller and all elements are equal.
// 1 is returned if the second range is smaller and all elements are equal.
func Cmp[T any](r1 InputRange[T], r2 InputRange[T], cmp func(a, b T) int) int {
	for !r1.Empty() && !r2.Empty() {
		cmpResult := cmp(r1.Front(), r2.Front())

		if cmpResult != 0 {
			return cmpResult
		}

		r1.PopFront()
		r2.PopFront()
	}

	if !r2.Empty() {
		return -1
	}

	if !r1.Empty() {
		return 1
	}

	return 0
}

// CmpFunc produces a comparison function for any orderable type.
func CmpFunc[T Ordered](a, b T) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}
