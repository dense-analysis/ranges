package ranges

type retroResult[T any] struct{ r BidirectionalRange[T] }

func (r *retroResult[T]) Empty() bool                  { return r.r.Empty() }
func (r *retroResult[T]) Front() T                     { return r.r.Back() }
func (r *retroResult[T]) PopFront()                    { r.r.PopBack() }
func (r *retroResult[T]) Back() T                      { return r.r.Front() }
func (r *retroResult[T]) PopBack()                     { r.r.PopFront() }
func (r *retroResult[T]) Save() ForwardRange[T]        { return &retroResult[T]{r.r.SaveB()} }
func (r *retroResult[T]) SaveB() BidirectionalRange[T] { return &retroResult[T]{r.r.SaveB()} }

// Retro iterates a BidirectionalRange in reverse.
func Retro[T any](r BidirectionalRange[T]) BidirectionalRange[T] {
	return &retroResult[T]{r}
}
