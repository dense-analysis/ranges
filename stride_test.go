package ranges

// Tests for stride rely slices.go, and iota.go.
import "testing"

func TestStride(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		Slice(Stride[int](Iota(10), 1)),
		SliceB(Iota(10)),
	)
	assertEqual(
		t,
		Slice(Stride[int](Iota(10), 3)),
		[]int{0, 3, 6, 9},
	)
}

func TestStrideF(t *testing.T) {
	t.Parallel()

	assertEqual(
		t,
		SliceF(StrideF(Iota(10), 1)),
		SliceB(Iota(10)),
	)
	assertEqual(
		t,
		SliceF(StrideF(Iota(10), 3)),
		[]int{0, 3, 6, 9},
	)
}

func TestStrideS(t *testing.T) {
	t.Parallel()

	assertEmpty(t, StrideS([]int{}, 1))

	assertEqual(
		t,
		SliceF(StrideS(SliceF(Iota(10)), 1)),
		SliceB(Iota(10)),
	)
	assertEqual(
		t,
		SliceF(StrideS(SliceB(Iota(10)), 3)),
		[]int{0, 3, 6, 9},
	)
}

func TestStrideInvalidStepPanic(t *testing.T) {
	t.Parallel()

	assertPanic(t, "step < 1 for Stride", func() { Stride[int](Iota(0), 0) })
	assertPanic(t, "step < 1 for Stride", func() { StrideF(Iota(0), 0) })
	assertPanic(t, "step < 1 for Stride", func() { StrideS([]int{}, 0) })
}
