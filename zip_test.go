package ranges

import "testing"

func TestZip2(t *testing.T) {
	output := Slice(
		Map(
			Zip2(
				Only(1, 2, 3, 4),
				Only(4, 5, 6),
			),
			func(e Pair[int, int]) []int {
				return []int{e.A, e.B}
			},
		),
	)

	assertEqual(t, output, [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	})
}

func TestZip3(t *testing.T) {
	output := Slice(
		Map(
			Zip3(
				Only(1, 2, 3, 4),
				Only(4.0, 5.0, 6.0, 7.0),
				Only(9, 8, 7),
			),
			func(e Triplet[int, float64, int]) []int {
				return []int{e.A, int(e.B), e.C}
			},
		),
	)

	assertEqual(t, output, [][]int{
		{1, 4, 9},
		{2, 5, 8},
		{3, 6, 7},
	})
}

func TestZip4F(t *testing.T) {
	r := MapF(
		Zip4F(
			Only(1, 2, 3, 4),
			Only(4.0, 5.0, 6.0, 7.0),
			Only(9, 8, 7),
			Only(-1.0, -2.0, -3.0),
		),
		func(e Quartet[int, float64, int, float64]) []int {
			return []int{e.A, int(e.B), e.C, int(e.D)}
		},
	)

	assertHasFront(t, r, []int{1, 4, 9, -1})

	r.PopFront()

	assertHasFront(t, r, []int{2, 5, 8, -2})

	rSave := r.Save()
	r.PopFront()

	assertHasFront(t, rSave, []int{2, 5, 8, -2})
	assertHasFront(t, r, []int{3, 6, 7, -3})

	r.PopFront()

	assertEmpty(t, r)
}
