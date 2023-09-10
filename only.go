package ranges

type nullRange[T any] struct{}

func (nr *nullRange[T]) Empty() bool                  { return true }
func (nr *nullRange[T]) Front() T                     { panic("Front() access on NullRange") }
func (nr *nullRange[T]) PopFront()                    { panic("PopFront() access on NullRange") }
func (nr *nullRange[T]) Back() T                      { panic("Back() access on NullRange") }
func (nr *nullRange[T]) PopBack()                     { panic("PopBack() access on NullRange") }
func (nr *nullRange[T]) Get(index int) T              { panic("Get() access on NullRange") }
func (nr *nullRange[T]) Len() int                     { return 0 }
func (nr *nullRange[T]) Save() ForwardRange[T]        { return nr }
func (nr *nullRange[T]) SaveB() BidirectionalRange[T] { return nr }
func (nr *nullRange[T]) SaveR() RandomAccessRange[T]  { return nr }

// Returns a range that's empty
func Null[T any]() RandomAccessRange[T] { return (*nullRange[T])(nil) }

// Only returns a ForwardRange through the arguments provided.
func Only[T any](values ...T) RandomAccessRange[T] {
	return SliceRange(values)
}
