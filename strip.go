package ranges

// stripLeftComparableResult implements StripLeftComparable
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

// stripLeftComparableForwardResult implements StripLeftComparableF
type stripLeftComparableForwardResult[T comparable] struct {
	stripLeftComparableResult[T]
}

func (slr *stripLeftComparableForwardResult[T]) Save() ForwardRange[T] {
	return &stripLeftComparableForwardResult[T]{
		stripLeftComparableResult[T]{slr.r.(ForwardRange[T]).Save(), slr.value, slr.isPrimed},
	}
}

// stripLeftComparableBidirectionalResult implements StripLeftComparableB
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

// stripLeftComparableRandomAccessResult implements StripLeftComparableR
type stripLeftComparableRandomAccessResult[T comparable] struct {
	stripLeftComparableBidirectionalResult[T]
}

func (slr *stripLeftComparableRandomAccessResult[T]) Len() int {
	slr.prime()

	return slr.r.(RandomAccessRange[T]).Len()
}

func (slr *stripLeftComparableRandomAccessResult[T]) Get(index int) T {
	slr.prime()

	return slr.r.(RandomAccessRange[T]).Get(index)
}

func (slr *stripLeftComparableRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	slr.prime()

	return &stripLeftComparableRandomAccessResult[T]{
		stripLeftComparableBidirectionalResult[T]{
			stripLeftComparableResult[T]{slr.r.(RandomAccessRange[T]).SaveR(), slr.value, slr.isPrimed},
		},
	}
}

// StripLeftComparable removes elements equal to `value` from the front of a range.
func StripLeftComparable[T comparable](r InputRange[T], value T) InputRange[T] {
	return &stripLeftComparableResult[T]{r, value, false}
}

// StripLeftComparableF is `StripLeftComparable` where the position can be saved.
func StripLeftComparableF[T comparable](r ForwardRange[T], value T) ForwardRange[T] {
	return &stripLeftComparableForwardResult[T]{
		stripLeftComparableResult[T]{r, value, false},
	}
}

// StripLeftComparableB is `StripLeftComparableF` where the range can be shrunk from the back.
func StripLeftComparableB[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripLeftComparableBidirectionalResult[T]{
		stripLeftComparableResult[T]{r, value, false},
	}
}

// StripLeftComparableR is `StripLeftComparableB` with random access.
func StripLeftComparableR[T comparable](r RandomAccessRange[T], value T) RandomAccessRange[T] {
	return &stripLeftComparableRandomAccessResult[T]{
		stripLeftComparableBidirectionalResult[T]{
			stripLeftComparableResult[T]{r, value, false},
		},
	}
}

// StripLeftComparableS is `StripLeftComparableR` accepting a slice.
func StripLeftComparableS[T comparable](r []T, value T) RandomAccessRange[T] {
	return StripLeftComparableR(SliceRange(r), value)
}

// stripLeftResult implements StripLeft
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

// stripLeftForwardResult implements StripLeftF
type stripLeftForwardResult[T any] struct {
	stripLeftResult[T]
}

func (slr *stripLeftForwardResult[T]) Save() ForwardRange[T] {
	return &stripLeftForwardResult[T]{
		stripLeftResult[T]{slr.r.(ForwardRange[T]).Save(), slr.cb, slr.isPrimed},
	}
}

// stripLeftBidirectionalResult implements StripLeftB
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

// stripLeftRandomAccessResult implements StripLeftR
type stripLeftRandomAccessResult[T any] struct {
	stripLeftBidirectionalResult[T]
}

func (slr *stripLeftRandomAccessResult[T]) Len() int {
	slr.prime()

	return slr.r.(RandomAccessRange[T]).Len()
}

func (slr *stripLeftRandomAccessResult[T]) Get(index int) T {
	slr.prime()

	return slr.r.(RandomAccessRange[T]).Get(index)
}

func (slr *stripLeftRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	return &stripLeftRandomAccessResult[T]{
		stripLeftBidirectionalResult[T]{
			stripLeftResult[T]{slr.r.(RandomAccessRange[T]).SaveR(), slr.cb, slr.isPrimed},
		},
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

// StripLeftR is `StripLeftB` with random access.
func StripLeftR[T any](r RandomAccessRange[T], cb func(a T) bool) RandomAccessRange[T] {
	return &stripLeftRandomAccessResult[T]{
		stripLeftBidirectionalResult[T]{stripLeftResult[T]{r, cb, false}},
	}
}

// StripLeftS is `StripLeftR` accepting a slice.
func StripLeftS[T any](r []T, cb func(a T) bool) RandomAccessRange[T] {
	return StripLeftR(SliceRange(r), cb)
}

// stripRightComparableResult implements StripRightComparable
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

// stripRightComparableRandomAccessResult implements StripRightComparableR
type stripRightComparableRandomAccessResult[T comparable] struct {
	stripRightComparableResult[T]
}

func (srr *stripRightComparableRandomAccessResult[T]) Len() int {
	srr.prime()

	return srr.r.(RandomAccessRange[T]).Len()
}

func (srr *stripRightComparableRandomAccessResult[T]) Get(index int) T {
	srr.prime()

	return srr.r.(RandomAccessRange[T]).Get(index)
}

func (srr *stripRightComparableRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	return &stripRightComparableRandomAccessResult[T]{
		stripRightComparableResult[T]{
			srr.r.(RandomAccessRange[T]).SaveR(),
			srr.value,
			srr.isPrimed,
		},
	}
}

// StripRightComparable removes elements equal to `value` from the back of a range.
func StripRightComparable[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripRightComparableResult[T]{r, value, false}
}

// StripRightComparableR is `StripRightComparable` with random access.
func StripRightComparableR[T comparable](r RandomAccessRange[T], value T) RandomAccessRange[T] {
	return &stripRightComparableRandomAccessResult[T]{
		stripRightComparableResult[T]{r, value, false},
	}
}

// StripRightComparableS is `StripRightComparableF` accepting a slice.
func StripRightComparableS[T comparable](r []T, value T) RandomAccessRange[T] {
	return StripRightComparableR(SliceRange(r), value)
}

// stripRightResult implements StripRight
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

// stripRightRandomAccessResult implements StripRightR
type stripRightRandomAccessResult[T any] struct {
	stripRightResult[T]
}

func (srr *stripRightRandomAccessResult[T]) Len() int {
	srr.prime()

	return srr.r.(RandomAccessRange[T]).Len()
}

func (srr *stripRightRandomAccessResult[T]) Get(index int) T {
	srr.prime()

	return srr.r.(RandomAccessRange[T]).Get(index)
}

func (srr *stripRightResult[T]) SaveR() RandomAccessRange[T] {
	return &stripRightRandomAccessResult[T]{
		stripRightResult[T]{srr.r.(RandomAccessRange[T]).SaveR(), srr.cb, srr.isPrimed},
	}
}

// StripRight removes elements where `cb(a) == true` from the back of a range.
func StripRight[T any](r BidirectionalRange[T], cb func(a T) bool) BidirectionalRange[T] {
	return &stripRightResult[T]{r, cb, false}
}

// StripRightR is `StripRight` with random access.
func StripRightR[T any](r RandomAccessRange[T], cb func(a T) bool) RandomAccessRange[T] {
	return &stripRightRandomAccessResult[T]{
		stripRightResult[T]{r, cb, false},
	}
}

// StripRightS is `StripRightR` accepting a slice.
func StripRightS[T any](r []T, cb func(a T) bool) RandomAccessRange[T] {
	return StripRightR(SliceRange(r), cb)
}

// stripComparableResult implements StripComparable
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

type stripComparableRandomAccessResult[T comparable] struct {
	stripComparableResult[T]
}

func (sr *stripComparableRandomAccessResult[T]) Len() int {
	sr.prime()

	return sr.r.(RandomAccessRange[T]).Len()
}

func (sr *stripComparableRandomAccessResult[T]) Get(index int) T {
	sr.prime()

	return sr.r.(RandomAccessRange[T]).Get(index)
}

func (sr *stripComparableRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	return &stripComparableRandomAccessResult[T]{
		stripComparableResult[T]{sr.r.(RandomAccessRange[T]).SaveR(), sr.value, sr.isPrimed},
	}
}

// StripComparable removes elements equal to `value` from the front and back of a range.
func StripComparable[T comparable](r BidirectionalRange[T], value T) BidirectionalRange[T] {
	return &stripComparableResult[T]{r, value, false}
}

// StripComparableR is `StripComparable` with random access.
func StripComparableR[T comparable](r RandomAccessRange[T], value T) RandomAccessRange[T] {
	return &stripComparableRandomAccessResult[T]{
		stripComparableResult[T]{r, value, false},
	}
}

// StripComparableS is `StripComparable` accepting a slice.
func StripComparableS[T comparable](r []T, value T) RandomAccessRange[T] {
	return StripComparableR(SliceRange(r), value)
}

// stripResult implement Strip
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

// stripRandomAccessResult implements StripR
type stripRandomAccessResult[T any] struct {
	stripResult[T]
}

func (sr *stripRandomAccessResult[T]) Get(index int) T {
	sr.prime()

	return sr.r.(RandomAccessRange[T]).Get(index)
}

func (sr *stripRandomAccessResult[T]) Len() int {
	sr.prime()

	return sr.r.(RandomAccessRange[T]).Len()
}

func (sr *stripRandomAccessResult[T]) SaveR() RandomAccessRange[T] {
	return &stripRandomAccessResult[T]{
		stripResult[T]{sr.r.(RandomAccessRange[T]).SaveR(), sr.cb, sr.isPrimed},
	}
}

// Strip removes elements where `cb(a) == true` from the front and back of a range.
func Strip[T any](r BidirectionalRange[T], cb func(a T) bool) BidirectionalRange[T] {
	return &stripResult[T]{r, cb, false}
}

// StripR is `Strip` with random access.
func StripR[T any](r RandomAccessRange[T], cb func(a T) bool) RandomAccessRange[T] {
	return &stripRandomAccessResult[T]{stripResult[T]{r, cb, false}}
}

// StripS is `StripR` accepting a slice.
func StripS[T any](r []T, cb func(a T) bool) RandomAccessRange[T] {
	return StripR(SliceRange(r), cb)
}
