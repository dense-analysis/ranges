package ranges

type cycleResult[T any] struct {
	start   ForwardRange[T]
	current ForwardRange[T]
}

func (cr *cycleResult[T]) Empty() bool {
	return cr.start.Empty()
}

func (cr *cycleResult[T]) Front() T {
	return cr.current.Front()
}

func (cr *cycleResult[T]) PopFront() {
	cr.current.PopFront()

	if cr.current.Empty() {
		cr.current = cr.start.Save()
	}
}

func (cr *cycleResult[T]) Save() ForwardRange[T] {
	return &cycleResult[T]{cr.start.Save(), cr.current.Save()}
}

// Cycle repeats a ForwardRange infinitely.
func Cycle[T any](r ForwardRange[T]) ForwardRange[T] {
	return &cycleResult[T]{r.Save(), r}
}
