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

// MapS is `MapF` accepting a slice.
func MapS[T any, U any](r []T, cb func(a T) U) ForwardRange[U] {
	return &forwardMapResult[T, U]{mapResult[T, U]{cb, SliceRange(r)}}
}
