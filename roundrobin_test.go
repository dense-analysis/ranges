package ranges

import "testing"

func TestRoundRobin(t *testing.T) {
	t.Parallel()

	output := Slice(
		RoundRobin[int](
			Only(1, 2, 3),
			Only(4, 5),
			Only(7, 8, 9),
		),
	)

	assertEqual(t, output, []int{1, 4, 7, 2, 5, 8, 3, 9})

	output2 := Slice(RoundRobin[int](Only(1, 2, 3), Only(10, 20, 30, 40)))

	assertEqual(t, output2, []int{1, 10, 2, 20, 3, 30, 40})

	// Ensure you can call PopFront() and Front() before Empty()
	r := RoundRobin[int](Only(1, 3), Only(2))

	r.PopFront()

	val1 := r.Front()
	r.PopFront()
	val2 := r.Front()
	r.PopFront()

	assertEqual(t, []int{val1, val2}, []int{2, 3})
}

func TestRoundRobinF(t *testing.T) {
	t.Parallel()

	output := SliceF(
		RoundRobinF(
			Only(1, 2, 3),
			Only(4, 5),
			Only(7, 8, 9),
		),
	)

	assertEqual(t, output, []int{1, 4, 7, 2, 5, 8, 3, 9})

	r := RoundRobinF(Only(1, 3), Only(2))
	r.PopFront()

	r2 := r.Save()

	r.PopFront()

	assertHasFront(t, r2, 2)
	assertHasFront(t, r, 3)
}
