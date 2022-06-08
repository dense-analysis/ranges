package ranges

type cacheResult[T any] struct {
	r     InputRange[T]
	value T
}

func (c *cacheResult[T]) Empty() bool {
	return c.r.Empty()
}

func (c *cacheResult[T]) Front() T {
	if c.r.Empty() {
		panic("Front() called on empty Cache result")
	}

	return c.value
}

func (c *cacheResult[T]) PopFront() {
	c.r.PopFront()

	if !c.r.Empty() {
		c.value = c.r.Front()
	}
}

type cacheForwardResult[T any] struct {
	cacheResult[T]
}

func (c *cacheForwardResult[T]) Save() ForwardRange[T] {
	return &cacheForwardResult[T]{cacheResult[T]{c.r.(ForwardRange[T]).Save(), c.value}}
}

// Cache eagerly caches the first element in a range and caches values so
// `Front()` is only called once per element for the original range.
func Cache[T any](r InputRange[T]) InputRange[T] {
	if r.Empty() {
		return Null[T]()
	}

	return &cacheResult[T]{r, r.Front()}
}

// CacheF is `Cache` for forward ranges.
//
// If the range is traversed once, `Front()` will be called once per element.
// If the range is saved, `Front()` will be called multiple times.
func CacheF[T any](r ForwardRange[T]) ForwardRange[T] {
	if r.Empty() {
		return Null[T]()
	}

	return &cacheForwardResult[T]{cacheResult[T]{r, r.Front()}}
}
