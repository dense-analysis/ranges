package ranges

import "testing"

func TestNull(t *testing.T) {
	t.Parallel()

	null := Null[int]()

	assertEmptyB(t, null)
}

func TestOnly(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(Only[int]()), []int{})
	assertEqual(t, SliceB(Only(1)), []int{1})
	assertEqual(t, SliceB(Only(1, 2)), []int{1, 2})
	assertEqual(t, SliceB(Only(1, 2, 3)), []int{1, 2, 3})
}
