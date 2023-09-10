package ranges

// All returns true if `cb` returns `true` for all elements
func All[T any](r InputRange[T], cb func(element T) bool) bool {
	for !r.Empty() {
		if !cb(r.Front()) {
			return false
		}

		r.PopFront()
	}

	return true
}

// AllS is `All` accepting a slice.
func AllS[T any](r []T, cb func(element T) bool) bool {
	return All[T](SliceRange(r), cb)
}

// Any returns true if `cb` returns `true` for any element
func Any[T any](r InputRange[T], cb func(element T) bool) bool {
	for !r.Empty() {
		if cb(r.Front()) {
			return true
		}

		r.PopFront()
	}

	return false
}

// AnyS is `Any` accepting a slice.
func AnyS[T any](r []T, cb func(element T) bool) bool {
	return Any[T](SliceRange(r), cb)
}

// Among returns `true` is `value` is equal to any of the `values` according to `eq`
func Among[T any](eq func(a, b T) bool, value T, values ...T) bool {
	for _, b := range values {
		if eq(value, b) {
			return true
		}
	}

	return false
}

// AmongEq implements `Among` with `==` for any comparable type.
func AmongEq[T comparable](value T, values ...T) bool {
	for _, b := range values {
		if value == b {
			return true
		}
	}

	return false
}

// StartsWith returns `true` if `r1` starts with `r2` according to `cb`
func StartsWith[T any, U any](
	r1 InputRange[T],
	r2 InputRange[U],
	cb func(element1 T, element2 U) bool,
) bool {
	for !r1.Empty() && !r2.Empty() {
		if !cb(r1.Front(), r2.Front()) {
			return false
		}

		r1.PopFront()
		r2.PopFront()
	}

	return r2.Empty()
}

// SkipOver sets `haystack` to the range after `needle`, if and only if
// `haystack` starts with `needle`.
func SkipOver[T any, U any](
	haystack *ForwardRange[T],
	needle InputRange[U],
	cb func(element1 T, element2 U) bool,
) bool {
	if haystack == nil || (*haystack).Empty() {
		return true
	}

	skipped := (*haystack).Save()

	for !skipped.Empty() && !needle.Empty() {
		if !cb(skipped.Front(), needle.Front()) {
			return false
		}

		skipped.PopFront()
		needle.PopFront()
	}

	*haystack = skipped

	return true
}

// Length returns the length of a range in O(n) time.
//
// The range is exhausted.
//
// If your range is a RandomAccessRange, use `r.Len()` instead.
func Length[T any](r InputRange[T]) int {
	count := 0

	for !r.Empty() {
		count++
		r.PopFront()
	}

	return count
}

// Count returns the number of elements where `cb` returns `true`
func Count[T any](r InputRange[T], cb func(element T) bool) int {
	count := 0

	for !r.Empty() {
		if cb(r.Front()) {
			count++
		}

		r.PopFront()
	}

	return count
}

// CountUntil returns the number of elements until `cb` returns `true`
func CountUntil[T any](r InputRange[T], cb func(element T) bool) int {
	count := 0

	for !r.Empty() {
		if cb(r.Front()) {
			return count
		}

		count++
		r.PopFront()
	}

	return count
}
