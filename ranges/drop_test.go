package ranges

import "testing"

func TestDrop(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice(Drop[int](Only(1, 2, 3, 4, 5), 3))

	assertEqual(t, sliceCopy, []int{4, 5})
}

func TestDropZero(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice(Drop[int](Only(1, 2, 3), 0))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestDropNegative(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice(Drop[int](Only(1, 2, 3), -10))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestDropCallFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := Drop[int](Only(1, 2, 3, 4, 5), 3)

	assertHasFront(t, dropRange, 4)
}

func TestDropCallPopFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := Drop[int](Only(1, 2, 3, 4, 5), 3)

	dropRange.PopFront()
	assertHasFront(t, dropRange, 5)
}

func TestDropF(t *testing.T) {
	t.Parallel()

	takeForwardResult := DropF(Only(1, 2, 3, 4, 5), 3)

	sliceCopy := SliceF(takeForwardResult.Save())
	sliceCopy2 := SliceF(takeForwardResult.Save())

	assertEqual(t, sliceCopy, []int{4, 5})
	assertEqual(t, sliceCopy2, []int{4, 5})
}

func TestDropFCallFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := DropF(Only(1, 2, 3, 4, 5), 3)

	assertHasFrontF(t, dropRange, 4)
}

func TestDropFCallPopFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := DropF(Only(1, 2, 3, 4, 5), 3)

	dropRange.PopFront()
	assertHasFrontF(t, dropRange, 5)
}

func TestDropFZero(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice[int](DropF(Only(1, 2, 3), 0))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestDropFNegative(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice[int](DropF(Only(1, 2, 3), -10))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}
