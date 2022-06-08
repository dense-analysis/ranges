package ranges

type teeResult[T any] struct {
	input  InputRange[T]
	output OutputRange[T]
}

func (tr *teeResult[T]) Empty() bool {
	return tr.input.Empty()
}

func (tr *teeResult[T]) Front() T {
	return tr.input.Front()
}

func (tr *teeResult[T]) PopFront() {
	tr.output.Put(tr.input.Front())
	tr.input.PopFront()
}

// Tee produces a range that outputs to a given 'output' when elements are popped.
func Tee[T any](r InputRange[T], output OutputRange[T]) InputRange[T] {
	return &teeResult[T]{r, output}
}

// TeeS is `Tee` accepting a slice.
func TeeS[T any](r []T, output OutputRange[T]) InputRange[T] {
	return &teeResult[T]{SliceRange(r), output}
}
