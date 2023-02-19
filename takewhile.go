package ranges

type takeWhileResult[T any] struct {
	r        InputRange[T]
	cb       func(element T) bool
	isPrimed bool
}

func (u *takeWhileResult[T]) prime() {
	if !u.isPrimed {
		if !u.r.Empty() {
			front := u.r.Front()

			if !u.cb(front) {
				u.r = Null[T]()
			}
		}

		u.isPrimed = true
	}
}

func (u *takeWhileResult[T]) Empty() bool {
	u.prime()

	return u.r.Empty()
}

func (u *takeWhileResult[T]) Front() T {
	u.prime()

	return u.r.Front()
}

func (u *takeWhileResult[T]) PopFront() {
	u.prime()
	u.r.PopFront()
	u.isPrimed = false
}

type takeWhileForwardResult[T any] struct {
	takeWhileResult[T]
}

func (u *takeWhileForwardResult[T]) Save() ForwardRange[T] {
	return &takeWhileForwardResult[T]{takeWhileResult[T]{u.r.(ForwardRange[T]).Save(), u.cb, u.isPrimed}}
}

// TakeWhile advances a range while `cb` returns `true`
func TakeWhile[T any](r InputRange[T], cb func(element T) bool) InputRange[T] {
	return &takeWhileResult[T]{r, cb, false}
}

// TakeWhileF is `TakeWhile` producing a `ForwardRange`
func TakeWhileF[T any](r ForwardRange[T], cb func(element T) bool) ForwardRange[T] {
	return &takeWhileForwardResult[T]{takeWhileResult[T]{r, cb, false}}
}
