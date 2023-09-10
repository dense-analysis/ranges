package ranges

import "testing"

func TestNull(t *testing.T) {
	t.Parallel()

	null := Null[int]()

	assertEmpty(t, null)
	assertEqual(t, null.Len(), 0)
}

func TestOnly(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceR(Only[int]()), []int{})
	assertEqual(t, SliceR(Only(1)), []int{1})
	assertEqual(t, SliceR(Only(1, 2)), []int{1, 2})
	assertEqual(t, SliceR(Only(1, 2, 3)), []int{1, 2, 3})
}
