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

	r := Only(1, 2)
	fr := MapF(r, func(element int) float64 { return float64(element * 2) })
	fr2 := fr.Save()

	fr.PopFront()

	assertHasFrontF(t, fr2, 2.0)

	fr2.PopFront()

	assertHasFrontF(t, fr2, 4.0)
}

func TestMapB(t *testing.T) {
	t.Parallel()

	r := Only(1, 2)
	br := MapB(r, func(element int) float64 { return float64(element * 2) })
	br2 := br.SaveB()

	br.PopBack()

	assertHasBack(t, br2, 4.0)

	br2.PopBack()

	assertHasBack(t, br2, 2.0)
}

func TestMapR(t *testing.T) {
	r := Only(1, 2)
	rr := MapR(r, func(element int) float64 { return float64(element * 2) })

	assertEqual(t, rr.Len(), 2)
	assertEqual(t, rr.Get(0), 2.0)
	assertEqual(t, rr.Get(1), 4.0)

	// Set saving the range works.
	rr2 := rr.SaveR()
	rr.PopFront()

	assertEqual(t, rr2.Front(), 2.0)
	rr2.PopFront()
	assertEqual(t, rr2.Front(), 4.0)
	assertEqual(t, rr2.Len(), 1)
	assertEqual(t, rr2.Get(0), 4.0)
}
