package ranges

type mapResult[T any, U any] struct {
	cb func(element T) U
	ir InputRange[T]
}

func (r *mapResult[T, U]) Empty() bool { return r.ir.Empty() }
func (r *mapResult[T, U]) Front() U    { return r.cb(r.ir.Front()) }
func (r *mapResult[T, U]) PopFront()   { r.ir.PopFront() }

type forwardMapResult[T any, U any] struct {
	mapResult[T, U]
}

func (r *forwardMapResult[T, U]) Save() ForwardRange[U] {
	return &forwardMapResult[T, U]{mapResult[T, U]{r.cb, r.ir.(ForwardRange[T]).Save()}}
}

type bidirectionalMapResult[T any, U any] struct {
	forwardMapResult[T, U]
}

func (r *bidirectionalMapResult[T, U]) Back() U  { return r.cb(r.ir.(BidirectionalRange[T]).Back()) }
func (r *bidirectionalMapResult[T, U]) PopBack() { r.ir.(BidirectionalRange[T]).PopBack() }
func (r *bidirectionalMapResult[T, U]) SaveB() BidirectionalRange[U] {
	return &bidirectionalMapResult[T, U]{
		forwardMapResult[T, U]{mapResult[T, U]{r.cb, r.ir.(BidirectionalRange[T]).SaveB()}},
	}
}

// Map transforms elements from one to another through `cb(element)`
//
// `cb` will be called each time `Front()` is called.
//
// The `Cache` function can be used to cache the result when generating ranges.
func Map[T any, U any](r InputRange[T], cb func(a T) U) InputRange[U] {
	return &mapResult[T, U]{cb, r}
}

// MapF is `Map`, where the position can be saved.
func MapF[T any, U any](r ForwardRange[T], cb func(a T) U) ForwardRange[U] {
	return &forwardMapResult[T, U]{mapResult[T, U]{cb, r}}
}

// MapB is `MapF`, that can be shrunk from the back.
func MapB[T any, U any](r BidirectionalRange[T], cb func(a T) U) BidirectionalRange[U] {
	return &bidirectionalMapResult[T, U]{
		forwardMapResult[T, U]{mapResult[T, U]{cb, r}},
	}
}

// MapS is `MapB` accepting a slice.
func MapS[T any, U any](r []T, cb func(a T) U) BidirectionalRange[U] {
	return MapB[T, U](SliceRange(r), cb)
}
