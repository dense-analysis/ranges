package ranges

import "testing"

func TestGroup(t *testing.T) {
	t.Parallel()

	result := Slice(
		Group(
			Only("a", "b", "b", "b", "c", "c", "d"),
			func(a, b string) bool { return a == b },
		),
	)

	assertEqual(t, result, []Pair[string, int]{
		{"a", 1},
		{"b", 3},
		{"c", 2},
		{"d", 1},
	})
}

func TestGroupS(t *testing.T) {
	t.Parallel()

	result := Slice(
		GroupS(
			[]string{"a", "b", "b", "b", "c", "c", "d"},
			func(a, b string) bool { return a == b },
		),
	)

	assertEqual(t, result, []Pair[string, int]{
		{"a", 1},
		{"b", 3},
		{"c", 2},
		{"d", 1},
	})
}

func TestGroupComparable(t *testing.T) {
	t.Parallel()

	result := Slice(
		GroupComparable(Only("a", "b", "b", "b", "c", "c", "d")),
	)

	assertEqual(t, result, []Pair[string, int]{
		{"a", 1},
		{"b", 3},
		{"c", 2},
		{"d", 1},
	})
}

func TestGroupComparableS(t *testing.T) {
	t.Parallel()

	result := Slice(
		GroupComparableS([]string{"a", "b", "b", "b", "c", "c", "d"}),
	)

	assertEqual(t, result, []Pair[string, int]{
		{"a", 1},
		{"b", 3},
		{"c", 2},
		{"d", 1},
	})
}
