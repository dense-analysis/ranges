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

// Uniq removes adjacent entries where `cb(a, b) == true`
func Uniq[T any](r InputRange[T], cb func(a, b T) bool) InputRange[T] {
	return &uniqResult[T]{r, cb}
}

// UniqF is `Uniq` where the range can be saved.
func UniqF[T any](r ForwardRange[T], cb func(a, b T) bool) ForwardRange[T] {
	return &uniqForwardResult[T]{uniqResult[T]{r, cb}}
}

// UniqS is `UniqF` accepting a slice.
func UniqS[T any](r []T, cb func(a, b T) bool) ForwardRange[T] {
	return UniqF(SliceRange(r), cb)
}

// UniqComparable removes adjacent entries where `a == b`
func UniqComparable[T comparable](r InputRange[T]) InputRange[T] {
	return &uniqComparableResult[T]{r}
}

// UniqComparableF is `UniqComparable` where the range can be saved.
func UniqComparableF[T comparable](r ForwardRange[T]) ForwardRange[T] {
	return &uniqComparableForwardResult[T]{uniqComparableResult[T]{r}}
}

// UniqComparableS is `UniqComparableF` accepting a slice.
func UniqComparableS[T comparable](r []T) ForwardRange[T] {
	return UniqComparableF(SliceRange(r))
}
