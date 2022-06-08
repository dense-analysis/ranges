package ranges

type stripLeftComparableResult[T comparable] struct {
	r        InputRange[T]
	value    T
	isPrimed bool
}

func (slr *stripLeftComparableResult[T]) prime() {
	for !slr.isPrimed && !slr.r.Empty() {
		if slr.value == slr.r.Front() {
			slr.r.PopFront()
		} else {
			slr.isPrimed = true
		}
	}
}

func (slr *stripLeftComparableResult[T]) Empty() bool {
	slr.prime()

	return slr.r.Empty()
}

func (slr *stripLeftComparableResult[T]) Front() T {
	slr.prime()

	return slr.r.Front()
}

func (slr *stripLeftComparableResult[T]) PopFront() {
	slr.prime()
	slr.r.PopFront()
}

type stripLeftComparableForwardResult[T comparable] struct {
	stripLeftComparableResult[T]
}

func (slr *stripLeftComparableForwardResult[T]) Save() ForwardRange[T] {
	return &stripLeftComparableForwardResult[T]{
		stripLeftComparableResult[T]{slr.r.(ForwardRange[T]).Save(), slr.value, slr.isPrimed},
	}
}

// StripLeftComparable removes elements equal to `value` from the front of a range.
func StripLeftComparable[T comparable](r InputRange[T], value T) InputRange[T] {
	return &stripLeftComparableResult[T]{r, value, false}
}

// StripLeftComparableF is `StripLeftComparable`, where the position can be saved.
func StripLeftComparableF[T comparable](r ForwardRange[T], value T) ForwardRange[T] {
	return &stripLeftComparableForwardResult[T]{
		stripLeftComparableResult[T]{r, value, false},
	}
}

// StripLeftComparableS is `StripLeftComparableF` accepting a slice.
func StripLeftComparableS[T comparable](r []T, value T) ForwardRange[T] {
	return StripLeftComparableF(SliceRange(r), value)
}

type stripLeftResult[T any] struct {
	r        InputRange[T]
	cb       func(a T) bool
	isPrimed bool
}

func (slr *stripLeftResult[T]) prime() {
	for !slr.isPrimed && !slr.r.Empty() {
		if slr.cb(slr.r.Front()) {
			slr.r.PopFront()
		} else {
			slr.isPrimed = true
		}
	}
}

func (slr *stripLeftResult[T]) Empty() bool {
	slr.prime()

	return slr.r.Empty()
}

func (slr *stripLeftResult[T]) Front() T {
	slr.prime()

	return slr.r.Front()
}

func (slr *stripLeftResult[T]) PopFront() {
	slr.prime()
	slr.r.PopFront()
}

type stripLeftForwardResult[T any] struct {
	stripLeftResult[T]
}

func (slr *stripLeftForwardResult[T]) Save() ForwardRange[T] {
	return &stripLeftForwardResult[T]{
		stripLeftResult[T]{slr.r.(ForwardRange[T]).Save(), slr.cb, slr.isPrimed},
	}
}

// StripLeft removes elements where `cb(a) == true` from the front of a range.
func StripLeft[T comparable](r InputRange[T], cb func(a T) bool) InputRange[T] {
	return &stripLeftResult[T]{r, cb, false}
}

// StripLeftF is `StripLeft`, where the position can be saved.
func StripLeftF[T comparable](r ForwardRange[T], cb func(a T) bool) ForwardRange[T] {
	return &stripLeftForwardResult[T]{
		stripLeftResult[T]{r, cb, false},
	}
}

// StripLeftS is `StripLeftF` accepting a slice.
func StripLeftS[T comparable](r []T, cb func(a T) bool) ForwardRange[T] {
	return StripLeftF(SliceRange(r), cb)
}
