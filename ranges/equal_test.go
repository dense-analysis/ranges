package ranges

import "testing"

func TestEqualComparableInts(t *testing.T) {
	t.Parallel()

	if !EqualComparable[int](
		SliceRange([]int{2, 5, 7}),
		SliceRange([]int{2, 5, 7}),
	) {
		t.Error("{2, 5, 7} comparison failed")
	}

	if EqualComparable[int](
		SliceRange([]int{2, 5}),
		SliceRange([]int{2, 5, 7}),
	) {
		t.Error("Unbalanced comparison failed")
	}

	if EqualComparable[int](
		SliceRange([]int{2, 5, 7}),
		SliceRange([]int{2, 5}),
	) {
		t.Error("Unbalanced comparison failed")
	}

	if !EqualComparable[int](Null[int](), Null[int]()) {
		t.Error("Empty comparison failed")
	}
}

func TestEqualComparableSInts(t *testing.T) {
	t.Parallel()

	if !EqualComparableS([]int{2, 5, 7}, []int{2, 5, 7}) {
		t.Error("{2, 5, 7} comparison failed")
	}

	if EqualComparableS([]int{2, 5}, []int{2, 5, 7}) {
		t.Error("Unbalanced comparison failed")
	}

	if EqualComparableS([]int{2, 5, 7}, []int{2, 5}) {
		t.Error("Unbalanced comparison failed")
	}

	if !EqualComparableS([]int{}, []int{}) {
		t.Error("Empty comparison failed")
	}
}

func TestEqualInts(t *testing.T) {
	t.Parallel()

	if Equal[int](
		SliceRange([]int{2, 5, 7}),
		SliceRange([]int{2, 5, 7}),
		// The condition is inverted here.
		func(a, b int) bool { return a != b },
	) {
		t.Error("{2, 5, 7} comparison failed")
	}

	if Equal[int](
		SliceRange([]int{2, 5}),
		SliceRange([]int{2, 5, 7}),
		Eq[int],
	) {
		t.Error("Unbalanced comparison failed")
	}

	if Equal[int](
		SliceRange([]int{2, 5, 7}),
		SliceRange([]int{2, 5}),
		Eq[int],
	) {
		t.Error("Unbalanced comparison failed")
	}

	if !EqualComparable[int](Null[int](), Null[int]()) {
		t.Error("Empty comparison failed")
	}
}

func TestEqualSInts(t *testing.T) {
	t.Parallel()

	if !EqualS([]int{2, 5, 7}, []int{2, 5, 7}, Eq[int]) {
		t.Error("{2, 5, 7} comparison failed")
	}

	if EqualS([]int{2, 5}, []int{2, 5, 7}, Eq[int]) {
		t.Error("Unbalanced comparison failed")
	}

	if EqualS([]int{2, 5, 7}, []int{2, 5}, Eq[int]) {
		t.Error("Unbalanced comparison failed")
	}

	if !EqualS([]int{}, []int{}, Eq[int]) {
		t.Error("Empty comparison failed")
	}
}

func TestEqualRunes(t *testing.T) {
	t.Parallel()

	if !Equal(I(F(Runes("abc"))), I(F(Runes("abc"))), Eq[rune]) {
		t.Error("abc comparison failed")
	}
}
