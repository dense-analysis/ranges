package ranges

import "testing"

func TestFind(t *testing.T) {
	t.Parallel()

	result := Slice(Find(I(F(B(Only(1, 2, 3, 4)))), Ge[int], 3))
	assertEqual(t, result, []int{3, 4})
}

func TestFindF(t *testing.T) {
	t.Parallel()

	assertHasSaveableFront(t, FindF(F(B(Only(1, 2, 3, 4))), Ge[int], 3), 3)
}

func TestFindS(t *testing.T) {
	t.Parallel()

	assertHasSaveableFront(t, FindS([]int{1, 2, 3, 4}, Ge[int], 3), 3)
}

func TestFindComparable(t *testing.T) {
	t.Parallel()

	assertEqual(t, Slice(FindComparable(I(F(B(Only(1, 2, 3, 4)))), 3)), []int{3, 4})
}

func TestFindComparableF(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(FindComparableF(F(B(Only(1, 2, 3, 4))), 3)), []int{3, 4})
}

func TestFindComparableS(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(FindComparableS([]int{1, 2, 3, 4}, 3)), []int{3, 4})
}

func TestFindEqual(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqual(F(B(Runes("hello world"))), Eq[rune], F(B(Runes("or")))))

	assertEqual(t, string(result), "orld")

	result2 := SliceF(FindEqual(F(B(Only(1, 3, 2, 3, 4, 5))), Eq[int], F(B(Only(3, 4)))))

	assertEqual(t, result2, []int{3, 4, 5})
}

func TestFindEqualS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualS([]rune("hello world"), Eq[rune], F(B(Runes("or")))))

	assertEqual(t, string(result), "orld")
}

func TestFindEqualComparable(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualComparable(F(B(Runes("hello world"))), F(B(Runes("or")))))

	assertEqual(t, string(result), "orld")

	result2 := SliceF(FindEqualComparable(F(B(Only(1, 3, 2, 3, 4, 5))), F(B(Only(3, 4)))))

	assertEqual(t, result2, []int{3, 4, 5})
}

func TestFindEqualComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindEqualComparableS([]rune("hello world"), F(B(Runes("or")))))
	assertEqual(t, string(result), "orld")
}

func TestFindAdjacent(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacent(F(B(Runes("abba"))), Eq[rune]))

	assertEqual(t, string(result), "bba")

	result2 := SliceF(FindAdjacent(F(B(Runes("abc"))), Eq[rune]))

	assertEqual(t, string(result2), "")

	result3 := SliceF(FindAdjacent(F(B(Only(1, 2, 3, 4, 4, 5))), Eq[int]))

	assertEqual(t, result3, []int{4, 4, 5})
}

func TestFindAdjacentS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentS([]rune("abba"), Eq[rune]))
	assertEqual(t, string(result), "bba")
}

func TestFindAdjacentComparable(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentComparable(F(B(Runes("abba")))))

	assertEqual(t, string(result), "bba")

	result2 := SliceF(FindAdjacentComparable(F(B(Runes("abc")))))

	assertEqual(t, string(result2), "")

	result3 := SliceF(FindAdjacentComparable(F(B(Only(1, 2, 3, 4, 4, 5)))))

	assertEqual(t, result3, []int{4, 4, 5})
}

func TestFindAdjacentComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAdjacentComparableS([]rune("abba")))
	assertEqual(t, string(result), "bba")
}

func TestFindAmong(t *testing.T) {
	t.Parallel()

	result := Slice(FindAmong(I(F(B(Runes("abcd")))), Eq[rune], F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongF(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongF(F(B(Runes("abcd"))), Eq[rune], F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
	assertHasSaveableFront(t, FindAmongF(F(B(Runes("abcd"))), Eq[rune], F(B(Runes("qcx")))), 'c')
}

func TestFindAmongS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongS([]rune("abcd"), Eq[rune], F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongComparable(t *testing.T) {
	t.Parallel()

	result := Slice(FindAmongComparable(I(F(B(Runes("abcd")))), F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
}

func TestFindAmongComparableF(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongComparableF(F(B(Runes("abcd"))), F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
	assertHasSaveableFront(t, FindAmongComparableF(F(B(Runes("abcd"))), F(B(Runes("qcx")))), 'c')
}

func TestFindAmongComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(FindAmongComparableS([]rune("abcd"), F(B(Runes("qcx")))))

	assertEqual(t, string(result), "cd")
}

func TestCanFind(t *testing.T) {
	t.Parallel()

	if !CanFind(I(F(B(Only(1, 2, 3, 4)))), Ge[int], 3) {
		t.Error("Couldn't find x >= 3 in (1, 2, 3, 4)")
	}

	if CanFind(I(F(B(Only(1, 2, 3, 4)))), Ge[int], 10) {
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

	if !CanFindComparable(I(F(B(Only(1, 2, 3, 4)))), 3) {
		t.Error("Couldn't find x == 3 in (1, 2, 3, 4)")
	}

	if CanFindComparable(I(F(B(Only(1, 2, 3, 4)))), 10) {
		t.Error("Incorrectly found x == 10 in (1, 2, 3, 4)")
	}
}

func TestCanFindComparableS(t *testing.T) {
	t.Parallel()

	if !CanFindComparableS([]int{1, 2, 3, 4}, 3) {
		t.Error("Couldn't find x == 3 in (1, 2, 3, 4)")
	}
}
