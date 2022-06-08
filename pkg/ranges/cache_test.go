package ranges

import "testing"

func TestCache(t *testing.T) {
	t.Parallel()

	assertEqual(t, Slice(Cache(I(Null[int]()))), []int{})
	assertEqual(t, Slice(Cache(I(Only(1, 2, 3)))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := Cache(
		Map(I(Only(5, 6, 7)), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
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

	assertEqual(t, SliceF(CacheF(Null[int]())), []int{})
	assertEqual(t, SliceF(CacheF(Only(1, 2, 3))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := CacheF(
		MapF(Only(5, 6, 7), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
	)

	r.Front()
	r.Front()
	r.PopFront()

	r2Save := r.Save()

	r.Front()
	r.Front()
	r.PopFront()

	r2Save.Front()
	r2Save.PopFront()

	assertEqual(t, storedValues, []int{5, 6, 7, 7})
}
