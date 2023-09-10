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

	dropRange := DropF(F(B(Only(1, 2, 3, 4, 5))), 3)

	sliceCopy := SliceF(dropRange.Save())
	sliceCopy2 := SliceF(dropRange.Save())

	assertEqual(t, sliceCopy, []int{4, 5})
	assertEqual(t, sliceCopy2, []int{4, 5})
}

func TestDropFCallFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := DropF(F(B(Only(1, 2, 3, 4, 5))), 3)

	assertHasFrontF(t, dropRange, 4)
}

func TestDropFCallPopFrontFirst(t *testing.T) {
	t.Parallel()

	dropRange := DropF(F(B(Only(1, 2, 3, 4, 5))), 3)

	dropRange.PopFront()
	assertHasFrontF(t, dropRange, 5)
}

func TestDropFZero(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice[int](DropF(F(B(Only(1, 2, 3))), 0))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestDropFNegative(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice[int](DropF(F(B(Only(1, 2, 3))), -10))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestDropB(t *testing.T) {
	t.Parallel()

	dropRange := DropB(B(Only(1, 2, 3, 4, 5)), 3)

	sliceCopy := SliceB(dropRange.SaveB())
	sliceCopy2 := SliceB(dropRange.SaveB())

	assertEqual(t, sliceCopy, []int{4, 5})
	assertEqual(t, sliceCopy2, []int{4, 5})

	assertHasBack(t, dropRange, 5)
	dropRange.PopBack()
	assertHasBack(t, dropRange, 4)
	dropRange.PopBack()
	assertEmptyB(t, dropRange)
}

func TestDropR(t *testing.T) {
	t.Parallel()

	r := DropR(Only(1, 2, 3, 4, 5), 3)
	r2 := r.SaveR()
	r.PopFront()

	assertEqual(t, r2.Get(0), 4)
	assertEqual(t, SliceR(r), []int{5})
	assertEqual(t, SliceR(r2), []int{4, 5})
}

// Ensure we prime the drop range if we .Get() before other operations.
func TestDropRGetPrimed(t *testing.T) {
	t.Parallel()

	r := DropR(Only(1, 2, 3, 4, 5), 3)

	assertEqual(t, r.Get(0), 4)
	assertEqual(t, r.Get(1), 5)
}

// Ensure we prime the drop range if we check length before other operations.
func TestDropRLen(t *testing.T) {
	t.Parallel()

	r := DropR(Only(1, 2, 3, 4, 5), 3)

	assertEqual(t, r.Len(), 2)
}
