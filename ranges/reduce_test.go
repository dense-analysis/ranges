package ranges

import "testing"

func TestReduce(t *testing.T) {
	t.Parallel()

	result := Reduce(
		I(F(Only(100, 10, 1))),
		func(a float64, b int) float64 { return a + float64(b) },
		1000.0,
	)

	assertEqual(t, result, 1111.0)
}

func TestReduceS(t *testing.T) {
	t.Parallel()

	result := ReduceS(
		[]int{100, 10, 1},
		func(a float64, b int) float64 { return a + float64(b) },
		1000.0,
	)

	assertEqual(t, result, 1111.0)
}

func TestReduceNoSeed(t *testing.T) {
	t.Parallel()

	result := ReduceNoSeed(
		I(F(Only(100, 10, 1))),
		func(a, b int) int { return a + b },
	)

	assertEqual(t, result, 111)
}

func TestReduceNoSeedS(t *testing.T) {
	t.Parallel()

	result := ReduceNoSeedS(
		[]int{100, 10, 1},
		func(a, b int) int { return a + b },
	)

	assertEqual(t, result, 111)
}
