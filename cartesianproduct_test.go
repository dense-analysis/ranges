package ranges

import "testing"

func TestCartesianProduct2(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(CartesianProduct2(F(Iota(2)), F(Iota(3)))),
		[]Pair[int, int]{
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 0},
			{1, 1},
			{1, 2},
		},
	)
	assertEqual(
		t,
		SliceF(CartesianProduct2(F(Iota(3)), F(Iota(2)))),
		[]Pair[int, int]{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
			{2, 0},
			{2, 1},
		},
	)
}

func TestCartesianProduct3(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(CartesianProduct3(F(B(Runes("abc"))), F(Iota(2)), F(B(Runes("xyz"))))),
		[]Triplet[rune, int, rune]{
			{'a', 0, 'x'}, {'a', 0, 'y'}, {'a', 0, 'z'}, {'a', 1, 'x'}, {'a', 1, 'y'}, {'a', 1, 'z'}, {'b', 0, 'x'}, {'b', 0, 'y'}, {'b', 0, 'z'}, {'b', 1, 'x'}, {'b', 1, 'y'}, {'b', 1, 'z'}, {'c', 0, 'x'}, {'c', 0, 'y'}, {'c', 0, 'z'}, {'c', 1, 'x'}, {'c', 1, 'y'}, {'c', 1, 'z'},
		},
	)
}

func TestCartesianProduct4(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(CartesianProduct4(F(B(Runes("abc"))), F(Iota(2)), F(B(Runes("xyz"))), F(B(Runes("42"))))),
		[]Quartet[rune, int, rune, rune]{
			{'a', 0, 'x', '4'}, {'a', 0, 'x', '2'}, {'a', 0, 'y', '4'}, {'a', 0, 'y', '2'}, {'a', 0, 'z', '4'},
			{'a', 0, 'z', '2'}, {'a', 1, 'x', '4'}, {'a', 1, 'x', '2'}, {'a', 1, 'y', '4'}, {'a', 1, 'y', '2'},
			{'a', 1, 'z', '4'}, {'a', 1, 'z', '2'}, {'b', 0, 'x', '4'}, {'b', 0, 'x', '2'}, {'b', 0, 'y', '4'},
			{'b', 0, 'y', '2'}, {'b', 0, 'z', '4'}, {'b', 0, 'z', '2'}, {'b', 1, 'x', '4'}, {'b', 1, 'x', '2'},
			{'b', 1, 'y', '4'}, {'b', 1, 'y', '2'}, {'b', 1, 'z', '4'}, {'b', 1, 'z', '2'}, {'c', 0, 'x', '4'},
			{'c', 0, 'x', '2'}, {'c', 0, 'y', '4'}, {'c', 0, 'y', '2'}, {'c', 0, 'z', '4'}, {'c', 0, 'z', '2'},
			{'c', 1, 'x', '4'}, {'c', 1, 'x', '2'}, {'c', 1, 'y', '4'}, {'c', 1, 'y', '2'}, {'c', 1, 'z', '4'},
			{'c', 1, 'z', '2'},
		},
	)
}

func TestCartesianProduct10(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(CartesianProduct10(
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
			F(Iota(2)),
		)),
		SliceF(
			MapF(
				F(Iota(1024)),
				func(num int) Decade[int, int, int, int, int, int, int, int, int, int] {
					return MakeDecade(
						(num&512)>>9,
						(num&256)>>8,
						(num&128)>>7,
						(num&64)>>6,
						(num&32)>>5,
						(num&16)>>4,
						(num&8)>>3,
						(num&4)>>2,
						(num&2)>>1,
						num&1,
					)
				},
			),
		),
	)
}
