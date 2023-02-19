package ranges

import "testing"

func TestMax(t *testing.T) {
	t.Parallel()

	assertEqual(t, Max(1, 2), 2)
	assertEqual(t, Max(2, 1), 2)
	assertEqual(t, Max(2, 1, 3), 3)
	assertEqual(t, Max(7, -1, 2), 7)
}

func TestMin(t *testing.T) {
	t.Parallel()

	assertEqual(t, Min(1, 2), 1)
	assertEqual(t, Min(2, 1), 1)
	assertEqual(t, Min(2, 1, 3), 1)
	assertEqual(t, Min(7, -1, 2), -1)
}
