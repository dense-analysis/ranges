package ranges

// CommonPrefix returns `true` if `r1` starts with `r2` according to `cb`
func CommonPrefix[T any, U any](
	r1 ForwardRange[T],
	r2 InputRange[U],
	cb func(element1 T, element2 U) bool,
) ForwardRange[T] {
	saved := r1.Save()
	count := 0

	for !r1.Empty() && !r2.Empty() && cb(r1.Front(), r2.Front()) {
		r1.PopFront()
		r2.PopFront()
		count++
	}

	return TakeF(saved, count)
}

// CommonPrefixF is `CommonPrefix` accepting a second forward range.
func CommonPrefixF[T any, U any](
	r1 ForwardRange[T],
	r2 ForwardRange[U],
	cb func(element1 T, element2 U) bool,
) ForwardRange[T] {
	return CommonPrefix(r1, I(r2), cb)
}

// CommonPrefixS is `CommonPrefix` accepting slices.
func CommonPrefixS[T any, U any](r1 []T, r2 []U, cb func(element1 T, element2 U) bool) ForwardRange[T] {
	return CommonPrefix(F(SliceRange(r1)), I(F(SliceRange(r2))), cb)
}
