package ranges

type nullSink[T any] struct{}

func (nr *nullSink[T]) Put(element T) error { return nil }

// NullSink returns a sink that discards output.
func NullSink[T any]() OutputRange[T] {
	return (*nullSink[T])(nil)
}

type sliceSink[T any] struct {
	slice *[]T
}

func (s sliceSink[T]) Put(element T) error {
	*s.slice = append(*s.slice, element)

	return nil
}

// SliceSink creates a OutputRange that appends elements to the slice.
func SliceSink[T any](slice *[]T) OutputRange[T] {
	if slice == nil {
		panic("SliceSink called with a nil slice")
	}

	return sliceSink[T]{slice}
}

type assignSink[T any] struct {
	r InputRange[*T]
}

func (as assignSink[T]) Put(element T) error {
	*as.r.Front() = element
	as.r.PopFront()

	return nil
}

// AssignSink assigns values by derefencing pointers from a range.
func AssignSink[T any](r InputRange[*T]) OutputRange[T] {
	return assignSink[T]{r}
}
