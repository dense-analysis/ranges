package ranges

type strideResult[T any] struct {
	r    InputRange[T]
	size int
}

func (sr *strideResult[T]) Empty() bool {
	return sr.r.Empty()
}

func (sr *strideResult[T]) Front() T {
	return sr.r.Front()
}

func (sr *strideResult[T]) PopFront() {
	sr.r.PopFront()

	for i := 1; i < sr.size && !sr.r.Empty(); i++ {
		sr.r.PopFront()
	}
}

type strideForwardResult[T any] struct {
	strideResult[T]
}

func (sfr *strideForwardResult[T]) Save() ForwardRange[T] {
	return &strideForwardResult[T]{
		strideResult[T]{sfr.r.(ForwardRange[T]).Save(), sfr.size},
	}
}

type strideSliceResult[T any] struct {
	slice []T
	size  int
}

func (ssr *strideSliceResult[T]) Empty() bool {
	return len(ssr.slice) == 0
}

func (ssr *strideSliceResult[T]) Front() T {
	return ssr.slice[0]
}

func (ssr *strideSliceResult[T]) PopFront() {
	if ssr.size > len(ssr.slice) {
		ssr.slice = nil
	} else {
		ssr.slice = ssr.slice[ssr.size:]
	}
}

func (ssr *strideSliceResult[T]) Save() ForwardRange[T] {
	return &strideSliceResult[T]{ssr.slice, ssr.size}
}

// Stride produces every `step` elements in a range.
func Stride[T any](r InputRange[T], step int) InputRange[T] {
	if step < 1 {
		panic("step < 1 for Stride")
	}

	return &strideResult[T]{r, step}
}

// StrideF is `Stride` producing a `ForwardRange`
func StrideF[T any](r ForwardRange[T], step int) ForwardRange[T] {
	if step < 1 {
		panic("step < 1 for Stride")
	}

	return &strideForwardResult[T]{strideResult[T]{r, step}}
}

// StrideS is `StrideF` accepting a slice.
//
// This result is optimized for slices.
func StrideS[T any](r []T, step int) ForwardRange[T] {
	if step < 1 {
		panic("step < 1 for Stride")
	}

	return &strideSliceResult[T]{r, step}
}
