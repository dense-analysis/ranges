package ranges

import "testing"

func TestCopy(t *testing.T) {
	t.Parallel()

	output := make([]int, 0, 3)
	input := []int{4, 5, 6}

	Copy[int](SliceRange(input), SliceSink(&output))

	assertEqual(t, output, input)
}

func TestFill(t *testing.T) {
	t.Parallel()

	output := make([]int, 5)

	Fill[int](SlicePtrRange(output), 3)

	assertEqual(t, output, []int{3, 3, 3, 3, 3})
}

func TestFillPattern(t *testing.T) {
	t.Parallel()

	output := make([]int, 5)

	FillPattern[int](SlicePtrRange(output), Only(1, 2, 3))

	assertEqual(t, output, []int{1, 2, 3, 1, 2})
}
