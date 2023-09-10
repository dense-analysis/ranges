package ranges

import "testing"

func TestRetro(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(Retro(Only(1, 2, 3))), []int{3, 2, 1})

	r := Retro(Only("a", "b", "c"))
	s1 := r.Save()
	s2 := r.SaveB()

	s1.PopFront()
	s2.PopBack()

	assertHasFront(t, s1, "b")
	assertHasFront(t, s2, "c")
	assertHasBack(t, s2, "b")
	assertHasBack(t, r, "a")
}

func TestRetroR(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceR(RetroR(Only(1, 2, 3))), []int{3, 2, 1})

	r := RetroR(Only("a", "b", "c"))

	s1 := r.SaveB()
	s2 := r.SaveR()

	s1.PopFront()
	s2.PopBack()

	assertHasFront(t, s1, "b")
	assertHasFront(t, s2, "c")
	assertHasBack(t, s2, "b")
	assertHasBack(t, r, "a")
}
