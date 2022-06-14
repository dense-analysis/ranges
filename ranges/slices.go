package ranges

type baseSliceRange[T any] struct{ slice []T }

func (r *baseSliceRange[T]) Empty() bool { return r == nil || len(r.slice) == 0 }
func (r *baseSliceRange[T]) PopFront()   { r.slice = r.slice[1:] }

type sliceRange[T any] struct{ baseSliceRange[T] }

func (r *sliceRange[T]) Save() ForwardRange[T] {
	return &sliceRange[T]{baseSliceRange[T]{r.slice}}
}

func (r *sliceRange[T]) Front() T { return r.slice[0] }

// SliceRange creates a range from a slice.
func SliceRange[T any](slice []T) ForwardRange[T] {
	return &sliceRange[T]{baseSliceRange[T]{slice}}
}

type slicePtrRange[T any] struct{ baseSliceRange[T] }

func (r *slicePtrRange[T]) Save() ForwardRange[*T] {
	return &slicePtrRange[T]{baseSliceRange[T]{r.slice}}
}

func (r *slicePtrRange[T]) Front() *T { return &r.slice[0] }

// SlicePtrRange creates a range of pointers to values from a slice.
func SlicePtrRange[T any](slice []T) ForwardRange[*T] {
	return &slicePtrRange[T]{baseSliceRange[T]{slice}}
}

type baseSliceRetroRange[T any] struct{ slice []T }

func (r *baseSliceRetroRange[T]) Empty() bool { return r == nil || len(r.slice) == 0 }
func (r *baseSliceRetroRange[T]) PopFront()   { r.slice = r.slice[:len(r.slice)-1] }

type sliceRetroRange[T any] struct{ baseSliceRetroRange[T] }

func (r *sliceRetroRange[T]) Save() ForwardRange[T] {
	return &sliceRetroRange[T]{baseSliceRetroRange[T]{r.slice}}
}

func (r *sliceRetroRange[T]) Front() T { return r.slice[len(r.slice)-1] }

// SliceRetroRange creates a range from a slice in reverse.
func SliceRetroRange[T any](slice []T) ForwardRange[T] {
	return &sliceRetroRange[T]{baseSliceRetroRange[T]{slice}}
}

type slicePtrRetroRange[T any] struct{ baseSliceRetroRange[T] }

func (r *slicePtrRetroRange[T]) Save() ForwardRange[*T] {
	return &slicePtrRetroRange[T]{baseSliceRetroRange[T]{r.slice}}
}

func (r *slicePtrRetroRange[T]) Front() *T { return &r.slice[len(r.slice)-1] }

// SlicePtrRetroRange creates a range of pointers to values from a slice.
func SlicePtrRetroRange[T any](slice []T) ForwardRange[*T] {
	return &slicePtrRetroRange[T]{baseSliceRetroRange[T]{slice}}
}

// Slice creates a slice of memory from a range.
func Slice[T any](r InputRange[T]) []T {
	slice := make([]T, 0)

	for !r.Empty() {
		slice = append(slice, r.Front())
		r.PopFront()
	}

	return slice
}

// SliceF is `Slice` accepting a ForwardRange.
func SliceF[T any](r ForwardRange[T]) []T {
	return Slice[T](r)
}

// Bytes produces a range over the bytes of string.
func Bytes(s string) ForwardRange[byte] {
	return SliceRange([]byte(s))
}

// Runes produces a range over the runes of string.
func Runes(s string) ForwardRange[rune] {
	return SliceRange([]rune(s))
}

// String creates a new string from a range of runes.
func String(r InputRange[rune]) string {
	return string(Slice(r))
}

// StringF is `String` accepting a ForwardRange.
func StringF(r ForwardRange[rune]) string {
	return string(Slice(I(r)))
}

// StringS is `String` accepting a slice of runes.
//
// This can be used as a callback where `string` cannot be.
func StringS(r []rune) string {
	return string(r)
}
