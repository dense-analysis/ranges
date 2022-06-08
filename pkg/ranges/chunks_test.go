package ranges

import "testing"

func TestChunks(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunks := Chunks[int](SliceRange(input), 2)
	allValues := make([][]int, 0, len(input)/2)

	for !chunks.Empty() {
		c := chunks.Front()
		cValues := make([]int, 0, 2)

		for !c.Empty() {
			cValues = append(cValues, c.Front())
			c.PopFront()
		}

		allValues = append(allValues, cValues)
		chunks.PopFront()
	}

	assertEqual(t, allValues, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}})
}

func TestChunksNotEmptyAfterChunkExhausted(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4}
	chunks := Chunks[int](SliceRange(input), 2)

	firstChunk := chunks.Front()
	firstChunk.PopFront()
	firstChunk.PopFront()

	assertNotEmpty(t, chunks)
}

func TestChunksModifiedByReference(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4}
	chunks := Chunks[int](SliceRange(input), 2)

	sliceCopy := Slice(chunks.Front())

	// The chunk should be correct
	assertEqual(t, sliceCopy, []int{1, 2})
	// The current chunk should be empty when we ask for it again.
	assertEmpty(t, chunks.Front())

	chunks.PopFront()

	secondChunk := chunks.Front()
	chunks.Front().PopFront()

	// The previously loaded chunk reference should be advanced when we pop a new value.
	assertHasFront(t, secondChunk, 4)
}

func TestChunkF(t *testing.T) {
	t.Parallel()

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunks := ChunksF(SliceRange(input), 2)
	allValues := make([][]int, 0, len(input)/2)
	const saveAt = 2
	i := 0

	var savedChunk ForwardRange[ForwardRange[int]]

	for !chunks.Empty() {
		if i == saveAt {
			savedChunk = chunks.Save()
		}

		c := chunks.Front()
		cValues := make([]int, 0, 2)

		for !c.Empty() {
			cValues = append(cValues, c.Front())
			c.PopFront()
		}

		allValues = append(allValues, cValues)
		chunks.PopFront()

		i++
	}

	assertEqual(t, allValues, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}})

	savedResult := SliceF(MapF(savedChunk, SliceF[int]))

	assertEqual(t, savedResult, [][]int{{5, 6}, {7, 8}, {9}})
}
