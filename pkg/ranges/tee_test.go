package ranges

import "testing"

func TestTee(t *testing.T) {
	t.Parallel()

	output := make([]int, 0)

	tee := Tee[int](Only(1, 2), SliceSink(&output))

	assertNotEmpty(t, tee)
	assertEqual(t, tee.Front(), 1)

	assertEqual(t, output, []int{})
	tee.PopFront()
	assertEqual(t, output, []int{1})

	assertNotEmpty(t, tee)
	assertEqual(t, tee.Front(), 2)

	assertEqual(t, output, []int{1})
	tee.PopFront()
	assertEqual(t, output, []int{1, 2})
}

func TestTeeS(t *testing.T) {
	t.Parallel()

	output := make([]int, 0)

	tee := TeeS([]int{1, 2}, SliceSink(&output))

	assertNotEmpty(t, tee)
	assertEqual(t, tee.Front(), 1)

	assertEqual(t, output, []int{})
	tee.PopFront()
	assertEqual(t, output, []int{1})

	assertNotEmpty(t, tee)
	assertEqual(t, tee.Front(), 2)

	assertEqual(t, output, []int{1})
	tee.PopFront()
	assertEqual(t, output, []int{1, 2})
}
