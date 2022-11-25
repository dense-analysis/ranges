package ranges

import "testing"

func TestMap(t *testing.T) {
	t.Parallel()

	r := Only(1, 2)
	sliceCopy := Slice(Map[int](r, func(element int) float64 { return float64(element * 2) }))

	assertEqual(t, sliceCopy, []float64{2.0, 4.0})
}

func TestMapF(t *testing.T) {
	t.Parallel()

	r := F(Only(1, 2))
	fr := MapF(r, func(element int) float64 { return float64(element * 2) })
	fr2 := fr.Save()

	fr.PopFront()

	assertHasFrontF(t, fr2, 2.0)

	fr2.PopFront()

	assertHasFrontF(t, fr2, 4.0)
}
