package ranges

import "testing"

func TestFilter(t *testing.T) {
	t.Parallel()

	r := Only(1, 2, 3, 4)
	sliceCopy := Slice(Filter[int](r, func(element int) bool { return element%2 == 0 }))

	assertEqual(t, sliceCopy, []int{2, 4})
}

func TestFilterIsLazy(t *testing.T) {
	t.Parallel()

	// This will panic if it's not lazy.
	Filter[int](nil, nil)
}

func TestFilterF(t *testing.T) {
	t.Parallel()

	r := F(Only(1, 2, 3, 4))
	fr := FilterF(r, func(element int) bool { return element%2 == 0 })
	fr2 := fr.Save()

	fr.PopFront()
	fr.PopFront()

	assertHasFrontF(t, fr2, 2)

	fr2.PopFront()

	assertHasFrontF(t, fr2, 4)
}
