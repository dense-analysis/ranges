package ranges

type dropResult[T any] struct {
	remaining int
	ir        InputRange[T]
}

func (dr *dropResult[T]) prime() {
	for dr.remaining != 0 && !dr.ir.Empty() {
		dr.ir.PopFront()
		dr.remaining--
	}
}

func (dr *dropResult[T]) Empty() bool {
	dr.prime()

	return dr.ir.Empty()
}

func (dr *dropResult[T]) Front() T {
	dr.prime()

	return dr.ir.Front()
}

func (dr *dropResult[T]) PopFront() {
	dr.prime()

	dr.ir.PopFront()
}

type dropForwardResult[T any] struct {
	dropResult[T]
}

func (dr *dropForwardResult[T]) Save() ForwardRange[T] {
	return &dropForwardResult[T]{dropResult[T]{dr.remaining, dr.ir.(ForwardRange[T]).Save()}}
}

// drop creates a range without the first up to `count` elements
func Drop[T any](r InputRange[T], count int) InputRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropResult[T]{count, r}
}

// dropF is `Drop` where the range can be saved.
func DropF[T any](r ForwardRange[T], count int) ForwardRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropForwardResult[T]{dropResult[T]{count, r}}
}
