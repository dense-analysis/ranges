package ranges

import "testing"

func TestStripLeftComparable(t *testing.T) {
	t.Parallel()

	result := Slice(StripLeftComparable(I(F(Only(5, 5, 3, 5, 2))), 5))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparable(I(F(Only(5, 5, 3, 5, 2))), 5)

	assertEqual(t, result2.Front(), 3)

	result3 := StripLeftComparable(I(F(Only(5, 5, 3, 5, 2))), 5)

	result3.PopFront()
	assertHasFront(t, result3, 5)

	result3.PopFront()
	assertHasFront(t, result3, 2)
}

func TestStripLeftComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(StripLeftComparableF(F(Only(5, 5, 3, 5, 2)), 5))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparableF(F(Only(5, 5, 3, 5, 2)), 5)

	assertHasSaveableFront(t, result2, 3)
}

func TestStripLeftComparableB(t *testing.T) {
	t.Parallel()

	r := StripLeftComparableB(Only(5, 5, 3, 5, 2), 5)

	assertHasSaveableBack(t, r, 2)
	assertEqual(t, SliceB(Retro(r)), []int{5, 3})
}

func TestStripLeftComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(StripLeftComparableS([]int{5, 5, 3, 5, 2}, 5))

	assertEqual(t, result, []int{3, 5, 2})
}

func TestStripLeft(t *testing.T) {
	t.Parallel()

	result := Slice(StripLeft(I(F(Only(5, 5, 3, 5, 2))), func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftComparable(I(F(Only(5, 5, 3, 5, 2))), 5)

	result2.PopFront()
	assertHasFront(t, result2, 5)

	result2.PopFront()
	assertHasFront(t, result2, 2)
}

func TestStripLeftF(t *testing.T) {
	t.Parallel()

	result := SliceF(StripLeftF(F(Only(5, 5, 3, 5, 2)), func(a int) bool { return a == 5 }))

	assertEqual(t, result, []int{3, 5, 2})

	result2 := StripLeftF(F(Only(5, 5, 3, 5, 2)), func(a int) bool { return a == 5 })

	assertHasSaveableFront(t, result2, 3)
}

func TestStripLeftB(t *testing.T) {
	t.Parallel()

	r := StripLeftB(Only(5, 5, 3, 5, 2), func(a int) bool { return a == 5 })

	assertHasSaveableBack(t, r, 2)
	assertEqual(t, SliceB(Retro(r)), []int{5, 3})
}
