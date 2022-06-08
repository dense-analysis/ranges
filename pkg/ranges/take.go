package ranges

type takeResult[T any] struct {
	remaining int
	ir InputRange[T]
}

func (tr *takeResult[T]) Empty() bool {
	val := tr.remaining == 0 || tr.ir.Empty()

	return val
}

func (tr *takeResult[T]) Front() T {
	val := tr.ir.Front()

	return val
}

func (tr *takeResult[T]) PopFront() {
	tr.ir.PopFront()
	tr.remaining--
}

type takeForwardResult[T any] struct {
	takeResult[T]
}

func (tr *takeForwardResult[T]) Save() ForwardRange[T] {
	return &takeForwardResult[T]{takeResult[T]{tr.remaining, tr.ir.(ForwardRange[T]).Save()}}
}

// Take creates a range of up to `count` elements from the start of a range.
func Take[T any](r InputRange[T], count int) InputRange[T] {
	if count < 0 {
		count = 0
	}

	return &takeResult[T]{count, r}
}

// TakeF is `Take` where the range can be saved.
func TakeF[T any](r ForwardRange[T], count int) ForwardRange[T] {
	if count < 0 {
		count = 0
	}

	return &takeForwardResult[T]{takeResult[T]{count, r}}
}
