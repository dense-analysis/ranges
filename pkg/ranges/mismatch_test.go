package ranges

import "testing"

func TestMismatch2(t *testing.T) {
	t.Parallel()

	r := Mismatch2(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Pair[int, int]{{42, 69}, {1, 1}})
}

func TestMismatch3(t *testing.T) {
	t.Parallel()

	r := Mismatch3(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Triplet[int, int, int]{{42, 42, 69}, {1, 1, 1}})
}

func TestMismatch4(t *testing.T) {
	t.Parallel()

	r := Mismatch4(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Quartet[int, int, int, int]{{42, 42, 42, 69}, {1, 1, 1, 1}})
}

func TestMismatch5(t *testing.T) {
	t.Parallel()

	r := Mismatch5(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Quintet[int, int, int, int, int]{{42, 42, 42, 42, 69}, {1, 1, 1, 1, 1}})
}

func TestMismatch6(t *testing.T) {
	t.Parallel()

	r := Mismatch6(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Sextet[int, int, int, int, int, int]{
		{42, 42, 42, 42, 42, 69},
		{1, 1, 1, 1, 1, 1},
	})
}

func TestMismatch7(t *testing.T) {
	t.Parallel()

	r := Mismatch7(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Septet[int, int, int, int, int, int, int]{
		{42, 42, 42, 42, 42, 42, 69},
		{1, 1, 1, 1, 1, 1, 1},
	})
}

func TestMismatch8(t *testing.T) {
	t.Parallel()

	r := Mismatch8(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Octet[int, int, int, int, int, int, int, int]{
		{42, 42, 42, 42, 42, 42, 42, 69},
		{1, 1, 1, 1, 1, 1, 1, 1},
	})
}

func TestMismatch9(t *testing.T) {
	t.Parallel()

	r := Mismatch9(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Ennead[int, int, int, int, int, int, int, int, int]{
		{42, 42, 42, 42, 42, 42, 42, 42, 69},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	})
}

func TestMismatch10(t *testing.T) {
	t.Parallel()

	r := Mismatch10(
		func(a, b int) bool { return a == b },
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 42, 1)),
		I(Only(1, 2, 69, 1)),
	)

	assertEqual(t, Slice(r), []Decade[int, int, int, int, int, int, int, int, int, int]{
		{42, 42, 42, 42, 42, 42, 42, 42, 42, 69},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	})
}
