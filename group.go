package ranges

// Group yields pairs of (value, size) from a range where values are equal according to `cb(a, b)`
func Group[T any](r InputRange[T], cb func(a, b T) bool) InputRange[Pair[T, int]] {
	return Map(
		ChunkBy(r, cb),
		func(x InputRange[T]) Pair[T, int] { return MakePair(x.Front(), Length(x)) },
	)
}

// GroupS is `Group` accepting a slice.
func GroupS[T any](r []T, cb func(a, b T) bool) InputRange[Pair[T, int]] {
	return Group(SliceRange(r), cb)
}

// GroupComparable is `Group` where `a == b`
func GroupComparable[T comparable](r InputRange[T]) InputRange[Pair[T, int]] {
	return Group(r, Eq[T])
}

// GroupComparableS is `GroupComparable` accepting a slice.
func GroupComparableS[T comparable](r []T) InputRange[Pair[T, int]] {
	return GroupComparable(SliceRange(r))
}
