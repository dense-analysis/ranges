package ranges

// enumerateResult implements Enumerate
type enumerateResult[T any] struct {
	index int
	r     InputRange[T]
}

func (r *enumerateResult[T]) Empty() bool         { return r.r.Empty() }
func (r *enumerateResult[T]) Front() Pair[int, T] { return MakePair(r.index, r.r.Front()) }
func (r *enumerateResult[T]) PopFront()           { r.r.PopFront(); r.index++ }

// enumerateForwardResult implements EnumerateF
type enumerateForwardResult[T any] struct{ enumerateResult[T] }

func (r *enumerateForwardResult[T]) Save() ForwardRange[Pair[int, T]] {
	return &enumerateForwardResult[T]{
		enumerateResult[T]{r.index, r.r.(ForwardRange[T]).Save()},
	}
}

// Enumerate yields elements with indexes starting from 0
func Enumerate[T any](r InputRange[T]) InputRange[Pair[int, T]] {
	return EnumerateN(r, 0)
}

// EnumerateF is `Enumerate` where the range can be saved.
func EnumerateF[T any](r ForwardRange[T]) ForwardRange[Pair[int, T]] {
	return EnumerateNF(r, 0)
}

// EnumerateN yields elements with indexes starting from n
func EnumerateN[T any](r InputRange[T], n int) InputRange[Pair[int, T]] {
	return &enumerateResult[T]{n, r}
}

// EnumerateNF is `EnumerateN` where the range can be saved.
func EnumerateNF[T any](r ForwardRange[T], n int) ForwardRange[Pair[int, T]] {
	return &enumerateForwardResult[T]{enumerateResult[T]{n, r}}
}
