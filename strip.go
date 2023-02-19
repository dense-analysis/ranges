package ranges

type stripLeftComparableResult[T comparable] struct {
	r        InputRange[T]
	value    T
	isPrimed bool
}

func (slr *stripLeftComparableResult[T]) prime() {
	if !slr.isPrimed {
		for !slr.r.Empty() && slr.value == slr.r.Front() {
			slr.r.PopFront()
		}

		slr.isPrimed = true
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
	if !slr.isPrimed {
		for !slr.r.Empty() && slr.cb(slr.r.Front()) {
			slr.r.PopFront()
		}

		slr.isPrimed = true
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

type stripRightComparableResult[T comparable] struct {
	r        BidirectionalRange[T]
	value    T
	isPrimed bool
}

func (srr *stripRightComparableResult[T]) prime() {
	if !srr.isPrimed {
		for !srr.r.Empty() && srr.value == srr.r.Back() {
			srr.r.PopBack()
		}

		srr.isPrimed = true
	}
}

func (srr *stripRightComparableResult[T]) Empty() bool {
	srr.prime()

	return srr.r.Empty()
}

func (srr *stripRightComparableResult[T]) Front() T {
	srr.prime()

	return srr.r.Front()
}

func (srr *stripRightComparableResult[T]) PopFront() {
	srr.prime()
	srr.r.PopFront()
}

func (srr *stripRightComparableResult[T]) Back() T {
	srr.prime()

	return srr.r.Back()
}

func (srr *stripRightComparableResult[T]) PopBack() {
	srr.prime()
	srr.r.PopBack()
}

func (srr *stripRightComparableResult[T]) Save() ForwardRange[T] {
	return srr.SaveB()
}

func (srr *stripRightComparableResult[T]) SaveB() BidirectionalRange[T] {
	return &stripRightComparableResult[T]{srr.r.SaveB(), srr.value, srr.isPrimed}
}

// StripRightComparable removes elements equal to `value` from the back of a range.
func StripRightComparable[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripRightComparableResult[T]{r, value, false}
}

// StripRightComparableS is `StripRightComparable` accepting a slice.
func StripRightComparableS[T comparable](r []T, value T) BidirectionalRange[T] {
	return StripRightComparable(SliceRange(r), value)
}

type stripRightResult[T any] struct {
	r        BidirectionalRange[T]
	cb       func(a T) bool
	isPrimed bool
}

func (srr *stripRightResult[T]) prime() {
	if !srr.isPrimed {
		for !srr.r.Empty() && srr.cb(srr.r.Back()) {
			srr.r.PopBack()
		}

		srr.isPrimed = true
	}
}

func (srr *stripRightResult[T]) Empty() bool {
	srr.prime()

	return srr.r.Empty()
}

func (srr *stripRightResult[T]) Front() T {
	srr.prime()

	return srr.r.Front()
}

func (srr *stripRightResult[T]) PopFront() {
	srr.prime()
	srr.r.PopFront()
}

func (srr *stripRightResult[T]) Back() T {
	srr.prime()

	return srr.r.Back()
}

func (srr *stripRightResult[T]) PopBack() {
	srr.prime()
	srr.r.PopBack()
}

func (srr *stripRightResult[T]) Save() ForwardRange[T] {
	return srr.SaveB()
}

func (srr *stripRightResult[T]) SaveB() BidirectionalRange[T] {
	return &stripRightResult[T]{srr.r.SaveB(), srr.cb, srr.isPrimed}
}

// StripRight removes elements where `cb(a) == true` from the back of a range.
func StripRight[T any](r BidirectionalRange[T], cb func(a T) bool) BidirectionalRange[T] {
	return &stripRightResult[T]{r, cb, false}
}

// StripRightS is `StripRight` accepting a slice.
func StripRightS[T any](r []T, cb func(a T) bool) BidirectionalRange[T] {
	return StripRight(SliceRange(r), cb)
}

type stripComparableResult[T comparable] struct {
	r        BidirectionalRange[T]
	value    T
	isPrimed bool
}

func (sr *stripComparableResult[T]) prime() {
	if !sr.isPrimed {
		for !sr.r.Empty() && sr.value == sr.r.Front() {
			sr.r.PopFront()
		}

		for !sr.r.Empty() && sr.value == sr.r.Back() {
			sr.r.PopBack()
		}

		sr.isPrimed = true
	}
}

func (sr *stripComparableResult[T]) Empty() bool {
	sr.prime()

	return sr.r.Empty()
}

func (sr *stripComparableResult[T]) Front() T {
	sr.prime()

	return sr.r.Front()
}

func (sr *stripComparableResult[T]) PopFront() {
	sr.prime()
	sr.r.PopFront()
}

func (sr *stripComparableResult[T]) Back() T {
	sr.prime()

	return sr.r.Back()
}

func (sr *stripComparableResult[T]) PopBack() {
	sr.prime()
	sr.r.PopBack()
}

func (sr *stripComparableResult[T]) Save() ForwardRange[T] {
	return sr.SaveB()
}

func (sr *stripComparableResult[T]) SaveB() BidirectionalRange[T] {
	return &stripComparableResult[T]{sr.r.SaveB(), sr.value, sr.isPrimed}
}

// StripComparable removes elements equal to `value` from the front and back of a range.
func StripComparable[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripComparableResult[T]{r, value, false}
}

// StripComparableS is `StripComparable` accepting a slice.
func StripComparableS[T comparable](r []T, value T) BidirectionalRange[T] {
	return StripComparable(SliceRange(r), value)
}

type stripResult[T any] struct {
	r        BidirectionalRange[T]
	cb       func(a T) bool
	isPrimed bool
}

func (sr *stripResult[T]) prime() {
	if !sr.isPrimed {
		for !sr.r.Empty() && sr.cb(sr.r.Front()) {
			sr.r.PopFront()
		}

		for !sr.r.Empty() && sr.cb(sr.r.Back()) {
			sr.r.PopBack()
		}

		sr.isPrimed = true
	}
}

func (sr *stripResult[T]) Empty() bool {
	sr.prime()

	return sr.r.Empty()
}

func (sr *stripResult[T]) Front() T {
	sr.prime()

	return sr.r.Front()
}

func (sr *stripResult[T]) PopFront() {
	sr.prime()
	sr.r.PopFront()
}

func (sr *stripResult[T]) Back() T {
	sr.prime()

	return sr.r.Back()
}

func (sr *stripResult[T]) PopBack() {
	sr.prime()
	sr.r.PopBack()
}

func (sr *stripResult[T]) Save() ForwardRange[T] {
	return sr.SaveB()
}

func (sr *stripResult[T]) SaveB() BidirectionalRange[T] {
	return &stripResult[T]{sr.r.SaveB(), sr.cb, sr.isPrimed}
}

// Strip removes elements where `cb(a) == true` from the front and back of a range.
func Strip[T any](r BidirectionalRange[T], cb func(a T) bool) BidirectionalRange[T] {
	return &stripResult[T]{r, cb, false}
}

// StripS is `Strip` accepting a slice.
func StripS[T any](r []T, cb func(a T) bool) BidirectionalRange[T] {
	return Strip(SliceRange(r), cb)
}
