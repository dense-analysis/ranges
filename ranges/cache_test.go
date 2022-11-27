package ranges

import "testing"

func TestCache(t *testing.T) {
	t.Parallel()

	assertEqual(t, Slice(Cache(I(F(Null[int]())))), []int{})
	assertEqual(t, Slice(Cache(I(F(Only(1, 2, 3))))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := Cache(
		Map(I(F(Only(5, 6, 7))), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
	)

	r.Front()
	r.Front()
	r.PopFront()
	r.Front()
	r.Front()
	r.PopFront()
	r.Front()
	r.Front()
	r.PopFront()

	assertEqual(t, storedValues, []int{5, 6, 7})
}

func TestCacheF(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(CacheF(F(Null[int]()))), []int{})
	assertEqual(t, SliceF(CacheF(F(Only(1, 2, 3)))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := CacheF(
		MapF(F(Only(5, 6, 7)), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
	)

	r.Front()
	r.Front()
	r.PopFront()

	r2Save := r.Save()

	r.Front()
	r.Front()
	r.PopFront()

	r2Save.Front()
	r2Save.Front()
	r2Save.PopFront()

	r.Front()
	r.Front()
	r.PopFront()

	assertEqual(t, storedValues, []int{5, 6, 6, 7})
}
