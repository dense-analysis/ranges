package ranges

// IsPermutation returns `true` if two ranges are permutations of each other by allocating a temporary map.
func IsPermutation[T comparable](r1 InputRange[T], r2 InputRange[T]) bool {
	// Skip all elements at the front of each range that are equal.
	for !r2.Empty() && !r1.Empty() && r1.Front() == r2.Front() {
		r1.PopFront()
		r2.PopFront()
	}

	if r1.Empty() {
		return r2.Empty()
	}

	counts := make(map[T]int)

	Each(r1, func(v T) {
		count, _ := counts[v]
		counts[v] = count + 1
	})

	return All(
		r2,
		func(v T) bool {
			count, _ := counts[v]

			if count == 0 {
				return false
			} else if count == 1 {
				delete(counts, v)
			} else {
				counts[v] = count - 1
			}

			return true
		},
	) && len(counts) == 0
}

// IsPermutationNoAlloc returns `true` if two ranges are permutations of each
// other in O(m * n) time without allocating any new memory itself.
func IsPermutationNoAlloc[T comparable](r1 ForwardRange[T], r2 ForwardRange[T]) bool {
	// Skip all elements at the front of each range that are equal.
	for !r2.Empty() && !r1.Empty() && r1.Front() == r2.Front() {
		r1.PopFront()
		r2.PopFront()
	}

	if r1.Empty() {
		return r2.Empty()
	}

	if !IsSameLength(r1.Save(), r2.Save()) {
		return false
	}

	for r1Saved := r1.Save(); !r1Saved.Empty(); r1Saved.PopFront() {
		this := r1Saved.Front()
		r1Count := Count(r1.Save(), func(other T) bool { return this == other })
		r2Count := Count(r2.Save(), func(other T) bool { return this == other })

		if r1Count != r2Count {
			return false
		}
	}

	return true
}
