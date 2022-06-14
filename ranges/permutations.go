package ranges

type permutationsResult[T any] struct {
	r        []T
	c        []int
	i        int
	isPrimed bool
}

func (pr *permutationsResult[T]) Empty() bool {
	return pr.i == len(pr.c)
}

func (pr *permutationsResult[T]) Front() []T {
	if pr.Empty() {
		panic("Front() accessed on an empty range")
	}

	return pr.r
}

func (pr *permutationsResult[T]) PopFront() {
	if pr.Empty() {
		panic("PopFront() accessed on an empty range")
	}

	if pr.isPrimed {
		pr.c[pr.i]++
		pr.i = 0
		pr.isPrimed = false
	}

	for !pr.Empty() {
		if pr.c[pr.i] < pr.i {
			copied := make([]T, len(pr.r))
			copy(copied, pr.r)
			pr.r = copied

			if pr.i%2 == 0 {
				pr.r[0], pr.r[pr.i] = pr.r[pr.i], pr.r[0]
			} else {
				pr.r[pr.c[pr.i]], pr.r[pr.i] = pr.r[pr.i], pr.r[pr.c[pr.i]]
			}

			pr.isPrimed = true

			return
		} else {
			pr.c[pr.i] = 0
			pr.i++
		}
	}
}

func (pr *permutationsResult[T]) Save() ForwardRange[[]T] {
	newResult := permutationsResult[T]{
		make([]T, len(pr.r)),
		make([]int, len(pr.c)),
		pr.i,
		pr.isPrimed,
	}

	copy(newResult.r, pr.r)
	copy(newResult.c, pr.c)

	return &newResult
}

// Permutations returns all permutations of a slice in a forward range.
func Permutations[T any](r []T) ForwardRange[[]T] {
	return &permutationsResult[T]{r, make([]int, len(r)), 0, false}
}
