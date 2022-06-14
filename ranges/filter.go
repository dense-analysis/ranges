package ranges

type filterResult[T any] struct {
	cb func(element T) bool
	ir InputRange[T]
	isPrimed bool
}

func (r *filterResult[T]) prime() {
	if !r.isPrimed {
		for !r.ir.Empty() && !r.cb(r.ir.Front()) {
			r.ir.PopFront()
		}

		r.isPrimed = true
	}
}

func (r *filterResult[T]) Empty() bool {
	r.prime()

	return r.ir.Empty()
}

func (r *filterResult[T]) Front() T {
	r.prime()

	return r.ir.Front()
}

func (r *filterResult[T]) PopFront() {
	r.prime()

	for {
		r.ir.PopFront()

		if r.ir.Empty() || r.cb(r.ir.Front()) {
			break
		}
	}
}

type forwardFilterResult[T any] struct {
	filterResult[T]
}

func (r *forwardFilterResult[T]) Save() ForwardRange[T] {
	return &forwardFilterResult[T]{filterResult[T]{r.cb, r.ir.(ForwardRange[T]).Save(), r.isPrimed}}
}

// Filter filters down to elements where `cb(element)` returns `true`
func Filter[T any](r InputRange[T], cb func(element T) bool) InputRange[T] {
	return &filterResult[T]{cb, r, false}
}

// FilterF is `Filter` where the range can be saved.
func FilterF[T any](r ForwardRange[T], cb func(element T) bool) ForwardRange[T] {
	return &forwardFilterResult[T]{filterResult[T]{cb, r, false}}
}

// FilterS is `FilterF` accepting a slice.
func FilterS[T any](slice []T, cb func(element T) bool) ForwardRange[T] {
	return FilterF(SliceRange(slice), cb)
}
