package ranges

type repeatResult[T any] struct {
	value T
}

func (repeatResult[T]) Empty() bool                    { return false }
func (r repeatResult[T]) Front() T                     { return r.value }
func (repeatResult[T]) PopFront()                      {}
func (r repeatResult[T]) Back() T                      { return r.value }
func (repeatResult[T]) PopBack()                       {}
func (r repeatResult[T]) Save() ForwardRange[T]        { return r }
func (r repeatResult[T]) SaveB() BidirectionalRange[T] { return r }

// Repeat repeats a value infinitely
func Repeat[T any](value T) BidirectionalRange[T] {
	return repeatResult[T]{value}
}

type generateResult[T any] struct{ cb func() T }

func (generateResult[T]) Empty() bool                    { return false }
func (r generateResult[T]) Front() T                     { return r.cb() }
func (generateResult[T]) PopFront()                      {}
func (r generateResult[T]) Back() T                      { return r.cb() }
func (generateResult[T]) PopBack()                       {}
func (r generateResult[T]) Save() ForwardRange[T]        { return r }
func (r generateResult[T]) SaveB() BidirectionalRange[T] { return r }

// Generate genereates a value infinitely by calling a function
func Generate[T any](cb func() T) BidirectionalRange[T] {
	return generateResult[T]{cb}
}
