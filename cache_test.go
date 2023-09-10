package ranges

import (
	"testing"
)

func TestCache(t *testing.T) {
	t.Parallel()

	assertEqual(t, Slice(Cache(I(F(B(Null[int]()))))), []int{})
	assertEqual(t, Slice(Cache(I(F(B(Only(1, 2, 3)))))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := Cache(
		Map(I(F(B(Only(5, 6, 7)))), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
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

	assertEqual(t, SliceF(CacheF(F(B(Null[int]())))), []int{})
	assertEqual(t, SliceF(CacheF(F(B(Only(1, 2, 3))))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := CacheF(
		MapF(F(B(Only(5, 6, 7))), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
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

func TestCacheB(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(CacheB(B(Null[int]()))), []int{})
	assertEqual(t, SliceB(CacheB(B(Only(1, 2, 3)))), []int{1, 2, 3})

	storedValues := make([]int, 0)

	r := CacheB(
		MapB(B(Only(5, 6, 7)), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
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

func TestCacheBBack(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(Retro(CacheB(B(Null[int]())))), []int{})
	assertEqual(t, SliceB(Retro(CacheB(B(Only(1, 2, 3))))), []int{3, 2, 1})

	storedValues := make([]int, 0)

	r := CacheB(
		MapB(B(Only(5, 6, 7)), func(elem int) int { storedValues = append(storedValues, elem); return elem }),
	)

	r.Back()
	r.Back()
	r.PopBack()

	r2Save := r.SaveB()

	r.Back()
	r.Back()
	r.PopBack()

	r2Save.Back()
	r2Save.Back()
	r2Save.PopBack()

	r.Back()
	r.Back()
	r.PopBack()

	assertEqual(t, storedValues, []int{7, 6, 6, 5})
}
