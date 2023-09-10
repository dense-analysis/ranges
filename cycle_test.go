package ranges

import "testing"

func TestCycle(t *testing.T) {
	t.Parallel()

	r := Cycle(F(B(Only(1, 2, 3))))

	r1 := r.Front()
	r.PopFront()
	r.Front()
	r2 := r.Front()
	r.PopFront()
	r.Front()
	r3 := r.Front()
	r.PopFront()
	r4 := r.Front()

	assertEqual(t, []int{r1, r2, r3, r4}, []int{1, 2, 3, 1})
}
