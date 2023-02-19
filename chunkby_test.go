package ranges

import "testing"

func TestChunkBy(t *testing.T) {
	t.Parallel()

	result := Slice(
		Map(
			ChunkBy(
				I(F(Only(1, 2, 3, 3, 4, 5, 6, 6, 7, 8))),
				func(a, b int) bool { return a == b },
			),
			Slice[int],
		),
	)

	assertEqual(t, result, [][]int{{1}, {2}, {3, 3}, {4}, {5}, {6, 6}, {7}, {8}})

	result2 := Slice(
		Map(
			ChunkBy(
				I(F(Only(1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9))),
				func(a, b int) bool { return a < b },
			),
			Slice[int],
		),
	)

	assertEqual(t, result2, [][]int{{1, 2, 3}, {3, 4, 5, 6}, {6, 7, 8}, {8, 9}, {9}})
}

func TestChunkByRepeatedFront(t *testing.T) {
	t.Parallel()

	result := ChunkBy(
		I(F(Only(1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 8, 9, 9))),
		func(a, b int) bool { return a < b },
	)

	assertHasFront(t, result.Front(), 1)
	assertHasFront(t, result.Front(), 1)

	result.PopFront()

	assertHasFront(t, result.Front(), 3)
	assertHasFront(t, result.Front(), 3)

	result.PopFront()

	assertHasFront(t, result.Front(), 6)
	assertHasFront(t, result.Front(), 6)

	result.PopFront()
	result.PopFront()

	assertHasFront(t, result.Front(), 9)

	result.PopFront()

	assertEmpty(t, result)
}

type chunkByValueTestStruct struct {
	x int
}

func TestChunkByValue(t *testing.T) {
	t.Parallel()

	result := Slice(
		Map(
			ChunkByValue(
				I(F(SliceRange([]chunkByValueTestStruct{{1}, {2}, {3}, {3}, {4}, {5}, {6}, {6}, {7}, {8}}))),
				func(a chunkByValueTestStruct) int { return a.x },
			),
			Slice[chunkByValueTestStruct],
		),
	)

	assertEqual(t, result, [][]chunkByValueTestStruct{
		{{1}},
		{{2}},
		{{3}, {3}},
		{{4}},
		{{5}},
		{{6}, {6}},
		{{7}},
		{{8}},
	})
}
