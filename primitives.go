package ranges

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

// BidirectionalRange is a ForwardRange that can be accessed from both directions.
type BidirectionalRange[T any] interface {
	ForwardRange[T]
	Back() T
	PopBack()
	SaveB() BidirectionalRange[T]
}

// RandomAccessRange is a Bidirectional ranges with a known length.
//
// Len() ought to be a constant-time access operation, as in O(1)
type RandomAccessRange[T any] interface {
	BidirectionalRange[T]
	Get(index int) T
	Len() int
	SaveR() RandomAccessRange[T]
}

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

// Any basic real number type
type RealNumber interface {
	Integer | Float
}

// Ordered is any type can can be compared with < and >
type Ordered interface {
	RealNumber | ~string
}
