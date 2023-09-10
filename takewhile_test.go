package ranges

import "testing"

func TestTakeWhile(t *testing.T) {
	t.Parallel()

	sr := Only(1, 3, 2, 4)
	r := TakeWhile[int](sr, func(x int) bool { return x%2 == 1 })

	assertHasFront(t, sr, 1)

	r.PopFront()
	assertEqual(t, Slice(r), []int{3})

	r2 := TakeWhile[int](Only(1, 3, 2, 4), func(x int) bool { return x%2 == 0 })

	assertEmpty(t, r2)
}

func TestTakeWhileF(t *testing.T) {
	t.Parallel()

	r := TakeWhileF(Only(1, 2, 3), func(x int) bool { return x%2 == 1 })

	assertNotEmpty(t, r)

	rSave := r.Save()
	r.PopFront()

	assertEmpty(t, r)
	assertNotEmpty(t, rSave)
}
