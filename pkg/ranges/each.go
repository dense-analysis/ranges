package ranges

// Each calls the callback with each element of the range.
func Each[T any](r InputRange[T], cb func(element T)) {
	for !r.Empty() {
		cb(r.Front())

		r.PopFront()
	}
}

// EachS is `Each` accepting a slice.
func EachS[T any](r []T, cb func(element T)) {
	Each(I(SliceRange(r)), cb)
}

// Exhaust steps through every element of a range until the range is empty.
//
// Front() will never be called in the range.
func Exhaust[T any](r InputRange[T]) {
	for !r.Empty() {
		r.PopFront()
	}
}
