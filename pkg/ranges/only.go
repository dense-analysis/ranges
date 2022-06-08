package ranges

type nullRange[T any] struct { }

func (nr *nullRange[T]) Empty() bool { return true }
func (nr *nullRange[T]) Front() T { panic("Front() access on NullRange") }
func (nr *nullRange[T]) PopFront() { panic("PopFront() access on NullRange") }
func (nr *nullRange[T]) Save() ForwardRange[T] { return nr }

// Returns a ForwardRange that's empty
func Null[T any]() ForwardRange[T] { return (*nullRange[T])(nil) }

// Only returns a ForwardRange through the arguments provided.
func Only[T any](values ...T) ForwardRange[T] {
	return SliceRange(values)
}
