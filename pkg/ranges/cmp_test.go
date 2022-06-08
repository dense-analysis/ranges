package ranges

import "testing"

func TestCmpFunc(t *testing.T) {
	t.Parallel()

	assertEqual(t, CmpFunc(1, 1), 0)
	assertEqual(t, CmpFunc(1, 2), -1)
	assertEqual(t, CmpFunc(2, 1), 1)
}

func TestCmp(t *testing.T) {
	t.Parallel()

	assertEqual(t, Cmp[int](Only[int](), Only[int](), CmpFunc[int]), 0)
	assertEqual(t, Cmp[int](Only[int](), Only(1), CmpFunc[int]), -1)
	assertEqual(t, Cmp[int](Only(1), Only[int](), CmpFunc[int]), 1)

	assertEqual(t, Cmp[int](Only(1), Only(1), CmpFunc[int]), 0)
	assertEqual(t, Cmp[int](Only(4), Only(3), CmpFunc[int]), 1)
	assertEqual(t, Cmp[int](Only(2), Only(3), CmpFunc[int]), -1)

	assertEqual(t, Cmp[int](Only(1, 2, 3), Only(1, 2, 3), CmpFunc[int]), 0)
	assertEqual(t, Cmp[int](Only(1, 2), Only(1, 2, 3), CmpFunc[int]), -1)
	assertEqual(t, Cmp[int](Only(1, 2, 3), Only(1, 2), CmpFunc[int]), 1)

	assertEqual(t, Cmp[int](Only(1, 2, 4), Only(1, 2, 3), CmpFunc[int]), 1)
	assertEqual(t, Cmp[int](Only(1, 2, 2), Only(1, 2, 3), CmpFunc[int]), -1)

	assertEqual(t, Cmp[int](Only(4), Only(1, 2, 3), CmpFunc[int]), 1)
	assertEqual(t, Cmp[int](Only(1, 2, 3), Only(4), CmpFunc[int]), -1)
}
