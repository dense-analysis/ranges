package ranges

type filterResult[T any] struct {
	cb       func(element T) bool
	ir       InputRange[T]
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

type filterForwardResult[T any] struct {
	filterResult[T]
}

func (r *filterForwardResult[T]) Save() ForwardRange[T] {
	return &filterForwardResult[T]{filterResult[T]{r.cb, r.ir.(ForwardRange[T]).Save(), r.isPrimed}}
}

// Filter filters down to elements where `cb(element)` returns `true`
func Filter[T any](r InputRange[T], cb func(element T) bool) InputRange[T] {
	return &filterResult[T]{cb, r, false}
}

// FilterF is `Filter` where the range can be saved.
func FilterF[T any](r ForwardRange[T], cb func(element T) bool) ForwardRange[T] {
	return &filterForwardResult[T]{filterResult[T]{cb, r, false}}
}

// FilterS is `FilterF` accepting a slice.
//
// Returns a ForwardRange, which is more efficient when moving forwards.
// `FilterSB` can be advanced in both directions.
func FilterS[T any](slice []T, cb func(element T) bool) ForwardRange[T] {
	return FilterF(F(SliceRange(slice)), cb)
}

type filterBidirectionalResult[T any] struct {
	cb       func(element T) bool
	br       BidirectionalRange[T]
	isPrimed bool
}

func (r *filterBidirectionalResult[T]) prime() {
	if !r.isPrimed {
		for !r.br.Empty() && !r.cb(r.br.Front()) {
			r.br.PopFront()
		}

		for !r.br.Empty() && !r.cb(r.br.Back()) {
			r.br.PopBack()
		}

		r.isPrimed = true
	}
}

func (r *filterBidirectionalResult[T]) Empty() bool {
	r.prime()

	return r.br.Empty()
}

func (r *filterBidirectionalResult[T]) Front() T {
	r.prime()

	return r.br.Front()
}

func (r *filterBidirectionalResult[T]) PopFront() {
	r.prime()

	for {
		r.br.PopFront()

		if r.br.Empty() || r.cb(r.br.Front()) {
			break
		}
	}
}

func (r *filterBidirectionalResult[T]) Back() T {
	r.prime()

	return r.br.Back()
}

func (r *filterBidirectionalResult[T]) PopBack() {
	r.prime()

	for {
		r.br.PopBack()

		if r.br.Empty() || r.cb(r.br.Back()) {
			break
		}
	}
}

func (r *filterBidirectionalResult[T]) Save() ForwardRange[T] {
	return r.SaveB()
}

func (r *filterBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &filterBidirectionalResult[T]{r.cb, r.br.SaveB(), r.isPrimed}
}

// FilterB is `FilterF` that can move in both directions.
//
// This is less efficient for moving forward than `FilterF`, as the filtered
// range must be primed in both directions.
func FilterB[T any](r BidirectionalRange[T], cb func(element T) bool) BidirectionalRange[T] {
	return &filterBidirectionalResult[T]{cb, r, false}
}

// FilterS is `FilterB` accepting a slice.
//
// This is less efficient for moving forward than `FilterS`, as the filtered
// range must be primed in both directions.
func FilterSB[T any](slice []T, cb func(element T) bool) BidirectionalRange[T] {
	return FilterB(SliceRange(slice), cb)
}
