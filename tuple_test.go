package ranges

import "testing"

func TestPair(t *testing.T) {
	t.Parallel()

	a, b := MakePair('a', 2).Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
}

func TestTriplet(t *testing.T) {
	t.Parallel()

	a, b, c := MakeTriplet('a', 2, 'c').Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
}

func TestQuartet(t *testing.T) {
	t.Parallel()

	a, b, c, d := MakeQuartet('a', 2, 'c', 4).Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
}

func TestQuintet(t *testing.T) {
	t.Parallel()

	a, b, c, d, e := MakeQuintet('a', 2, 'c', 4, 'e').Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
}

func TestSextet(t *testing.T) {
	t.Parallel()

	a, b, c, d, e, f := MakeSextet('a', 2, 'c', 4, 'e', 6).Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
	assertEqual(t, f, 6)
}

func TestSeptet(t *testing.T) {
	t.Parallel()

	a, b, c, d, e, f, g := MakeSeptet('a', 2, 'c', 4, 'e', 6, 'g').Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
	assertEqual(t, f, 6)
	assertEqual(t, g, 'g')
}

func TestOctet(t *testing.T) {
	t.Parallel()

	a, b, c, d, e, f, g, h := MakeOctet('a', 2, 'c', 4, 'e', 6, 'g', 8).Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
	assertEqual(t, f, 6)
	assertEqual(t, g, 'g')
	assertEqual(t, h, 8)
}

func TestEnnead(t *testing.T) {
	t.Parallel()

	a, b, c, d, e, f, g, h, i := MakeEnnead('a', 2, 'c', 4, 'e', 6, 'g', 8, 'i').Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
	assertEqual(t, f, 6)
	assertEqual(t, g, 'g')
	assertEqual(t, h, 8)
	assertEqual(t, i, 'i')
}

func TestDecade(t *testing.T) {
	t.Parallel()

	a, b, c, d, e, f, g, h, i, j := MakeDecade('a', 2, 'c', 4, 'e', 6, 'g', 8, 'i', 10).Get()

	assertEqual(t, a, 'a')
	assertEqual(t, b, 2)
	assertEqual(t, c, 'c')
	assertEqual(t, d, 4)
	assertEqual(t, e, 'e')
	assertEqual(t, f, 6)
	assertEqual(t, g, 'g')
	assertEqual(t, h, 8)
	assertEqual(t, i, 'i')
	assertEqual(t, j, 10)
}
