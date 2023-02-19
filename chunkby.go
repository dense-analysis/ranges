package ranges

func takeChunks[T any](
	cr *chunkByResult[T],
	takeWhileFunc func(r InputRange[T], cb func(element T) bool) InputRange[T],
) InputRange[T] {
	hasLast := false
	var last T

	cr.takeWhile = takeWhileFunc(cr.r, func(x T) bool {
		if !hasLast {
			hasLast = true
			last = x

			return true
		}

		if cr.cb(last, x) {
			last = x

			return true
		}

		return false
	})

	return cr.takeWhile
}

type chunkByResult[T any] struct {
	r         InputRange[T]
	cb        func(a, b T) bool
	takeWhile InputRange[T]
}

func (cr *chunkByResult[T]) Empty() bool {
	return cr.r.Empty()
}

func (cr *chunkByResult[T]) Front() InputRange[T] {
	return takeChunks(cr, TakeWhile[T])
}

func (cr *chunkByResult[T]) PopFront() {
	if cr.takeWhile == nil {
		cr.Front()
		cr.takeWhile.PopFront()
	}

	Exhaust(cr.takeWhile)
	cr.takeWhile = nil
}

type chunkByForwardResult[T any] struct {
	r         InputRange[T]
	cb        func(a, b T) bool
	takeWhile InputRange[T]
}

// Returns a range that splits a range into subranges when `cb(a, b)` returns `false`.
func ChunkBy[T any](r InputRange[T], cb func(a, b T) bool) InputRange[InputRange[T]] {
	return &chunkByResult[T]{r, cb, nil}
}

// Returns `ChunkBy` with `cb(a) == cb(b)`
func ChunkByValue[T any, U comparable](r InputRange[T], cb func(a T) U) InputRange[InputRange[T]] {
	return ChunkBy(r, func(a, b T) bool { return cb(a) == cb(b) })
}
