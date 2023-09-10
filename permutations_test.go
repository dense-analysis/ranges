package ranges

import "testing"

func TestPermutations(t *testing.T) {
	t.Parallel()

	assertEmpty(t, Permutations([]int{}))

	assertEqual(
		t,
		SliceF(Permutations([]int{1, 2, 3})),
		[][]int{
			{1, 2, 3}, {2, 1, 3}, {3, 1, 2},
			{1, 3, 2}, {2, 3, 1}, {3, 2, 1},
		},
	)

	assertEqual(
		t,
		SliceF(
			MapF(
				Permutations(SliceF(Runes("abcd"))),
				StringS,
			),
		),
		[]string{
			"abcd", "bacd", "cabd", "acbd",
			"bcad", "cbad", "dbac", "bdac",
			"adbc", "dabc", "badc", "abdc",
			"acdb", "cadb", "dacb", "adcb",
			"cdab", "dcab", "dcba", "cdba",
			"bdca", "dbca", "cbda", "bcda",
		},
	)
}

func TestPermutationsSave(t *testing.T) {
	t.Parallel()

	r := Permutations([]int{1, 2, 3})

	assertHasFront(t, r, []int{1, 2, 3})
	assertHasFront(t, r, []int{1, 2, 3})

	r.PopFront()

	assertHasFront(t, r, []int{2, 1, 3})

	rSave := r.Save()

	r.PopFront()

	assertHasFront(t, r, []int{3, 1, 2})
	assertHasFront(t, rSave, []int{2, 1, 3})
}
