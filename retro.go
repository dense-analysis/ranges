package ranges

type retroResult[T any] struct{ r BidirectionalRange[T] }

func (r *retroResult[T]) Empty() bool                  { return r.r.Empty() }
func (r *retroResult[T]) Front() T                     { return r.r.Back() }
func (r *retroResult[T]) PopFront()                    { r.r.PopBack() }
func (r *retroResult[T]) Back() T                      { return r.r.Front() }
func (r *retroResult[T]) PopBack()                     { r.r.PopFront() }
func (r *retroResult[T]) Save() ForwardRange[T]        { return &retroResult[T]{r.r.SaveB()} }
func (r *retroResult[T]) SaveB() BidirectionalRange[T] { return &retroResult[T]{r.r.SaveB()} }

type randomAccessRetroResult[T any] struct{ retroResult[T] }

func (r *randomAccessRetroResult[T]) Len() int {
	return r.r.(RandomAccessRange[T]).Len()
}

func (r *randomAccessRetroResult[T]) Get(index int) T {
	random := r.r.(RandomAccessRange[T])

	// Index to the reversed position.
	return random.Get(random.Len() - 1 - index)
}

func (r *randomAccessRetroResult[T]) SaveR() RandomAccessRange[T] {
	return RetroR(r.r.(RandomAccessRange[T]).SaveR())
}

// Retro iterates a BidirectionalRange in reverse.
func Retro[T any](r BidirectionalRange[T]) BidirectionalRange[T] {
	return &retroResult[T]{r}
}

// RetroR is Retro, producing a RandomAccessRange.
func RetroR[T any](r RandomAccessRange[T]) RandomAccessRange[T] {
	return &randomAccessRetroResult[T]{retroResult[T]{r}}
}
