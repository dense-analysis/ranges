package ranges

type padRightResult[T any] struct {
	r         InputRange[T]
	value     T
	remaining int
}

func (pr *padRightResult[T]) Empty() bool {
	return pr.r.Empty() && pr.remaining == 0
}

func (pr *padRightResult[T]) Front() T {
	if pr.r.Empty() {
		if pr.remaining == 0 {
			panic("Front() accessed on an empty PadRight result")
		}

		return pr.value
	}

	return pr.r.Front()
}

func (pr *padRightResult[T]) PopFront() {
	if pr.r.Empty() {
		if pr.remaining != 0 {
			pr.remaining--
		} else {
			panic("PopFront() accessed on an empty PadRight result")
		}
	} else {
		if pr.remaining != 0 {
			pr.remaining--
		}

		pr.r.PopFront()
	}
}

type padRightFResult[T any] struct {
	padRightResult[T]
}

func (pr *padRightFResult[T]) Save() ForwardRange[T] {
	return &padRightFResult[T]{padRightResult[T]{pr.r.(ForwardRange[T]).Save(), pr.value, pr.remaining}}
}

// PadRight adds up to `count` `value` elements to the end of a range to ensure it's at least `count` elements long.
func PadRight[T any](r InputRange[T], value T, count int) InputRange[T] {
	if count < 0 {
		count = 0
	}

	return &padRightResult[T]{r, value, count}
}

// PadRightF is `PadRight` returning a ForwardRange
func PadRightF[T any](r ForwardRange[T], value T, count int) ForwardRange[T] {
	if count < 0 {
		count = 0
	}

	return &padRightFResult[T]{padRightResult[T]{r, value, count}}
}

// PadRightS is `PadRightF` accepting a slice
func PadRightS[T any](r []T, value T, count int) ForwardRange[T] {
	return PadRightF(F(SliceRange(r)), value, count)
}
