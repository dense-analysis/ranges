package ranges

type uniqResult[T any] struct {
	r  InputRange[T]
	cb func(a, b T) bool
}

func (ur *uniqResult[T]) Empty() bool {
	return ur.r.Empty()
}

func (ur *uniqResult[T]) Front() T {
	return ur.r.Front()
}

func (ur *uniqResult[T]) PopFront() {
	front := ur.Front()
	ur.r.PopFront()

	for !ur.Empty() && ur.cb(front, ur.Front()) {
		ur.r.PopFront()
	}
}

type uniqForwardResult[T any] struct {
	uniqResult[T]
}

func (ufr *uniqForwardResult[T]) Save() ForwardRange[T] {
	return &uniqForwardResult[T]{uniqResult[T]{ufr.r.(ForwardRange[T]).Save(), ufr.cb}}
}

type uniqBidirectionalResult[T any] struct {
	uniqResult[T]
}

func (ubr *uniqBidirectionalResult[T]) Back() T {
	return ubr.r.(BidirectionalRange[T]).Back()
}

func (ubr *uniqBidirectionalResult[T]) PopBack() {
	back := ubr.Back()
	ubr.r.(BidirectionalRange[T]).PopBack()

	for !ubr.Empty() && ubr.cb(ubr.Back(), back) {
		ubr.r.(BidirectionalRange[T]).PopBack()
	}
}

func (ubr *uniqBidirectionalResult[T]) Save() ForwardRange[T] {
	return ubr.SaveB()
}

func (ubr *uniqBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &uniqBidirectionalResult[T]{uniqResult[T]{ubr.r.(BidirectionalRange[T]).SaveB(), ubr.cb}}
}

type uniqComparableResult[T comparable] struct {
	r InputRange[T]
}

func (ur *uniqComparableResult[T]) Empty() bool {
	return ur.r.Empty()
}

func (ur *uniqComparableResult[T]) Front() T {
	return ur.r.Front()
}

func (ur *uniqComparableResult[T]) PopFront() {
	front := ur.Front()
	ur.r.PopFront()

	for !ur.Empty() && front == ur.Front() {
		ur.r.PopFront()
	}
}

type uniqComparableForwardResult[T comparable] struct {
	uniqComparableResult[T]
}

func (ucfr *uniqComparableForwardResult[T]) Save() ForwardRange[T] {
	return &uniqComparableForwardResult[T]{uniqComparableResult[T]{ucfr.r.(ForwardRange[T]).Save()}}
}

type uniqComparableBidirectionalResult[T comparable] struct {
	uniqComparableResult[T]
}

func (ucbr *uniqComparableBidirectionalResult[T]) Back() T {
	return ucbr.r.(BidirectionalRange[T]).Back()
}

func (ucbr *uniqComparableBidirectionalResult[T]) PopBack() {
	back := ucbr.Back()
	ucbr.r.(BidirectionalRange[T]).PopBack()

	for !ucbr.Empty() && ucbr.Back() == back {
		ucbr.r.(BidirectionalRange[T]).PopBack()
	}
}

func (ucbr *uniqComparableBidirectionalResult[T]) Save() ForwardRange[T] {
	return ucbr.SaveB()
}

func (ucbr *uniqComparableBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &uniqComparableBidirectionalResult[T]{uniqComparableResult[T]{ucbr.r.(BidirectionalRange[T]).SaveB()}}
}

// Uniq removes adjacent entries where `cb(a, b) == true`
func Uniq[T any](r InputRange[T], cb func(a, b T) bool) InputRange[T] {
	return &uniqResult[T]{r, cb}
}

// UniqF is `Uniq` where the range can be saved.
func UniqF[T any](r ForwardRange[T], cb func(a, b T) bool) ForwardRange[T] {
	return &uniqForwardResult[T]{uniqResult[T]{r, cb}}
}

// UniqB is `UniqF` that can be shrunk from the back.
func UniqB[T any](r BidirectionalRange[T], cb func(a, b T) bool) BidirectionalRange[T] {
	return &uniqBidirectionalResult[T]{uniqResult[T]{r, cb}}
}

// UniqS is `UniqB` accepting a slice.
func UniqS[T any](r []T, cb func(a, b T) bool) BidirectionalRange[T] {
	return UniqB(SliceRange(r), cb)
}

// UniqComparable removes adjacent entries where `a == b`
func UniqComparable[T comparable](r InputRange[T]) InputRange[T] {
	return &uniqComparableResult[T]{r}
}

// UniqComparableF is `UniqComparable` where the range can be saved.
func UniqComparableF[T comparable](r ForwardRange[T]) ForwardRange[T] {
	return &uniqComparableForwardResult[T]{uniqComparableResult[T]{r}}
}

// UniqComparableB is `UniqComparableF` that can be shrunk from the back.
func UniqComparableB[T comparable](r BidirectionalRange[T]) BidirectionalRange[T] {
	return &uniqComparableBidirectionalResult[T]{uniqComparableResult[T]{r}}
}

// UniqComparableS is `UniqComparableB` accepting a slice.
func UniqComparableS[T comparable](r []T) BidirectionalRange[T] {
	return UniqComparableB(SliceRange(r))
}
