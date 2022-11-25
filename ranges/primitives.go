package ranges

// Signed is any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer integer type.
type Integer interface {
	Signed | Unsigned
}

// Float is any floating point type.
type Float interface {
	~float32 | ~float64
}

// Ordered is any type can can be compared with < and >
type Ordered interface {
	Integer | Float | ~string
}

// Any basic real number type
type RealNumber interface {
	int | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// HasLength is any type that has a length.
//
// Length ought to be a constant-time access operation.
type HasLength interface {
	Length() int
}

// OutputRange is any type that values can be output to.
type OutputRange[T any] interface {
	Put(element T) error
}

// InputRange is any type you can read sequentially.
type InputRange[T any] interface {
	Empty() bool
	Front() T
	PopFront()
}

// ForwardRange is an InputRange you can save the position of.
type ForwardRange[T any] interface {
	InputRange[T]
	Save() ForwardRange[T]
}

// I is a convenience function for passing a ForwardRange as an InputRange.
func I[T any](r ForwardRange[T]) InputRange[T] { return r }

// BidirectionalRange is a ForwardRange that can be accessed from both directions.
type BidirectionalRange[T any] interface {
	ForwardRange[T]
	SaveB() BidirectionalRange[T]
	Back() T
	PopBack()
}

// F is a convenience function for passing a BidirectionalRange as a ForwardRange
func F[T any](r BidirectionalRange[T]) ForwardRange[T] { return r }
