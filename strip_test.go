package ranges

import "testing"

func TestStripLeftComparable(t *testing.T) {
	t.Parallel()

	result := Slice(StripLeftComparable(Only(5, 5, 3, 5, 2), 5))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparable(Only(5, 5, 3, 5, 2), 5)

	assertEqual(t, result2.Front(), 3)

	result3 := StripLeftComparable(Only(5, 5, 3, 5, 2), 5)

	result3.PopFront()
	assertHasFront(t, result3, 5)

	result3.PopFront()
	assertHasFront(t, result3, 2)
}

func TestStripLeftComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(StripLeftComparableF(Only(5, 5, 3, 5, 2), 5))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparableF(Only(5, 5, 3, 5, 2), 5)

	assertHasSaveableFront(t, result2, 3)
}

func TestStripLeftComparableB(t *testing.T) {
	t.Parallel()

	r := StripLeftComparableB(Only(5, 5, 3, 5, 2), 5)

	assertHasSaveableBack(t, r, 2)
	assertEqual(t, SliceB(Retro(r)), []int{5, 3})
}

func TestStripLeftComparableR(t *testing.T) {
	t.Parallel()

	r := StripLeftComparableR(Only(5, 5, 3, 5, 2), 5)

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Check that Len() and Get() prime correctly.
	assertEqual(t, r2.Len(), 3)
	assertEqual(t, r3.Get(1), 5)

	// Popping the original shouldn't change a saved one.
	r.PopFront()
	r.PopBack()
	assertHasFrontR(t, r2, 3)
	assertHasBackR(t, r2, 2)
}

func TestStripLeftComparableS(t *testing.T) {
	t.Parallel()

	result := SliceR(StripLeftComparableS([]int{5, 5, 3, 5, 2}, 5))

	assertEqual(t, result, []int{3, 5, 2})
}

func TestStripLeft(t *testing.T) {
	t.Parallel()

	result := Slice(StripLeft(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparable(Only(5, 5, 3, 5, 2), 5)

	result2.PopFront()
	assertHasFront(t, result2, 5)

	result2.PopFront()
	assertHasFront(t, result2, 2)
}

func TestStripLeftF(t *testing.T) {
	t.Parallel()

	result := SliceF(StripLeftF(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftF(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 })

	assertHasSaveableFront(t, result2, 3)
}

func TestStripLeftB(t *testing.T) {
	t.Parallel()

	r := StripLeftB(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 })

	assertHasSaveableBack(t, r, 2)
	assertEqual(t, SliceB(Retro(r)), []int{5, 3})
}

func TestStripLeftR(t *testing.T) {
	t.Parallel()

	r := StripLeftR(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 })

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Check that Len() and Get() prime correctly.
	assertEqual(t, r2.Len(), 3)
	assertEqual(t, r3.Get(1), 5)

	// Popping the original shouldn't change a saved one.
	r.PopFront()
	r.PopBack()
	assertHasFrontR(t, r2, 3)
	assertHasBackR(t, r2, 2)
}

func TestStripRightComparable(t *testing.T) {
	t.Parallel()

	result := SliceB(StripRightComparable(Only(3, 5, 2, 5, 5), 5))

	assertEqual(t, result, []int{3, 5, 2})

	r := StripRightComparable(Only(0, 1, 2, 3, 0, 0), 0)

	assertHasSaveableBack(t, r, 3)
	assertHasSaveableFrontB(t, r, 0)
}

func TestStripRightComparableR(t *testing.T) {
	t.Parallel()

	r := StripRightComparableR(Only(3, 5, 2, 5, 5), 5)

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Check that Len() and Get() prime correctly.
	assertEqual(t, r2.Len(), 3)
	assertEqual(t, r3.Get(1), 5)

	// Popping the original shouldn't change a saved one.
	r.PopFront()
	r.PopBack()
	assertHasFrontR(t, r2, 3)
	assertHasBackR(t, r2, 2)
}

func TestStripRightComparableS(t *testing.T) {
	t.Parallel()

	result := SliceR(StripRightComparableS([]int{3, 5, 2, 5, 5}, 5))

	assertEqual(t, result, []int{3, 5, 2})

	r := StripRightComparableS([]int{0, 1, 2, 3, 0, 0}, 0)

	assertHasSaveableBackR(t, r, 3)
	assertHasSaveableFrontR(t, r, 0)
}

func TestStripRight(t *testing.T) {
	t.Parallel()

	result := SliceB(StripRight(Only(3, 5, 2, 5, 5), func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	r := StripRight(Only(0, 1, 2, 3, 0, 0), func(a int) bool { return a == 0 })

	assertHasSaveableBack(t, r, 3)
	assertHasSaveableFrontB(t, r, 0)
}

func TestStripRightR(t *testing.T) {
	t.Parallel()

	r := StripRightR(Only(3, 5, 2, 5, 5), func(a int) bool { return a == 5 })

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Check that Len() and Get() prime correctly.
	assertEqual(t, r2.Len(), 3)
	assertEqual(t, r3.Get(1), 5)

	// Popping the original shouldn't change a saved one.
	r.PopFront()
	r.PopBack()
	assertHasFrontR(t, r2, 3)
	assertHasBackR(t, r2, 2)
}

func TestStripRightS(t *testing.T) {
	t.Parallel()

	result := SliceR(StripRightS([]int{3, 5, 2, 5, 5}, func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	r := StripRightS([]int{0, 1, 2, 3, 0, 0}, func(a int) bool { return a == 0 })

	assertHasSaveableBackR(t, r, 3)
	assertHasSaveableFrontR(t, r, 0)
}

func TestStripComparable(t *testing.T) {
	t.Parallel()

	r := StripComparable(Only(0, 0, 0, 1, 0, 2, 0, 0), 0)

	assertHasSaveableBack(t, r, 2)
	assertHasSaveableFrontB(t, r, 1)
	assertHasFrontB(t, r, 0)
}

func TestStripComparableR(t *testing.T) {
	t.Parallel()

	r := StripComparableR(Only(0, 0, 0, 1, 0, 2, 0, 0), 0)

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Test extra methods are primed and save correctly.
	assertEqual(t, r2.Get(2), 2)
	assertEqual(t, r3.Len(), 3)

	// Test that saving don't change the original range.
	r2.PopFront()
	assertHasFrontR(t, r, 1)
}

func TestStripComparableS(t *testing.T) {
	t.Parallel()

	r := StripComparableS([]int{0, 0, 0, 1, 0, 2, 0, 0}, 0)

	assertHasSaveableBackR(t, r, 2)
	assertHasSaveableFrontR(t, r, 1)
	assertHasFrontR(t, r, 0)
}

func TestStrip(t *testing.T) {
	t.Parallel()

	r := Strip(Only(0, 0, 0, 1, 0, 2, 0, 0), func(a int) bool { return a == 0 })

	assertHasSaveableBack(t, r, 2)
	assertHasSaveableFrontB(t, r, 1)
	assertHasFrontB(t, r, 0)
}

func TestStripR(t *testing.T) {
	t.Parallel()

	r := StripR(Only(0, 0, 0, 1, 0, 2, 0, 0), func(a int) bool { return a == 0 })

	r2 := r.SaveR()
	r3 := r.SaveR()

	// Test extra methods are primed and save correctly.
	assertEqual(t, r2.Get(2), 2)
	assertEqual(t, r3.Len(), 3)

	// Test that saving don't change the original range.
	r2.PopFront()
	assertHasFrontR(t, r, 1)
}

func TestStripS(t *testing.T) {
	t.Parallel()

	r := StripS([]int{0, 0, 0, 1, 0, 2, 0, 0}, func(a int) bool { return a == 0 })

	assertHasSaveableBackR(t, r, 2)
	assertHasSaveableFrontR(t, r, 1)
	assertHasFrontR(t, r, 0)
}
