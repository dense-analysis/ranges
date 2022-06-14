package ranges

import "testing"

func TestEach(t *testing.T) {
	t.Parallel()

	values := make([]int, 0, 3)

	Each(I(Only(7, 1, 2)), func(element int) { values = append(values, element) })

	assertEqual(t, values, []int{7, 1, 2})
}

func TestEachS(t *testing.T) {
	t.Parallel()

	values := make([]int, 0, 3)

	EachS([]int{7, 1, 2}, func(element int) { values = append(values, element) })

	assertEqual(t, values, []int{7, 1, 2})
}

func TestExhaust(t *testing.T) {
	t.Parallel()

	r := Only(1, 2, 3)

	Exhaust(I(r))

	assertEmpty(t, I(r))
}
