package ranges

import "testing"

func TestRetro(t *testing.T) {
	assertEqual(t, SliceB(Retro(Only(1, 2, 3))), []int{3, 2, 1})

	r := Retro(Only("a", "b", "c"))
	s1 := r.Save()
	s2 := r.SaveB()

	s1.PopFront()
	s2.PopBack()

	assertHasFrontF(t, s1, "b")
	assertHasFrontB(t, s2, "c")
	assertHasBack(t, s2, "b")
	assertHasBack(t, r, "a")
}
