package ranges

type zipNResult struct {
	ranges []InputRange[any]
}

func (z *zipNResult) Empty() bool {
	return AnyS(z.ranges, InputRange[any].Empty)
}

func (z *zipNResult) Front() []any {
	return SliceB(B(MapS(z.ranges, InputRange[any].Front)))
}

func (z *zipNResult) PopFront() {
	EachS(z.ranges, InputRange[any].PopFront)
}

type zipNFResult struct {
	ranges []ForwardRange[any]
}

func (z *zipNFResult) Empty() bool {
	return AnyS(z.ranges, ForwardRange[any].Empty)
}

func (z *zipNFResult) Front() []any {
	return SliceB(B(MapS(z.ranges, ForwardRange[any].Front)))
}

func (z *zipNFResult) PopFront() {
	EachS(z.ranges, ForwardRange[any].PopFront)
}

func (z *zipNFResult) Save() ForwardRange[[]any] {
	return &zipNFResult{SliceB(B(MapS(z.ranges, ForwardRange[any].Save)))}
}

// ZipN produces items from any n ranges in parallel.
func ZipN(ranges ...InputRange[any]) InputRange[[]any] {
	return &zipNResult{ranges}
}

// ZipNF is `ZipN` that allows the range to be saved.
func ZipNF(ranges ...ForwardRange[any]) ForwardRange[[]any] {
	return &zipNFResult{ranges}
}
