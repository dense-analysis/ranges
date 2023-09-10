package ranges

import (
	"testing"
)

func TestTakeZero(t *testing.T) {
	t.Parallel()

	assertEmpty(t, Take[int](Only(1), 0))
}

func TestTakeNegative(t *testing.T) {
	t.Parallel()

	assertEmpty(t, Take[int](Only(1), -10))
}

func TestTake(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice(Take[int](Only(1, 2, 3, 4, 5), 3))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestTakeF(t *testing.T) {
	t.Parallel()

	takeForwardResult := TakeF(Only(1, 2, 3, 4, 5), 3)

	sliceCopy := SliceF(takeForwardResult.Save())
	sliceCopy2 := SliceF(takeForwardResult.Save())

	assertEqual(t, sliceCopy, []int{1, 2, 3})
	assertEqual(t, sliceCopy2, []int{1, 2, 3})
}

func TestTakeFZero(t *testing.T) {
	t.Parallel()

	assertEmptyF(t, TakeF(Only(1), 0))
}

func TestTakeFNegative(t *testing.T) {
	t.Parallel()

	assertEmptyF(t, TakeF(Only(1), -10))
}
