package ranges

import "testing"

func TestUniq(t *testing.T) {
	t.Parallel()

	result := Slice(Uniq(I(F(Only(4, 1, 2, 2, 3, 4, 4, 5))), Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqF(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqF(F(Only(4, 1, 2, 2, 3, 4, 4, 5)), Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
	assertHasSaveableFront(t, UniqF(F(Only(4, 1)), Eq[int]), 4)
}

func TestUniqS(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqS([]int{4, 1, 2, 2, 3, 4, 4, 5}, Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqComparable(t *testing.T) {
	t.Parallel()

	result := Slice(UniqComparable(I(F(Only(4, 1, 2, 2, 3, 4, 4, 5)))))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqComparableF(F(Only(4, 1, 2, 2, 3, 4, 4, 5))))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
	assertHasSaveableFront(t, UniqComparableF(F(Only(4, 1))), 4)
}

func TestUniqComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqComparableS([]int{4, 1, 2, 2, 3, 4, 4, 5}))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}
