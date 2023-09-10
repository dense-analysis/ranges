package ranges

import "testing"

func TestFind(t *testing.T) {
	t.Parallel()

	result := Slice(Find(Only(1, 2, 3, 4), Ge[int], 3))
	assertEqual(t, result, []int{3, 4})
}

func TestFindF(t *testing.T) {
	t.Parallel()

	assertHasSaveableFront(t, FindF(Only(1, 2, 3, 4), Ge[int], 3), 3)
}

func TestFindS(t *testing.T) {
	t.Parallel()

	assertHasSaveableFront(t, FindS([]int{1, 2, 3, 4}, Ge[int], 3), 3)
}

func TestFindComparable(t *testing.T) {
	t.Parallel()

	assertEqual(t, Slice(FindComparable(Only(1, 2, 3, 4), 3)), []int{3, 4})
}

func TestFindComparableF(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(FindComparableF(Only(1, 2, 3, 4), 3)), []int{3, 4})
}

func TestFindComparableS(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(FindComparableS([]int{1, 2, 3, 4}, 3)), []int{3, 4})
}

func TestFindEqual(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqual(Runes("hello world"), Eq[rune], Runes("or")))

	assertEqual(t, string(result), "orld")

	result2 := SliceF(FindEqual(Only(1, 3, 2, 3, 4, 5), Eq[int], Only(3, 4)))

	assertEqual(t, result2, []int{3, 4, 5})
}

func TestFindEqualS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualS([]rune("hello world"), Eq[rune], Runes("or")))

	assertEqual(t, string(result), "orld")
}

func TestFindEqualComparable(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualComparable(Runes("hello world"), Runes("or")))

	assertEqual(t, string(result), "orld")

	result2 := SliceF(FindEqualComparable(Only(1, 3, 2, 3, 4, 5), Only(3, 4)))

	assertEqual(t, result2, []int{3, 4, 5})
}

func TestFindEqualComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualComparableS([]rune("hello world"), Runes("or")))
	assertEqual(t, string(result), "orld")
}

func TestFindAdjacent(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacent(Runes("abba"), Eq[rune]))

	assertEqual(t, string(result), "bba")

	result2 := SliceF(FindAdjacent(Runes("abc"), Eq[rune]))

	assertEqual(t, string(result2), "")

	result3 := SliceF(FindAdjacent(Only(1, 2, 3, 4, 4, 5), Eq[int]))

	assertEqual(t, result3, []int{4, 4, 5})
}

func TestFindAdjacentS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentS([]rune("abba"), Eq[rune]))
	assertEqual(t, string(result), "bba")
}

func TestFindAdjacentComparable(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentComparable(Runes("abba")))

	assertEqual(t, string(result), "bba")

	result2 := SliceF(FindAdjacentComparable(Runes("abc")))

	assertEqual(t, string(result2), "")

	result3 := SliceF(FindAdjacentComparable(Only(1, 2, 3, 4, 4, 5)))

	assertEqual(t, result3, []int{4, 4, 5})
}

func TestFindAdjacentComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentComparableS([]rune("abba")))
	assertEqual(t, string(result), "bba")
}

func TestFindAmong(t *testing.T) {
	t.Parallel()

	result := Slice(FindAmong(Runes("abcd"), Eq[rune], Runes("qcx")))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongF(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongF(Runes("abcd"), Eq[rune], Runes("qcx")))

	assertEqual(t, string(result), "cd")
	assertHasSaveableFront(t, FindAmongF(Runes("abcd"), Eq[rune], Runes("qcx")), 'c')
}

func TestFindAmongS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongS([]rune("abcd"), Eq[rune], Runes("qcx")))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongComparable(t *testing.T) {
	t.Parallel()

	result := Slice(FindAmongComparable(Runes("abcd"), Runes("qcx")))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongComparableF(Runes("abcd"), Runes("qcx")))

	assertEqual(t, string(result), "cd")
	assertHasSaveableFront(t, FindAmongComparableF(Runes("abcd"), Runes("qcx")), 'c')
}

func TestFindAmongComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongComparableS([]rune("abcd"), Runes("qcx")))

	assertEqual(t, string(result), "cd")
}

func TestCanFind(t *testing.T) {
	t.Parallel()

	if !CanFind(Only(1, 2, 3, 4), Ge[int], 3) {
		t.Error("Couldn't find x >= 3 in (1, 2, 3, 4)")
	}

	if CanFind(Only(1, 2, 3, 4), Ge[int], 10) {
		t.Error("Incorrectly found x >= 10 in (1, 2, 3, 4)")
	}
}

func TestCanFindS(t *testing.T) {
	t.Parallel()

	if !CanFindS([]int{1, 2, 3, 4}, Ge[int], 3) {
		t.Error("Couldn't find x >= 3 in (1, 2, 3, 4)")
	}
}

func TestCanFindComparable(t *testing.T) {
	t.Parallel()

	if !CanFindComparable(Only(1, 2, 3, 4), 3) {
		t.Error("Couldn't find x == 3 in (1, 2, 3, 4)")
	}

	if CanFindComparable(Only(1, 2, 3, 4), 10) {
		t.Error("Incorrectly found x == 10 in (1, 2, 3, 4)")
	}
}

func TestCanFindComparableS(t *testing.T) {
	t.Parallel()

	if !CanFindComparableS([]int{1, 2, 3, 4}, 3) {
		t.Error("Couldn't find x == 3 in (1, 2, 3, 4)")
	}
}
