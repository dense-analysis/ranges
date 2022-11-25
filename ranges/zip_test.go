package ranges

import "testing"

func TestZip2(t *testing.T) {
	output := Slice(
		Map(
			Zip2(
				I(F(Only(1, 2, 3, 4))),
				I(F(Only(4, 5, 6))),
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
				I(F(Only(1, 2, 3, 4))),
				I(F(Only(4.0, 5.0, 6.0, 7.0))),
				I(F(Only(9, 8, 7))),
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
