package ranges

// EqualComparable returns `true` if two input ranges of `comparable` values are equal.
func EqualComparable[T comparable](r1, r2 InputRange[T]) bool {
	for !r1.Empty() && !r2.Empty() {
		if r1.Front() != r2.Front() {
			return false
		}

		r1.PopFront()
		r2.PopFront()
	}

	return r2.Empty() && r1.Empty()
}

// EqualComparableS returns `true` if two slices of `comparable` values are equal.
func EqualComparableS[T comparable](r1, r2 []T) bool {
	if len(r1) != len(r2) {
		return false
	}

	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

// Equal returns `true` if two input ranges of are equal according to `cb`
func Equal[T any](r1, r2 InputRange[T], cb func(a, b T) bool) bool {
	for !r1.Empty() && !r2.Empty() {
		if !cb(r1.Front(), r2.Front()) {
			return false
		}

		r1.PopFront()
		r2.PopFront()
	}

	return r2.Empty() && r1.Empty()
}

// EqualS returns `true` if two slices of are equal according to `cb`
func EqualS[T any](r1, r2 []T, cb func(a, b T) bool) bool {
	if len(r1) != len(r2) {
		return false
	}

	for i := range r1 {
		if !cb(r1[i], r2[i]) {
			return false
		}
	}

	return true
}
