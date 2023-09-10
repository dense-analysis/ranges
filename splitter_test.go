package ranges

import (
	"testing"
)

func TestSplitter(t *testing.T) {
	t.Parallel()

	r := Splitter(
		Runes("123,456,789"),
		Eq[rune],
		Runes(","),
	)

	assertNotEmpty(t, r)
	assertEqual(t, string(SliceF(r.Front())), "123")

	r.PopFront()

	assertNotEmpty(t, r)
	assertEqual(t, string(SliceF(r.Front())), "456")

	r.PopFront()

	assertNotEmpty(t, r)
	assertEqual(t, string(SliceF(r.Front())), "789")

	r.PopFront()

	assertEmpty(t, r)

	assertPanic(t, "PopFront() called on an empty Splitter() range", func() { r.PopFront() })
}

func TestSplitterSave(t *testing.T) {
	t.Parallel()

	r := Splitter(
		Runes("123,456,789"),
		Eq[rune],
		Runes(","),
	)

	r1 := r.Save()
	r.PopFront()
	r2 := r.Save()
	r.PopFront()
	r3 := r.Save()
	r.PopFront()

	assertEqual(t, string(SliceF(r1.Front())), "123")
	assertEqual(t, string(SliceF(r2.Front())), "456")
	assertEqual(t, string(SliceF(r3.Front())), "789")
}

func TestSplitterToSlice(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			Splitter(
				Runes("123,456,789"),
				Eq[rune],
				Runes(","),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"123",
		"456",
		"789",
	})
}

func TestSplitterLeadingEmptySplit(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			Splitter(
				Only(0, 4, 5, 6, 0, 7, 8, 9),
				Eq[int],
				Only(0),
			),
			SliceF[int],
		),
	)

	assertEqual(t, result, [][]int{
		{},
		{4, 5, 6},
		{7, 8, 9},
	})
}

func TestSplitterTrailingEmptySplit(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			Splitter(
				Only(4, 5, 6, 0, 7, 8, 9, 0),
				Eq[int],
				Only(0),
			),
			SliceF[int],
		),
	)

	assertEqual(t, result, [][]int{
		{4, 5, 6},
		{7, 8, 9},
		{},
	})
}

func TestSplitterMultipleCharacters(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			Splitter(
				Runes("<>12<3<>45>6<>789<>"),
				Eq[rune],
				Runes("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitterEmpty(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			Splitter(Runes(""), Eq[rune], Runes(",")),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{""})
}

func TestSplitterS(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			SplitterS(
				[]rune("<>12<3<>45>6<>789<>"),
				Eq[rune],
				Runes("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitterSS(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			SplitterSS(
				[]rune("<>12<3<>45>6<>789<>"),
				Eq[rune],
				[]rune("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitterComparable(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			SplitterComparable(
				Runes("<>12<3<>45>6<>789<>"),
				Runes("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitterComparableS(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			SplitterComparableS(
				[]rune("<>12<3<>45>6<>789<>"),
				Runes("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitterComparableSS(t *testing.T) {
	t.Parallel()

	result := SliceF(
		MapF(
			SplitterComparableSS(
				[]rune("<>12<3<>45>6<>789<>"),
				[]rune("<>"),
			),
			Pipe2(SliceF[rune], StringS),
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitString(t *testing.T) {
	t.Parallel()

	result := SliceF(
		SplitString(
			"<>12<3<>45>6<>789<>",
			"<>",
		),
	)

	assertEqual(t, result, []string{
		"",
		"12<3",
		"45>6",
		"789",
		"",
	})
}

func TestSplitStringSave(t *testing.T) {
	t.Parallel()

	result := SplitString("<>12<3<>45>6<>789<>", "<>")

	r1 := result.Save()
	result.PopFront()
	r2 := result.Save()
	result.PopFront()
	r3 := result.Save()
	result.PopFront()
	r4 := result.Save()
	result.PopFront()
	r5 := result.Save()
	result.PopFront()

	assertEqual(t, r1.Front(), "")
	assertEqual(t, r2.Front(), "12<3")
	assertEqual(t, r3.Front(), "45>6")
	assertEqual(t, r4.Front(), "789")
	assertEqual(t, r5.Front(), "")
}

func TestSplitWhen(t *testing.T) {
	t.Parallel()

	result := Slice(
		Map(
			SplitWhen(
				SliceRange([]int{4, 3, 2, 11, 0, -3, -3, 5, 3, 0}),
				func(a, b int) bool { return a <= b },
			),
			Slice[int],
		),
	)

	assertEqual(t, result, [][]int{
		{4, 3, 2},
		{11, 0, -3},
		{-3},
		{5, 3, 0},
	})
}
