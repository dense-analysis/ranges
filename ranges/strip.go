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

type stripLeftComparableBidirectionalResult[T comparable] struct {
	stripLeftComparableResult[T]
}

func (slr *stripLeftComparableBidirectionalResult[T]) Back() T {
	slr.prime()

	return slr.r.(BidirectionalRange[T]).Back()
}

func (slr *stripLeftComparableBidirectionalResult[T]) PopBack() {
	slr.prime()
	slr.r.(BidirectionalRange[T]).PopBack()
}

func (slr *stripLeftComparableBidirectionalResult[T]) Save() ForwardRange[T] {
	return slr.SaveB()
}

func (slr *stripLeftComparableBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &stripLeftComparableBidirectionalResult[T]{
		stripLeftComparableResult[T]{slr.r.(BidirectionalRange[T]).SaveB(), slr.value, slr.isPrimed},
	}
}

// StripLeftComparable removes elements equal to `value` from the front of a range.
func StripLeftComparable[T comparable](r InputRange[T], value T) InputRange[T] {
	return &stripLeftComparableResult[T]{r, value, false}
}

// StripLeftComparableF is `StripLeftComparable` where the position can be saved.
func StripLeftComparableB[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripLeftComparableBidirectionalResult[T]{
		stripLeftComparableResult[T]{r, value, false},
	}
}

// StripLeftComparableB is `StripLeftComparable` where the range can be shrunk from the back.
func StripLeftComparableF[T comparable](r ForwardRange[T], value T) ForwardRange[T] {
	return &stripLeftComparableForwardResult[T]{
		stripLeftComparableResult[T]{r, value, false},
	}
}

// StripLeftComparableS is `StripLeftComparableB` accepting a slice.
func StripLeftComparableS[T comparable](r []T, value T) ForwardRange[T] {
	return StripLeftComparableB(SliceRange(r), value)
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

type stripLeftBidirectionalResult[T any] struct {
	stripLeftResult[T]
}

func (slr *stripLeftBidirectionalResult[T]) Back() T {
	slr.prime()

	return slr.r.(BidirectionalRange[T]).Back()
}

func (slr *stripLeftBidirectionalResult[T]) PopBack() {
	slr.prime()
	slr.r.(BidirectionalRange[T]).PopBack()
}

func (slr *stripLeftBidirectionalResult[T]) Save() ForwardRange[T] {
	return slr.SaveB()
}

func (slr *stripLeftBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &stripLeftBidirectionalResult[T]{
		stripLeftResult[T]{slr.r.(BidirectionalRange[T]).SaveB(), slr.cb, slr.isPrimed},
	}
}

// StripLeft removes elements where `cb(a) == true` from the front of a range.
func StripLeft[T any](r InputRange[T], cb func(a T) bool) InputRange[T] {
	return &stripLeftResult[T]{r, cb, false}
}

// StripLeftF is `StripLeft` where the position can be saved.
func StripLeftF[T any](r ForwardRange[T], cb func(a T) bool) ForwardRange[T] {
	return &stripLeftForwardResult[T]{stripLeftResult[T]{r, cb, false}}
}

// StripLeftB is `StripLeftF` that can be shrunk from the back.
func StripLeftB[T any](r BidirectionalRange[T], cb func(a T) bool) BidirectionalRange[T] {
	return &stripLeftBidirectionalResult[T]{stripLeftResult[T]{r, cb, false}}
}

// StripLeftS is `StripLeftB` accepting a slice.
func StripLeftS[T any](r []T, cb func(a T) bool) BidirectionalRange[T] {
	return StripLeftB(SliceRange(r), cb)
}
