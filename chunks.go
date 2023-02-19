package ranges

type chunkResult[T any] struct {
	size       int
	takeResult takeResult[T]
}

func (cr *chunkResult[T]) Empty() bool {
	return cr.takeResult.ir.Empty()
}

func (cr *chunkResult[T]) Front() InputRange[T] {
	return &cr.takeResult
}

func (cr *chunkResult[T]) PopFront() {
	for !cr.takeResult.Empty() {
		cr.takeResult.PopFront()
	}

	cr.takeResult = takeResult[T]{cr.size, cr.takeResult.ir}
}

type chunkForwardResult[T any] struct {
	size       int
	takeResult takeForwardResult[T]
}

func (cr *chunkForwardResult[T]) Empty() bool {
	return cr.takeResult.ir.Empty()
}

func (cr *chunkForwardResult[T]) Front() ForwardRange[T] {
	return &cr.takeResult
}

func (cr *chunkForwardResult[T]) PopFront() {
	for !cr.takeResult.Empty() {
		cr.takeResult.PopFront()
	}

	cr.takeResult = takeForwardResult[T]{takeResult[T]{cr.size, cr.takeResult.ir}}
}

func (cr *chunkForwardResult[T]) Save() ForwardRange[ForwardRange[T]] {
	return &chunkForwardResult[T]{
		cr.size,
		takeForwardResult[T]{
			takeResult[T]{
				cr.size,
				cr.takeResult.ir.(ForwardRange[T]).Save(),
			},
		},
	}
}

// Chunks returns sub-ranges that are at most `size` in length.
func Chunks[T any](r InputRange[T], size int) InputRange[InputRange[T]] {
	return &chunkResult[T]{size, takeResult[T]{size, r}}
}

// `ChunksF` is `Chunks` where the range can be saved.
func ChunksF[T any](r ForwardRange[T], size int) ForwardRange[ForwardRange[T]] {
	return &chunkForwardResult[T]{
		size,
		takeForwardResult[T]{takeResult[T]{size, r}},
	}
}
