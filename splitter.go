package ranges

type splitterState uint8

const (
	splitterInitialState splitterState = 0
	splitterIsPrimed     splitterState = 1 << iota
	splitterIsEmpty
	splitterHasFinalSplit
)

type splitterResult[T any] struct {
	r         ForwardRange[T]
	separator ForwardRange[T]
	current   ForwardRange[T]
	cb        func(a, b T) bool
	state     splitterState
}

func (sr *splitterResult[T]) Empty() bool {
	if (sr.state & splitterHasFinalSplit) != 0 {
		return false
	}

	if (sr.state & splitterIsEmpty) != 0 {
		return true
	}

	if (sr.state & splitterIsPrimed) == 0 {
		sr.current = TakeWhileF(sr.r, func(a T) bool {
			saved := sr.r.Save()

			if StartsWith[T](saved, I(sr.separator.Save()), sr.cb) {
				sr.r = saved

				if saved.Empty() {
					sr.state |= splitterHasFinalSplit
				}

				return false
			}

			return true
		})

		sr.state |= splitterIsPrimed
	}

	return false
}

func (sr *splitterResult[T]) Front() ForwardRange[T] {
	if sr.Empty() {
		panic("Front() called on an empty Splitter() range")
	}

	return sr.current
}

func (sr *splitterResult[T]) PopFront() {
	if (sr.state & splitterHasFinalSplit) != 0 {
		sr.state ^= splitterHasFinalSplit

		return
	}

	if sr.Empty() {
		panic("PopFront() called on an empty Splitter() range")
	}

	for !sr.current.Empty() {
		sr.current.PopFront()
	}

	sr.state ^= splitterIsPrimed

	if sr.r.Empty() {
		sr.state ^= splitterIsEmpty
	}
}

func (sr *splitterResult[T]) Save() ForwardRange[ForwardRange[T]] {
	return &splitterResult[T]{sr.r.Save(), sr.separator, sr.current.Save(), sr.cb, sr.state}
}

type splitWhenResult[T any] struct {
	r       ForwardRange[T]
	current ForwardRange[T]
	cb      func(a, b T) bool
	state   splitterState
}

func (swr *splitWhenResult[T]) Empty() bool {
	if (swr.state & splitterHasFinalSplit) != 0 {
		return false
	}

	if (swr.state & splitterIsEmpty) != 0 {
		return true
	}

	if (swr.state & splitterIsPrimed) == 0 {
		swr.current = TakeWhileF(swr.r, func(a T) bool {
			/*
				if StartsWith[T](saved, I(swr.separator.Save()), sr.cb) {
					swr.r = saved

					if saved.Empty() {
						swr.state |= splitterHasFinalSplit
					}

					return false
				}
			*/

			return true
		})

		swr.state |= splitterIsPrimed
	}

	return false
}

func (swr *splitWhenResult[T]) Front() ForwardRange[T] {
	if swr.Empty() {
		panic("Front() called on an empty SplitWhen() range")
	}

	return swr.current
}

func (swr *splitWhenResult[T]) PopFront() {
	if (swr.state & splitterHasFinalSplit) != 0 {
		swr.state ^= splitterHasFinalSplit

		return
	}

	if swr.Empty() {
		panic("PopFront() called on an empty SplitWhen() range")
	}

	for !swr.current.Empty() {
		swr.current.PopFront()
	}

	swr.state ^= splitterIsPrimed

	if swr.r.Empty() {
		swr.state ^= splitterIsEmpty
	}
}

func (swr *splitWhenResult[T]) Save() ForwardRange[ForwardRange[T]] {
	return &splitWhenResult[T]{swr.r.Save(), swr.current.Save(), swr.cb, swr.state}
}

// Splitter splits a range using a ForwardRange as a separator
// where `cb(a, b) == true` on each element.
func Splitter[T any](r ForwardRange[T], cb func(a, b T) bool, separator ForwardRange[T]) ForwardRange[ForwardRange[T]] {
	return &splitterResult[T]{r, separator, Null[T](), cb, splitterInitialState}
}

// SplitterS is `Splitter` accepting a slice.
func SplitterS[T any](r []T, cb func(a, b T) bool, separator ForwardRange[T]) ForwardRange[ForwardRange[T]] {
	return Splitter(F(B(SliceRange(r))), cb, separator)
}

// SplitterSS is `SplitterS` accepting a slice for both ranges.
func SplitterSS[T any](r []T, cb func(a, b T) bool, separator []T) ForwardRange[ForwardRange[T]] {
	return SplitterS(r, cb, F(B(SliceRange(separator))))
}

// SplitterComparable splits ranges using a ForwardRange as a separator where
// each element is equal.
func SplitterComparable[T comparable](r ForwardRange[T], separator ForwardRange[T]) ForwardRange[ForwardRange[T]] {
	return Splitter(r, Eq[T], separator)
}

// SplitterComparableS is `SplitterComparable` accepting a slice.
func SplitterComparableS[T comparable](r []T, separator ForwardRange[T]) ForwardRange[ForwardRange[T]] {
	return SplitterComparable(F(B(SliceRange(r))), separator)
}

// SplitterComparableSS is `SplitterComparable` accepting a slice for both ranges
func SplitterComparableSS[T comparable](r []T, separator []T) ForwardRange[ForwardRange[T]] {
	return SplitterComparableS(r, F(B(SliceRange(separator))))
}

// SplitString splits a string by a `separator` into a range of strings.
func SplitString(r string, separator string) ForwardRange[string] {
	return CacheF(
		MapF(
			SplitterComparable(F(B(Bytes(r))), F(B(Bytes(separator)))),
			Pipe2(SliceF[byte], func(arr []byte) string { return string(arr) }),
		),
	)
}

// SplitWhen splits a range where `cb(a, b) == true` for adjacent elements.
func SplitWhen[T any](r InputRange[T], cb func(a, b T) bool) InputRange[InputRange[T]] {
	return ChunkBy(r, func(a, b T) bool { return !cb(a, b) })
}
