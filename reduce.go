package ranges

// Reduce produces a singe value from a range by calling a callback to combine the values.
// The seed value will be used as the first value for `a` in `cb(a, b)`.
func Reduce[T, U any](r InputRange[T], cb func(a U, b T) U, seed U) U {
	for !r.Empty() {
		seed = cb(seed, r.Front())
		r.PopFront()
	}

	return seed
}

// ReduceS is `Reduce` accepting a slice.
func ReduceS[T, U any](r []T, cb func(a U, b T) U, seed U) U {
	return Reduce(I(F(SliceRange(r))), cb, seed)
}

// ReduceNoSeed is `Reduce` where the the range is assumed not to be empty, and
// the seed is the front of the range.
//
// Panics when the range is empty.
func ReduceNoSeed[T any](r InputRange[T], cb func(a T, b T) T) T {
	if r.Empty() {
		panic("Cannot reduce an empty range")
	}

	seed := r.Front()
	r.PopFront()

	for !r.Empty() {
		seed = cb(seed, r.Front())

		r.PopFront()
	}

	return seed
}

// ReduceNoSeedS is `ReduceNoSeed` accepting a slice.
func ReduceNoSeedS[T any](r []T, cb func(a T, b T) T) T {
	return ReduceNoSeed(I(F(SliceRange(r))), cb)
}
