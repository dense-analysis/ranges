package ranges

type zipNResult struct {
	ranges []InputRange[any]
}

func (z *zipNResult) Empty() bool {
	return AnyS(z.ranges, InputRange[any].Empty)
}

func (z *zipNResult) Front() []any {
	return SliceB(MapS(z.ranges, InputRange[any].Front))
}

func (z *zipNResult) PopFront() {
	EachS(z.ranges, InputRange[any].PopFront)
}

// ZipN produces items from any n ranges in parallel.
func ZipN(ranges ...InputRange[any]) InputRange[[]any] {
	return &zipNResult{ranges}
}
