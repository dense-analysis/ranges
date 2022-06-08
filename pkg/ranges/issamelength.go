package ranges

// IsSameLength returns `true` if two ranges have the same length in O(n) time.
func IsSameLength[T any, U any](r1 InputRange[T], r2 InputRange[U]) bool {
	return Length(r1) == Length(r2)
}
