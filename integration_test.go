package ranges

// This file tests how different algorithms integrate together.
import "testing"

func TestChunksAsSlices(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4, 5}
	chunks := Slice(Map(Chunks[int](SliceRange(input), 3), Slice[int]))

	assertEqual(t, chunks, [][]int{{1, 2, 3}, {4, 5}})
}

func TestFlattenChunks(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4, 5}
	chunkOutput := [][]int{}
	chunkRange := Chunks[int](SliceRange(input), 3)

	for !chunkRange.Empty() {
		chunkOutput = append(chunkOutput, Slice(chunkRange.Front()))
		chunkRange.PopFront()
	}

	assertEqual(t, chunkOutput, [][]int{{1, 2, 3}, {4, 5}})

	combinedOutput := Slice(Flatten(Chunks[int](SliceRange(input), 3)))

	assertEqual(t, combinedOutput, input)
}

func TestTakeCycleRepeat(t *testing.T) {
	t.Parallel()

	result := SliceF(
		TakeF(
			Cycle(
				ChainF(
					F(B(Only(1, 2, 3))),
					F(Repeat(4)),
				),
			),
			5,
		),
	)

	assertEqual(t, result, []int{1, 2, 3, 4, 4})
}
