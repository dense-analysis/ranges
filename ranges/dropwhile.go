package ranges

type dropWhileResult[T any] struct {
	r  InputRange[T]
	cb func(element T) bool
}

func (dwr *dropWhileResult[T]) prime() {
	if dwr.cb != nil {
		for !dwr.r.Empty() {
			front := dwr.r.Front()

			if dwr.cb(front) {
				dwr.r.PopFront()
			} else {
				dwr.cb = nil
				break
			}
		}
	}
}

func (dwr *dropWhileResult[T]) Empty() bool {
	dwr.prime()

	return dwr.r.Empty()
}

func (dwr *dropWhileResult[T]) Front() T {
	dwr.prime()

	return dwr.r.Front()
}

func (dwr *dropWhileResult[T]) PopFront() {
	dwr.prime()

	dwr.r.PopFront()
}

type dropWhileForwardResult[T any] struct {
	dropWhileResult[T]
}

func (dwfr *dropWhileForwardResult[T]) Save() ForwardRange[T] {
	dwfr.prime()

	return &dropWhileForwardResult[T]{dropWhileResult[T]{dwfr.r.(ForwardRange[T]).Save(), dwfr.cb}}
}

// DropWhile advances a range while cb(element) returns `true`
func DropWhile[T any](r InputRange[T], cb func(element T) bool) InputRange[T] {
	if cb == nil {
		panic("cb is nil")
	}

	return &dropWhileResult[T]{r, cb}
}

// DropWhileF is DropWhile where the range can be saved.
func DropWhileF[T any](r ForwardRange[T], cb func(element T) bool) ForwardRange[T] {
	if cb == nil {
		panic("cb is nil")
	}

	return &dropWhileForwardResult[T]{dropWhileResult[T]{r, cb}}
}
