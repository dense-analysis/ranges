package ranges

type sliceRange[T any] []T

func (r *sliceRange[T]) Empty() bool                  { return r == nil || len(*r) == 0 }
func (r *sliceRange[T]) Front() T                     { return (*r)[0] }
func (r *sliceRange[T]) PopFront()                    { *r = (*r)[1:] }
func (r *sliceRange[T]) Back() T                      { return (*r)[len(*r)-1] }
func (r *sliceRange[T]) PopBack()                     { *r = (*r)[:len(*r)-1] }
func (r *sliceRange[T]) Save() ForwardRange[T]        { return r.SaveR() }
func (r *sliceRange[T]) SaveB() BidirectionalRange[T] { return r.SaveR() }
func (r *sliceRange[T]) Get(index int) T              { return (*r)[index] }
func (r *sliceRange[T]) Len() int                     { return len(*r) }
func (r *sliceRange[T]) SaveR() RandomAccessRange[T]  { return SliceRange([]T(*r)) }

// SliceRange creates a range from a slice.
func SliceRange[T any](slice []T) RandomAccessRange[T] {
	result := sliceRange[T](slice)

	return &result
}

type slicePtrRange[T any] []T

func (r *slicePtrRange[T]) Empty() bool                   { return r == nil || len(*r) == 0 }
func (r *slicePtrRange[T]) Front() *T                     { return &(*r)[0] }
func (r *slicePtrRange[T]) PopFront()                     { *r = (*r)[1:] }
func (r *slicePtrRange[T]) Back() *T                      { return &(*r)[len(*r)-1] }
func (r *slicePtrRange[T]) PopBack()                      { *r = (*r)[:len(*r)-1] }
func (r *slicePtrRange[T]) Save() ForwardRange[*T]        { return r.SaveR() }
func (r *slicePtrRange[T]) SaveB() BidirectionalRange[*T] { return r.SaveR() }
func (r *slicePtrRange[T]) Get(index int) *T              { return &(*r)[index] }
func (r *slicePtrRange[T]) Len() int                      { return len(*r) }
func (r *slicePtrRange[T]) SaveR() RandomAccessRange[*T]  { return SlicePtrRange([]T(*r)) }

// SlicePtrRange creates a range of pointers to values from a slice.
func SlicePtrRange[T any](slice []T) RandomAccessRange[*T] {
	result := slicePtrRange[T](slice)

	return &result
}

// SliceRetroRange creates a range from a slice in reverse.
func SliceRetroRange[T any](slice []T) RandomAccessRange[T] {
	return RetroR(SliceRange(slice))
}

// SlicePtrRetroRange creates a range of pointers to values from a slice in reverse.
func SlicePtrRetroRange[T any](slice []T) RandomAccessRange[*T] {
	return RetroR(SlicePtrRange(slice))
}

// Slice creates a slice of memory from a range.
// If the type of the range at runtime is a RandomAccessRange, the Len() will be used to allocate the slice up-front.
func Slice[T any](r InputRange[T]) []T {
	// If the range is a RandomAccessRange, use the SliceR implementation to allocate up-front.
	if random, ok := r.(RandomAccessRange[T]); ok {
		// See implementation below.
		return SliceR(random)
	}

	slice := make([]T, 0)

	for !r.Empty() {
		slice = append(slice, r.Front())
		r.PopFront()
	}

	return slice
}

// SliceF is `Slice` accepting a ForwardRange.
func SliceF[T any](r ForwardRange[T]) []T {
	return Slice(r)
}

// SliceB is `Slice` accepting a BidirectionalRange.
func SliceB[T any](r BidirectionalRange[T]) []T {
	return Slice(r)
}

// SliceR is `Slice` accepting a RandomAccessRange.
// This variant will allocate the slice with the length of the range without casting the type.
func SliceR[T any](r RandomAccessRange[T]) []T {
	s := make([]T, 0, r.Len())

	for !r.Empty() {
		s = append(s, r.Front())
		r.PopFront()
	}

	return s
}

// Bytes produces a range over the bytes of string.
func Bytes(s string) RandomAccessRange[byte] {
	return SliceRange([]byte(s))
}

// Runes produces a range over the runes of string.
func Runes(s string) RandomAccessRange[rune] {
	return SliceRange([]rune(s))
}

// String creates a new string from a range of runes.
func String(r InputRange[rune]) string {
	return string(Slice(r))
}

// StringF is `String` accepting a ForwardRange.
func StringF(r ForwardRange[rune]) string {
	return string(SliceF(r))
}

// StringB is `String` accepting a BidirectionalRange.
func StringB(r BidirectionalRange[rune]) string {
	return string(SliceB(r))
}

// StringR is `String` accepting a RandomAccessRange.
// This variant is specialized so the rune slice will be allocated with the `Len()` of the range.
func StringR(r RandomAccessRange[rune]) string {
	return string(SliceR(r))
}

// StringS is `String` accepting a slice of runes.
//
// This can be used as a callback where `string` cannot be.
func StringS(r []rune) string {
	return string(r)
}
