package ranges

type joinerResult[T any] struct {
	r              InputRange[InputRange[T]]
	separator      ForwardRange[T]
	savedSeparator ForwardRange[T]
	isPrimed       bool
}

func (jr *joinerResult[T]) prime() {
	if !jr.isPrimed {
		for !jr.r.Empty() && jr.r.Front().Empty() {
			jr.r.PopFront()

			if !jr.r.Empty() && !jr.separator.Empty() {
				jr.savedSeparator = jr.separator.Save()
				break
			}
		}

		jr.isPrimed = true
	}
}

func (jr *joinerResult[T]) Empty() bool {
	jr.prime()

	return jr.r.Empty()
}

func (jr *joinerResult[T]) Front() T {
	jr.prime()

	if !jr.savedSeparator.Empty() {
		return jr.savedSeparator.Front()
	}

	return jr.r.Front().Front()
}

func (jr *joinerResult[T]) PopFront() {
	jr.prime()

	if !jr.savedSeparator.Empty() {
		jr.savedSeparator.PopFront()

		if jr.savedSeparator.Empty() {
			jr.isPrimed = false
		}
	} else {
		jr.r.Front().PopFront()
		jr.isPrimed = false
	}
}

type joinerForwardResult[T any] struct {
	r              ForwardRange[ForwardRange[T]]
	separator      ForwardRange[T]
	savedSeparator ForwardRange[T]
	isPrimed       bool
}

func (jr *joinerForwardResult[T]) prime() {
	if !jr.isPrimed {
		for !jr.r.Empty() && jr.r.Front().Empty() {
			jr.r.PopFront()

			if !jr.r.Empty() && !jr.separator.Empty() {
				jr.savedSeparator = jr.separator.Save()
				break
			}
		}

		jr.isPrimed = true
	}
}

func (jr *joinerForwardResult[T]) Empty() bool {
	jr.prime()

	return jr.r.Empty()
}

func (jr *joinerForwardResult[T]) Front() T {
	jr.prime()

	if !jr.savedSeparator.Empty() {
		return jr.savedSeparator.Front()
	}

	return jr.r.Front().Front()
}

func (jr *joinerForwardResult[T]) PopFront() {
	jr.prime()

	if !jr.savedSeparator.Empty() {
		jr.savedSeparator.PopFront()

		if jr.savedSeparator.Empty() {
			jr.isPrimed = false
		}
	} else {
		jr.r.Front().PopFront()
		jr.isPrimed = false
	}
}

func (jfr *joinerForwardResult[T]) Save() ForwardRange[T] {
	return &joinerForwardResult[T]{
		jfr.r.Save(),
		jfr.separator,
		jfr.savedSeparator.Save(),
		jfr.isPrimed,
	}
}

// Joiner joins ranges together using a given ForwardRange as a separator.
func Joiner[T any](r InputRange[InputRange[T]], separator ForwardRange[T]) InputRange[T] {
	return &joinerResult[T]{r, separator, Null[T](), false}
}

// JoinerF is `Joiner`, where the position can be saved.
func JoinerF[T any](r ForwardRange[ForwardRange[T]], separator ForwardRange[T]) ForwardRange[T] {
	return &joinerForwardResult[T]{r, separator, Null[T](), false}
}

// JoinerS is `JoinerF` accepting a slice of ranges.
func JoinerS[T any](r []ForwardRange[T], separator ForwardRange[T]) ForwardRange[T] {
	return JoinerF(SliceRange(r), separator)
}

// JoinerSS is `JoinerF` accepting a slice of lisces.
func JoinerSS[T any](r [][]T, separator []T) ForwardRange[T] {
	return JoinerF(
		CacheF(MapS(r, SliceRange[T])),
		SliceRange(separator),
	)
}

// JoinStrings joins a range of strings together as one string.
func JoinStrings(r ForwardRange[string], separator string) string {
	return string(SliceF(
		JoinerF(
			CacheF(MapF(r, func(x string) ForwardRange[byte] { return Bytes(x) })),
			Bytes(separator),
		),
	))
}
