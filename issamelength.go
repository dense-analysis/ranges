package ranges

// IsSameLength returns `true` if two ranges have the same length in O(n) time.
//
// The ranges are exhausted.
//
// If you have access to random access ranges, use `r1.Len() == r2.Len()` instead.
func IsSameLength[T any, U any](r1 InputRange[T], r2 InputRange[U]) bool {
	return Length(r1) == Length(r2)
}
