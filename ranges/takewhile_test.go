package ranges

import "testing"

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	sr := F(Only(1, 3, 2, 4))
	r := TakeWhile[int](sr, func(x int) bool { return x%2 == 1 })

	assertHasFrontF(t, sr, 1)

	r.PopFront()
	assertEqual(t, Slice(r), []int{3})

	r2 := TakeWhile[int](Only(1, 3, 2, 4), func(x int) bool { return x%2 == 0 })

	assertEmpty(t, r2)
}

func TestTakeWhileF(t *testing.T) {
	t.Parallel()

	r := TakeWhileF(F(Only(1, 2, 3)), func(x int) bool { return x%2 == 1 })

	assertNotEmptyF(t, r)

	rSave := r.Save()
	r.PopFront()

	assertEmptyF(t, r)
	assertNotEmptyF(t, rSave)
}
