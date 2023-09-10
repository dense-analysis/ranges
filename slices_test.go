package ranges

import (
	"testing"
)

func TestSliceRange(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})

	sliceCopy := make([]int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Front())
		r.PopFront()
	}

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestEmptyNilSliceRange(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int(nil))

	assertEmptyR(t, r)
}

func TestSlicePopFront(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})

	r.PopFront()

	sliceCopy := make([]int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Front())
		r.PopFront()
	}

	assertEqual(t, sliceCopy, []int{2, 3})
}

func TestSlicePopBack(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})

	r.PopBack()

	sliceCopy := make([]int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Back())
		r.PopBack()
	}

	assertEqual(t, sliceCopy, []int{2, 1})
}

func TestSliceRangeSave(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})
	r2 := r.Save()

	r.PopFront()
	r.PopFront()
	r.PopFront()

	if r2.Empty() || r2.Front() != 1 {
		t.Fatal("The range was not saved")
	}
}

func TestSliceRangeSaveB(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})
	r2 := r.SaveB()

	r.PopFront()
	r.PopFront()
	r.PopBack()

	if r2.Empty() || r2.Front() != 1 || r2.Back() != 3 {
		t.Fatal("The range was not saved")
	}
}

func TestSliceRetroRange(t *testing.T) {
	t.Parallel()

	r := SliceRetroRange([]int{1, 2, 3})

	sliceCopy := make([]int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Front())
		r.PopFront()
	}

	assertEqual(t, sliceCopy, []int{3, 2, 1})
}

func TestSliceRetroRangeBack(t *testing.T) {
	t.Parallel()

	r := SliceRetroRange([]int{1, 2, 3})

	sliceCopy := make([]int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Back())
		r.PopBack()
	}

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestEmptyNilSliceRetroRange(t *testing.T) {
	t.Parallel()

	r := SliceRetroRange([]int(nil))

	assertEmptyR(t, r)
}

func TestSliceRetroRangeSave(t *testing.T) {
	t.Parallel()

	r := SliceRetroRange([]int{1, 2, 3})
	r2 := r.Save()

	r.PopFront()
	r.PopFront()
	r.PopFront()

	if r2.Empty() || r2.Front() != 3 {
		t.Fatal("The range was not saved")
	}
}

func TestSliceRetroRangeSaveB(t *testing.T) {
	t.Parallel()

	r := SliceRetroRange([]int{1, 2, 3})
	r2 := r.SaveB()

	r.PopFront()
	r.PopFront()
	r.PopBack()

	if r2.Empty() || r2.Front() != 3 || r2.Back() != 1 {
		t.Fatal("The range was not saved")
	}
}

func TestSlicePtrRange(t *testing.T) {
	t.Parallel()

	r := SlicePtrRange([]int{1, 2, 3})

	sliceCopy := make([]*int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Front())
		r.PopFront()
	}

	if len(sliceCopy) != 3 || *sliceCopy[0] != 1 || *sliceCopy[1] != 2 || *sliceCopy[2] != 3 {
		t.Fatal("sliceCopy != []int{&1, &2, &3}")
	}
}

func TestSlicePtrRangeBack(t *testing.T) {
	t.Parallel()

	r := SlicePtrRange([]int{1, 2, 3})

	sliceCopy := make([]*int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Back())
		r.PopBack()
	}

	if len(sliceCopy) != 3 || *sliceCopy[0] != 3 || *sliceCopy[1] != 2 || *sliceCopy[2] != 1 {
		t.Fatal("sliceCopy != []int{&3, &2, &1}")
	}
}

func TestEmptyNilSlicePtrRange(t *testing.T) {
	t.Parallel()

	r := SlicePtrRange([]int(nil))

	assertEmptyR(t, r)
}

func TestSlicePtrRangeSave(t *testing.T) {
	t.Parallel()

	r := SlicePtrRange([]int{1, 2, 3})
	r2 := r.Save()

	r.PopFront()
	r.PopFront()
	r.PopFront()

	if r2.Empty() || *r2.Front() != 1 {
		t.Fatal("The range was not saved")
	}
}

func TestSlicePtrRangeSaveB(t *testing.T) {
	t.Parallel()

	r := SlicePtrRange([]int{1, 2, 3})
	r2 := r.SaveB()

	r.PopFront()
	r.PopFront()
	r.PopBack()

	if r2.Empty() || *r2.Front() != 1 || *r2.Back() != 3 {
		t.Fatal("The range was not saved")
	}
}

func TestSlicePtrRetroRange(t *testing.T) {
	t.Parallel()

	r := SlicePtrRetroRange([]int{1, 2, 3})

	sliceCopy := make([]*int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Front())
		r.PopFront()
	}

	if len(sliceCopy) != 3 || *sliceCopy[0] != 3 || *sliceCopy[1] != 2 || *sliceCopy[2] != 1 {
		t.Fatal("sliceCopy != []int{&3, &2, &1}")
	}
}

func TestSlicePtrRetroRangeBack(t *testing.T) {
	t.Parallel()

	r := SlicePtrRetroRange([]int{1, 2, 3})

	sliceCopy := make([]*int, 0)

	for !r.Empty() {
		sliceCopy = append(sliceCopy, r.Back())
		r.PopBack()
	}

	if len(sliceCopy) != 3 || *sliceCopy[0] != 1 || *sliceCopy[1] != 2 || *sliceCopy[2] != 3 {
		t.Fatal("sliceCopy != []int{&1, &2, &3}")
	}
}

func TestEmptyNilSlicePtrRetroRange(t *testing.T) {
	t.Parallel()

	r := SlicePtrRetroRange([]int(nil))

	assertEmptyR(t, r)
}

func TestSlicePtrRetroRangeSave(t *testing.T) {
	t.Parallel()

	r := SlicePtrRetroRange([]int{1, 2, 3})
	r2 := r.Save()

	r.PopFront()
	r.PopFront()
	r.PopFront()

	if r2.Empty() || *r2.Front() != 3 {
		t.Fatal("The range was not saved")
	}
}

func TestSlicePtrRetroRangeSaveB(t *testing.T) {
	t.Parallel()

	r := SlicePtrRetroRange([]int{1, 2, 3})
	r2 := r.SaveB()

	r.PopFront()
	r.PopFront()
	r.PopBack()

	if r2.Empty() || *r2.Front() != 3 || *r2.Back() != 1 {
		t.Fatal("The range was not saved")
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	sliceCopy := Slice[int](SliceRange([]int{1, 2, 3}))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestSliceF(t *testing.T) {
	t.Parallel()

	sliceCopy := SliceF(SliceRange([]int{1, 2, 3}))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestSliceB(t *testing.T) {
	t.Parallel()

	sliceCopy := SliceB(SliceRange([]int{1, 2, 3}))

	assertEqual(t, sliceCopy, []int{1, 2, 3})
}

func TestSliceR(t *testing.T) {
	t.Parallel()

	r := SliceRange([]int{1, 2, 3})
	sliceCopy := SliceR(r)

	assertEqual(t, sliceCopy, []int{1, 2, 3})
	// We should consume the range when we copy it.
	assertEmptyR(t, r)
}

// Test how we handle ranges with the wrong lengths for SliceR
func TestSliceRWithInvalidLengthRange(t *testing.T) {
	t.Parallel()

	// We should take more elements than the length says in case of a short length.
	sliceCopy := SliceR(badLengthRange(SliceRange([]int{1, 2, 3}), 2))
	assertEqual(t, sliceCopy, []int{1, 2, 3})

	// We should take as many as there are in case of a too long length.
	sliceCopy2 := SliceR(badLengthRange(SliceRange([]int{1, 2, 3}), 4))
	assertEqual(t, sliceCopy2, []int{1, 2, 3})
}

func TestBytes(t *testing.T) {
	t.Parallel()

	if string(SliceR(Bytes("abc"))) != "abc" {
		t.Error("Bytes did not represent the string correctly")
	}
}

func TestRunes(t *testing.T) {
	t.Parallel()

	r := Runes("日本語")

	if r.Empty() || r.Front() != '日' {
		t.Fatal("runes[0] != '日'")
	}

	r.PopFront()

	if r.Empty() || r.Front() != '本' {
		t.Fatal("runes[1] != '本'")
	}

	r.PopFront()

	if r.Empty() || r.Front() != '語' {
		t.Fatal("runes[2] != '語'")
	}

	r.PopFront()

	assertEmptyR(t, r)
}

func TestRunesBack(t *testing.T) {
	t.Parallel()

	r := Runes("日本語")

	if r.Empty() || r.Back() != '語' {
		t.Fatal("runes[2] != '語'")
	}

	r.PopBack()

	if r.Empty() || r.Back() != '本' {
		t.Fatal("runes[1] != '本'")
	}

	r.PopBack()

	if r.Empty() || r.Back() != '日' {
		t.Fatal("runes[0] != '日'")
	}

	r.PopBack()

	assertEmptyR(t, r)
}

func TestRunesToString(t *testing.T) {
	t.Parallel()

	if string(SliceB(Runes("日本語"))) != "日本語" {
		t.Error("We couldn't covert a string to runes and back again")
	}

	if String(Runes("日本語")) != "日本語" {
		t.Error("We couldn't covert a string to runes and back again")
	}
}

func TestRunsToStringRandomAccessSpecialization(t *testing.T) {
	t.Parallel()

	if string(SliceR(Runes("日本語"))) != "日本語" {
		t.Error("We couldn't covert a string to runes and back again")
	}

	if String(Runes("日本語")) != "日本語" {
		t.Error("We couldn't covert a string to runes and back again")
	}
}
