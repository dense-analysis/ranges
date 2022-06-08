package ranges

// Copy copies all values from an InputRange into an OutputRange
func Copy[T any](input InputRange[T], output OutputRange[T]) error {
	for !input.Empty() {
		if err := output.Put(input.Front()); err != nil {
			return err
		}

		input.PopFront()
	}

	return nil
}

// Fill assigns the given value to the locations of all addresses in a range.
func Fill[T any](r InputRange[*T], value T) {
	for !r.Empty() {
		*r.Front() = value
		r.PopFront()
	}
}

// FillPattern assigns a pattern of values to the locations of all addresses in a range.
func FillPattern[T any](r InputRange[*T], pattern ForwardRange[T]) {
	cycle := Cycle(pattern)

	for !r.Empty() {
		*r.Front() = cycle.Front()
		r.PopFront()
		cycle.PopFront()
	}
}
