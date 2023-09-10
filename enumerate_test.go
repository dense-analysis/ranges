package ranges

import "testing"

func TestEnumerate(t *testing.T) {
	t.Parallel()

	keys := Slice(Map(
		Enumerate[int](Only(7, 9, 1)),
		func(x Pair[int, int]) int { return x.A },
	))

	assertEqual(t, keys, []int{0, 1, 2})

	values := Slice(Map(
		Enumerate[int](Only(7, 9, 1)),
		func(x Pair[int, int]) int { return x.B },
	))

	assertEqual(t, values, []int{7, 9, 1})
}

func TestEnumerateN(t *testing.T) {
	t.Parallel()

	keys := Slice(Map(
		EnumerateN[int](Only(7, 9, 1), 3),
		func(x Pair[int, int]) int { return x.A },
	))

	assertEqual(t, keys, []int{3, 4, 5})
}

func TestEnumerateF(t *testing.T) {
	t.Parallel()

	enumerateRange := EnumerateF(F(B(Only(7, 9, 1))))

	keys := SliceF(MapF(
		enumerateRange.Save(),
		func(x Pair[int, int]) int { return x.A },
	))

	assertEqual(t, keys, []int{0, 1, 2})

	values := SliceF(MapF(
		enumerateRange,
		func(x Pair[int, int]) int { return x.B },
	))

	assertEqual(t, values, []int{7, 9, 1})
}

func TestEnumerateNF(t *testing.T) {
	t.Parallel()

	keys := SliceF(MapF(
		EnumerateNF(F(B(Only(7, 9, 1))), 3),
		func(x Pair[int, int]) int { return x.A },
	))

	assertEqual(t, keys, []int{3, 4, 5})
}
