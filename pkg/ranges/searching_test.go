package ranges

import "testing"

func TestAll(t *testing.T) {
	t.Parallel()

	if !All[int](Only(2, 4), func(x int) bool { return x%2 == 0 }) {
		t.Fatal("All did not return a positive result.")
	}

	if All[int](Only(2, 1), func(x int) bool { return x%2 == 0 }) {
		t.Fatal("All did not return a negative result.")
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	if Any[int](Only(1, 3), func(x int) bool { return x%2 == 0 }) {
		t.Fatal("Any did not return a negative result.")
	}

	if !Any[int](Only(2, 1), func(x int) bool { return x%2 == 0 }) {
		t.Fatal("Any did not return a positive result.")
	}
}

func TestAmong(t *testing.T) {
	t.Parallel()

	if !Among(func(a, b int) bool { return a == b }, 1, 1, 2, 3) {
		t.Fatal("1 is not in (1, 2, 3)")
	}

	if Among(func(a, b int) bool { return a == b }, 4, 1, 2, 3) {
		t.Fatal("4 was in (1, 2, 3)")
	}

	if Among(func(a, b int) bool { return a == b }, 1) {
		t.Fatal("1 is in ()")
	}
}

func TestAmongEq(t *testing.T) {
	t.Parallel()

	if !AmongEq(1, 1, 2, 3) {
		t.Fatal("1 is not in (1, 2, 3)")
	}

	if AmongEq(4, 1, 2, 3) {
		t.Fatal("4 was in (1, 2, 3)")
	}
}

func TestStartsWith(t *testing.T) {
	t.Parallel()

	if !StartsWith[int, float64](Null[int](), Null[float64](), nil) {
		t.Fatal("Two empty ranges do not satisfy StartsWith")
	}

	if !StartsWith[int, float64](Only(1), Null[float64](), nil) {
		t.Fatal("A non-empty range and empty range did not satisfy StartsWith")
	}

	if StartsWith[int, float64](Null[int](), Only(1.0), nil) {
		t.Fatal("An empty range should result in false if the search range is non-empty")
	}

	if !StartsWith[int, float64](Only(1, 2, 3), Only(1.0, 2.0), func(x int, y float64) bool { return x == int(y) }) {
		t.Fatal("[]int{1, 2, 3} did not contain []float64{1.0, 2.0}")
	}

	if StartsWith[int, float64](Only(1, 2, 3), Only(2.0), func(x int, y float64) bool { return x == int(y) }) {
		t.Fatal("[]int{1, 2, 3} should not start with []float64{2.0}")
	}
}

func TestSkipOver(t *testing.T) {
	t.Parallel()

	empty := Runes("")

	if !SkipOver(&empty, I(Runes("xyz")), Eq[rune]) {
		t.Error("SkipOver(\"\", \"xyz\") != true")
	}

	if !SkipOver(nil, I(Runes("xyz")), Eq[rune]) {
		t.Error("SkipOver(nil, \"xyz\") != true")
	}

	r1 := Runes("Hello world")

	if SkipOver(&r1, I(Runes("Ha")), Eq[rune]) {
		t.Fatal("SkipOver(\"Hello world\", \"Ha\") != false")
	}

	if !SkipOver(&r1, I(Runes("Hell")), Eq[rune]) {
		t.Fatal("SkipOver(\"Hello world\", \"Hell\") != true")
	}

	assertEqual(t, String(r1), "o world")
}

func TestLength(t *testing.T) {
	t.Parallel()

	assertEqual(t, Length(I(Only[int]())), 0)
	assertEqual(t, Length(I(Only(1))), 1)
	assertEqual(t, Length(I(Only(4, 5, 6))), 3)
}

func TestCount(t *testing.T) {
	t.Parallel()

	r := Only(1, 2, 3, 4, 5)
	res1 := Count[int](r, func(x int) bool { return x%2 == 0 })

	if res1 != 2 {
		t.Fatalf("Count counted %d instead of 2 even elements", res1)
	}

	if !r.Empty() {
		t.Fatal("The range was not exhaustively searched")
	}

	res2 := Count[int](Only(1, 2, 3, 4, 5), func(x int) bool { return x%2 == 1 })

	if res2 != 3 {
		t.Fatalf("Count counted %d instead of 3 odd elements", res2)
	}
}

func TestCountUntil(t *testing.T) {
	t.Parallel()

	r := Only(4, 2, 3, 4, 5)
	res1 := CountUntil[int](r, func(x int) bool { return x%2 == 1 })

	if res1 != 2 {
		t.Fatalf("CountUntil counted %d instead of 2 even elements", res1)
	}

	assertHasFrontF(t, r, 3)

	res2 := CountUntil[int](Only(1, 1, 3, 4, 5), func(x int) bool { return x%2 == 0 })

	if res2 != 3 {
		t.Fatalf("CountUntil counted %d instead of 3 odd elements", res2)
	}
}
