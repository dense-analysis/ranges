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

func TestFilterS(t *testing.T) {
	t.Parallel()

	r := FilterS([]int{1, 2, 3, 4}, func(element int) bool { return element%2 == 0 })
	assertEqual(t, SliceF(r), []int{2, 4})
}

func TestFilterB(t *testing.T) {
	t.Parallel()

	r := Only(1, 2, 3, 4)
	fr := FilterB(r, func(element int) bool { return element%2 == 0 })
	fr2 := fr.SaveB()

	fr.PopFront()
	fr.PopFront()

	assertEmpty(t, fr)
	assertHasFront(t, fr2, 2)
	assertHasBack(t, fr2, 4)

	fr2.PopBack()

	assertHasFront(t, fr2, 2)
	assertHasBack(t, fr2, 2)

	fr2.PopFront()

	assertEmpty(t, fr2)
}

func TestFilterSB(t *testing.T) {
	t.Parallel()

	r := FilterSB([]int{1, 2, 3, 4}, func(element int) bool { return element%2 == 0 })
	assertEqual(t, SliceB(Retro(r)), []int{4, 2})
}
