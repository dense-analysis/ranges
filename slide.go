package ranges

type slideResult[T any] struct {
	r          ForwardRange[T]
	windowSize int
	stepSize   int
}

func (sr *slideResult[T]) Empty() bool {
	return sr.r.Empty()
}

func (sr *slideResult[T]) Front() ForwardRange[T] {
	return TakeF(sr.r.Save(), sr.windowSize)
}

func (sr *slideResult[T]) PopFront() {
	for i := 0; i < sr.stepSize && !sr.r.Empty(); i++ {
		sr.r.PopFront()
	}

	// Check if there's a remaining window of the given size.
	if !sr.r.Empty() {
		r := sr.r.Save()
		r.PopFront()

		for i := sr.stepSize + 1; i < sr.windowSize && !r.Empty(); i++ {
			r.PopFront()
		}

		if r.Empty() {
			sr.r = r
		}
	}
}

func (sr *slideResult[T]) Save() ForwardRange[ForwardRange[T]] {
	return &slideResult[T]{sr.r.Save(), sr.windowSize, sr.stepSize}
}

// Slide yields chunks starting from every element in a range.
//
// This function will panic `if windowSize < 1`.
func Slide[T any](r ForwardRange[T], windowSize int) ForwardRange[ForwardRange[T]] {
	return SlideStep(r, windowSize, 1)
}

// SlideS is `Slide` accepting a slice.
func SlideS[T any](r []T, windowSize int) ForwardRange[ForwardRange[T]] {
	return SlideStepS(r, windowSize, 1)
}

// SlideStep is `Slide` steping over elements with a given `stepSize`
func SlideStep[T any](r ForwardRange[T], windowSize int, stepSize int) ForwardRange[ForwardRange[T]] {
	if windowSize < 1 {
		panic("windowSize < 1 for Slide")
	}

	return &slideResult[T]{r, windowSize, stepSize}
}

// SlideStepS is `Slide` accepting a slice.
func SlideStepS[T any](r []T, windowSize int, stepSize int) ForwardRange[ForwardRange[T]] {
	return SlideStep(F(SliceRange(r)), windowSize, stepSize)
}
