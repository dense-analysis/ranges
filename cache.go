package ranges

type cacheResult[T any] struct {
	r        InputRange[T]
	value    T
	isPrimed bool
}

func (c *cacheResult[T]) Empty() bool {
	return c.r.Empty()
}

func (c *cacheResult[T]) Front() T {
	if !c.isPrimed {
		c.value = c.r.Front()
		c.isPrimed = true
	}

	return c.value
}

func (c *cacheResult[T]) PopFront() {
	c.r.PopFront()
	c.isPrimed = false
}

type cacheForwardResult[T any] struct {
	cacheResult[T]
}

func (c *cacheForwardResult[T]) Save() ForwardRange[T] {
	return &cacheForwardResult[T]{cacheResult[T]{c.r.(ForwardRange[T]).Save(), c.value, c.isPrimed}}
}

type cacheBidirectionalResult[T any] struct {
	cacheResult[T]
	backValue    T
	backIsPrimed bool
}

func (c *cacheBidirectionalResult[T]) Back() T {
	if !c.backIsPrimed {
		c.backValue = c.r.(BidirectionalRange[T]).Back()
		c.backIsPrimed = true
	}

	return c.backValue
}

func (c *cacheBidirectionalResult[T]) PopBack() {
	c.r.(BidirectionalRange[T]).PopBack()
	c.backIsPrimed = false
}

func (c *cacheBidirectionalResult[T]) Save() ForwardRange[T] {
	return c.SaveB()
}

func (c *cacheBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &cacheBidirectionalResult[T]{
		cacheResult[T]{c.r.(BidirectionalRange[T]).SaveB(), c.value, c.isPrimed},
		c.backValue,
		c.backIsPrimed,
	}
}

// Cache eagerly caches the first element in a range and caches values so
// `Front()` is only called once per element for the original range.
func Cache[T any](r InputRange[T]) InputRange[T] {
	return &cacheResult[T]{r, *new(T), false}
}

// CacheF is `Cache` for forward ranges.
//
// If the range is traversed once, `Front()` will be called once per element.
// If the range is saved, `Front()` will be called multiple times.
func CacheF[T any](r ForwardRange[T]) ForwardRange[T] {
	return &cacheForwardResult[T]{cacheResult[T]{r, *new(T), false}}
}

// CacheB is `CacheF` for bidrectional ranges.
//
// Traversing in both directions could cause the same expression to be evaluated twice.
// If the range is saved, `Front()` or `Back()` will be called multiple times.
func CacheB[T any](r BidirectionalRange[T]) BidirectionalRange[T] {
	return &cacheBidirectionalResult[T]{cacheResult[T]{r, *new(T), false}, *new(T), false}
}
