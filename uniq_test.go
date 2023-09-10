package ranges

import "testing"

func TestUniq(t *testing.T) {
	t.Parallel()

	result := Slice(Uniq(I(F(B(Only(4, 1, 2, 2, 3, 4, 4, 5)))), Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqF(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqF(F(B(Only(4, 1, 2, 2, 3, 4, 4, 5))), Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
	assertHasSaveableFront(t, UniqF(F(B(Only(4, 1))), Eq[int]), 4)
}

func TestUniqS(t *testing.T) {
	t.Parallel()

	result := SliceB(UniqS([]int{4, 1, 2, 2, 3, 4, 4, 5}, Eq[int]))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqComparable(t *testing.T) {
	t.Parallel()

	result := Slice(UniqComparable(I(F(B(Only(4, 1, 2, 2, 3, 4, 4, 5))))))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

func TestUniqComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(UniqComparableF(F(B(Only(4, 1, 2, 2, 3, 4, 4, 5)))))
	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})

	assertHasSaveableFront(t, UniqComparableF(F(B(Only(4, 1)))), 4)
}

func TestUniqComparableB(t *testing.T) {
	t.Parallel()

	result := SliceB(UniqComparableB(B(Only(4, 1, 2, 2, 3, 4, 4, 5))))
	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
	result2 := SliceB(Retro(UniqComparableB(B(Only(4, 1, 2, 2, 3, 4, 4, 5)))))
	assertEqual(t, result2, []int{5, 4, 3, 2, 1, 4})

	assertHasSaveableFrontB(t, UniqComparableB(B(Only(4, 1))), 4)
	assertHasSaveableBack(t, UniqComparableB(B(Only(4, 1))), 1)
}

func TestUniqComparableS(t *testing.T) {
	t.Parallel()

	result := SliceB(UniqComparableS([]int{4, 1, 2, 2, 3, 4, 4, 5}))

	assertEqual(t, result, []int{4, 1, 2, 3, 4, 5})
}

// Check that we use cb(front, Front()) and cb(Back(), back) for iterations.
func TestUniqBConsistentOrdering(t *testing.T) {
	t.Parallel()

	result := SliceB(UniqB(B(Only(1, 2, 3, 4, 5, 6, 7, 8)), Le[int]))
	assertEqual(t, result, []int{1})

	result2 := SliceB(Retro(UniqB(B(Only(1, 2, 3, 4, 5, 6, 7, 8)), Le[int])))
	assertEqual(t, result2, []int{8})
}
