package ranges

type roundRobinResult[T any] struct {
	ranges []InputRange[T]
	index  int
}

func (rr *roundRobinResult[T]) prime() {
	start := rr.index

	for rr.ranges[rr.index].Empty() {
		rr.index++

		if rr.index >= len(rr.ranges) {
			rr.index = 0
		}

		if rr.index == start {
			panic("RoundRobin was empty")
		}
	}
}

func (rr *roundRobinResult[T]) Empty() bool {
	return AllS(rr.ranges, InputRange[T].Empty)
}

func (rr *roundRobinResult[T]) Front() T {
	rr.prime()

	return rr.ranges[rr.index].Front()
}

func (rr *roundRobinResult[T]) PopFront() {
	rr.prime()

	rr.ranges[rr.index].PopFront()

	rr.index++

	if rr.index >= len(rr.ranges) {
		rr.index = 0
	}
}

type roundRobinForwardResult[T any] struct {
	ranges []ForwardRange[T]
	index  int
}

func (rr *roundRobinForwardResult[T]) prime() {
	start := rr.index

	for rr.ranges[rr.index].Empty() {
		rr.index++

		if rr.index >= len(rr.ranges) {
			rr.index = 0
		}

		if rr.index == start {
			panic("RoundRobin was empty")
		}
	}
}

func (rr *roundRobinForwardResult[T]) Empty() bool {
	return AllS(rr.ranges, ForwardRange[T].Empty)
}

func (rr *roundRobinForwardResult[T]) Front() T {
	rr.prime()

	return rr.ranges[rr.index].Front()
}

func (rr *roundRobinForwardResult[T]) PopFront() {
	rr.prime()

	rr.ranges[rr.index].PopFront()

	rr.index++

	if rr.index >= len(rr.ranges) {
		rr.index = 0
	}
}

func (rr *roundRobinForwardResult[T]) Save() ForwardRange[T] {
	ranges := make([]ForwardRange[T], len(rr.ranges))

	for i := range rr.ranges {
		ranges[i] = rr.ranges[i].Save()
	}

	return &roundRobinForwardResult[T]{ranges, rr.index}
}

// RoundRobin yields the first elements of the ranges and cycles back around until all are consumed.
func RoundRobin[T any](ranges ...InputRange[T]) InputRange[T] {
	return &roundRobinResult[T]{ranges, 0}
}

// RoundRobinF is `RoundRobin` producing a ForwardRange
func RoundRobinF[T any](ranges ...ForwardRange[T]) ForwardRange[T] {
	return &roundRobinForwardResult[T]{ranges, 0}
}
