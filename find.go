package ranges

// Find advances a range until cb(a, b) returns `true`, comparing with needle.
func Find[T any](haystack InputRange[T], cb func(a, b T) bool, needle T) InputRange[T] {
	return DropWhile(haystack, func(a T) bool { return !cb(a, needle) })
}

// FindF is `Find` where the range can be saved.
func FindF[T any](haystack ForwardRange[T], cb func(a, b T) bool, needle T) ForwardRange[T] {
	return DropWhileF(haystack, func(a T) bool { return !cb(a, needle) })
}

// FindS is `FindF` accepting a slice.
func FindS[T any](haystack []T, cb func(a, b T) bool, needle T) ForwardRange[T] {
	return FindF(F(B(SliceRange(haystack))), cb, needle)
}

// FindComparable advances a range until Front() == needle
func FindComparable[T comparable](haystack InputRange[T], needle T) InputRange[T] {
	return Find(haystack, Eq[T], needle)
}

// FindComparableF is `FindComparable` where the range can be saved.
func FindComparableF[T comparable](haystack ForwardRange[T], needle T) ForwardRange[T] {
	return FindF(haystack, Eq[T], needle)
}

// FindComparableS is `FindComparableF` accepting a slice.
func FindComparableS[T comparable](haystack []T, needle T) ForwardRange[T] {
	return FindComparableF(F(B(SliceRange(haystack))), needle)
}

// FindEqual advances a range until cb(a, b) returns `true` for all elments of `needle`.
func FindEqual[T any](haystack ForwardRange[T], cb func(a, b T) bool, needle ForwardRange[T]) ForwardRange[T] {
	return DropWhileF(haystack, func(a T) bool { return !StartsWith[T](haystack.Save(), I(needle.Save()), cb) })
}

// FindEqualS is FindEqual accepting a slice.
func FindEqualS[T any](haystack []T, cb func(a, b T) bool, needle ForwardRange[T]) ForwardRange[T] {
	return FindEqual(F(B(SliceRange(haystack))), cb, needle)
}

// FindEqualComparable advances a range until `a == b` is satisifed for all elements of a `needle`.
func FindEqualComparable[T comparable](haystack ForwardRange[T], needle ForwardRange[T]) ForwardRange[T] {
	return DropWhileF(
		haystack,
		func(a T) bool {
			return !StartsWith[T](
				haystack.Save(),
				I(needle.Save()),
				func(a, b T) bool { return a == b },
			)
		},
	)
}

// FindEqualComparableS is FindEqualComparable accepting a slice.
func FindEqualComparableS[T comparable](haystack []T, needle ForwardRange[T]) ForwardRange[T] {
	return FindEqualComparable(F(B(SliceRange(haystack))), needle)
}

// FindAdjacent advances a range until cb(a, b) returns `true` for two adjacent elements.
func FindAdjacent[T any](haystack ForwardRange[T], cb func(a, b T) bool) ForwardRange[T] {
	return DropWhileF(haystack, func(a T) bool {
		saved := haystack.Save()
		saved.PopFront()

		return saved.Empty() || !cb(a, saved.Front())
	})
}

// FindAdjacentS is FindAdjacent accepting a slice.
func FindAdjacentS[T any](haystack []T, cb func(a, b T) bool) ForwardRange[T] {
	return FindAdjacent(F(B(SliceRange(haystack))), cb)
}

// FindAdjacentComparable advances a range until a == b for two adjacent elements.
func FindAdjacentComparable[T comparable](haystack ForwardRange[T]) ForwardRange[T] {
	return FindAdjacent(haystack, Eq[T])
}

// FindAdjacentComparableS is FindAdjacentComparable accepting a slice.
func FindAdjacentComparableS[T comparable](haystack []T) ForwardRange[T] {
	return FindAdjacentComparable(F(B(SliceRange(haystack))))
}

// FindAmong advances until `cb(a, b) == true` for any element of `needle`.
func FindAmong[T any](haystack InputRange[T], cb func(a, b T) bool, needle ForwardRange[T]) InputRange[T] {
	return DropWhile(haystack, func(a T) bool {
		return !Any(I(needle.Save()), func(b T) bool { return cb(a, b) })
	})
}

// FindAmongF is FindAmong where the range can be saved.
func FindAmongF[T any](haystack ForwardRange[T], cb func(a, b T) bool, needle ForwardRange[T]) ForwardRange[T] {
	return DropWhileF(haystack, func(a T) bool {
		return !Any(I(needle.Save()), func(b T) bool { return cb(a, b) })
	})
}

// FindAmongS is FindAmongF accepting a slice.
func FindAmongS[T any](haystack []T, cb func(a, b T) bool, needle ForwardRange[T]) ForwardRange[T] {
	return FindAmongF(F(B(SliceRange(haystack))), cb, needle)
}

// FindAmongComparable advances until `a == b` for any element of `needle`.
func FindAmongComparable[T comparable](haystack InputRange[T], needle ForwardRange[T]) InputRange[T] {
	return FindAmong(haystack, Eq[T], needle)
}

// FindAmongComparableF is FindAmongComparable where the range can be saved.
func FindAmongComparableF[T comparable](haystack ForwardRange[T], needle ForwardRange[T]) ForwardRange[T] {
	return FindAmongF(haystack, Eq[T], needle)
}

// FindAmongComparableS is FindAmongComparableF accepting a slice.
func FindAmongComparableS[T comparable](haystack []T, needle ForwardRange[T]) ForwardRange[T] {
	return FindAmongComparableF(F(B(SliceRange(haystack))), needle)
}

// CanFind returns `true` if `cb(a, b) == true`, comparing with needle.
func CanFind[T any](haystack InputRange[T], cb func(a, b T) bool, needle T) bool {
	return Any(haystack, func(a T) bool { return cb(a, needle) })
}

// CanFindS is `CanFind` accepting a slice.
func CanFindS[T any](haystack []T, cb func(a, b T) bool, needle T) bool {
	return CanFind[T](SliceRange(haystack), cb, needle)
}

// CanFindComparable returns `true` if `a == b`, comparing with needle.
func CanFindComparable[T comparable](haystack InputRange[T], needle T) bool {
	return Any(haystack, func(a T) bool { return a == needle })
}

// CanFindComparableS is CanFindComparable accepting a slice.
func CanFindComparableS[T comparable](haystack []T, needle T) bool {
	return CanFindComparable[T](SliceRange(haystack), needle)
}
