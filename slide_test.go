package ranges

import "testing"

func TestSlide(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(MapF(SlideStep(F(Iota(5)), 3, 3), SliceF[int])),
		[][]int{{0, 1, 2}, {3, 4}},
	)
	assertEqual(
		t,
		SliceF(MapF(Slide(F(Iota(5)), 3), SliceF[int])),
		[][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}},
	)
	assertEqual(
		t,
		SliceF(MapF(SlideStep(F(Iota(5)), 3, 2), SliceF[int])),
		[][]int{{0, 1, 2}, {2, 3, 4}},
	)
	assertEqual(
		t,
		SliceF(MapF(SlideStep(F(Iota(5)), 2, 2), SliceF[int])),
		[][]int{{0, 1}, {2, 3}},
	)
}

func TestSlidePoppingAndSaving(t *testing.T) {
	t.Parallel()

	r := Slide(F(Iota(5)), 3)
	r2 := r.Save()

	r.PopFront()

	assertEqual(t, SliceF(r.Front()), []int{1, 2, 3})
	assertEqual(t, SliceF(r2.Front()), []int{0, 1, 2})
}

func TestSlideS(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(MapF(SlideStepS([]int{0, 1, 2, 3, 4}, 3, 3), SliceF[int])),
		[][]int{{0, 1, 2}, {3, 4}},
	)
	assertEqual(
		t,
		SliceF(MapF(SlideS([]int{0, 1, 2, 3, 4}, 3), SliceF[int])),
		[][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}},
	)
}

func TestSlideSizeZeroPanics(t *testing.T) {
	t.Parallel()

	assertPanic(t, "windowSize < 1 for Slide", func() { Slide(F(Null[int]()), 0) })
}
