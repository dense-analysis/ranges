package ranges

// dropResult implements Drop
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

// dropForwardResult implements DropF
type dropForwardResult[T any] struct {
	dropResult[T]
}

func (dr *dropForwardResult[T]) Save() ForwardRange[T] {
	return &dropForwardResult[T]{dropResult[T]{dr.remaining, dr.ir.(ForwardRange[T]).Save()}}
}

// dropBidirectionalResult implements DropB
type dropBidirectionalResult[T any] struct {
	dropForwardResult[T]
}

func (dr *dropBidirectionalResult[T]) PopBack() {
	dr.prime()

	dr.ir.(BidirectionalRange[T]).PopBack()
}

func (dr *dropBidirectionalResult[T]) Back() T {
	dr.prime()

	return dr.ir.(BidirectionalRange[T]).Back()
}

func (dr *dropBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &dropBidirectionalResult[T]{
		dropForwardResult[T]{dropResult[T]{dr.remaining, dr.ir.(BidirectionalRange[T]).SaveB()}},
	}
}

// dropRandomAccessResult implements `DropR`
type dropRandomAccessResult[T any] struct {
	dropBidirectionalResult[T]
}

func (dr *dropRandomAccessResult[T]) Get(index int) T {
	dr.prime()

	return dr.ir.(RandomAccessRange[T]).Get(index)
}

func (dr *dropRandomAccessResult[T]) Len() int {
	dr.prime()

	return dr.ir.(RandomAccessRange[T]).Len()
}

func (dr *dropRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	return &dropRandomAccessResult[T]{
		dropBidirectionalResult[T]{
			dropForwardResult[T]{
				dropResult[T]{dr.remaining, dr.ir.(RandomAccessRange[T]).SaveR()},
			},
		},
	}
}

// Drop creates a range without the first up to `count` elements
func Drop[T any](r InputRange[T], count int) InputRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropResult[T]{count, r}
}

// DropF is `Drop` where the range can be saved.
func DropF[T any](r ForwardRange[T], count int) ForwardRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropForwardResult[T]{dropResult[T]{count, r}}
}

// DropB is `DropF` that can be shrunk from the back.
func DropB[T any](r BidirectionalRange[T], count int) BidirectionalRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropBidirectionalResult[T]{dropForwardResult[T]{dropResult[T]{count, r}}}
}

// DropR is `DropB` permitting random access.
func DropR[T any](r RandomAccessRange[T], count int) RandomAccessRange[T] {
	if count < 0 {
		count = 0
	}

	return &dropRandomAccessResult[T]{
		dropBidirectionalResult[T]{dropForwardResult[T]{dropResult[T]{count, r}}},
	}
}
