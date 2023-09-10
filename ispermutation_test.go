package ranges

import "testing"

func TestIsPermutation(t *testing.T) {
	t.Parallel()

	if !IsPermutation[int](Only[int](), Only[int]()) {
		t.Error("Two empty ranges were not pemutations of each other")
	}

	if IsPermutation[int](Only(1), Only[int]()) {
		t.Error("A non-empty range was considered a permutation of an empty one")
	}

	if IsPermutation[int](Only[int](), Only(1)) {
		t.Error("An empty range was considered a permutation of a non-empty one")
	}

	if !IsPermutation[int](Only(1, 2, 3), Only(1, 2, 3)) {
		t.Error("(1, 2, 3) was not considered a permutation of (1, 2, 3)")
	}

	if !IsPermutation[int](Only(3, 2, 1), Only(1, 2, 3)) {
		t.Error("(3, 2, 1) was not considered a permutation of (1, 2, 3)")
	}

	if !IsPermutation[int](Only(1, 2, 3), Only(3, 2, 1)) {
		t.Error("(1, 2, 3) was not considered a permutation of (3, 2, 1)")
	}

	if !IsPermutation[int](Only(1, 2, 3), Only(3, 2, 1)) {
		t.Error("(1, 2, 3) was not considered a permutation of (3, 2, 1)")
	}

	if IsPermutation[int](Only(1, 2, 2, 3, 3, 3), Only(3, 2, 1)) {
		t.Error("(1, 2, 2, 3, 3, 3) was considered a permutation of (3, 2, 1)")
	}

	if IsPermutation[int](Only(3, 2, 1), Only(1, 2, 2, 3, 3, 3)) {
		t.Error("(3, 2, 1) was considered a permutation of (1, 2, 2, 3, 3, 3)")
	}

	if !IsPermutation[int](Only(3, 2, 1, 3, 3, 2), Only(1, 2, 2, 3, 3, 3)) {
		t.Error("(3, 2, 1, 3, 3, 2) was not considered a permutation of (1, 2, 2, 3, 3, 3)")
	}
}

func TestIsPermutationNoAlloc(t *testing.T) {
	t.Parallel()

	if !IsPermutationNoAlloc(F(B(Only[int]())), F(B(Only[int]()))) {
		t.Error("Two empty ranges were not pemutations of each other")
	}

	if IsPermutationNoAlloc(F(B(Only(1))), F(B(Only[int]()))) {
		t.Error("A non-empty range was considered a permutation of an empty one")
	}

	if IsPermutationNoAlloc(F(B(Only[int]())), F(B(Only(1)))) {
		t.Error("An empty range was considered a permutation of a non-empty one")
	}

	if !IsPermutationNoAlloc(F(B(Only(1, 2, 3))), F(B(Only(1, 2, 3)))) {
		t.Error("(1, 2, 3) was not considered a permutation of (1, 2, 3)")
	}

	if !IsPermutationNoAlloc(F(B(Only(3, 2, 1))), F(B(Only(1, 2, 3)))) {
		t.Error("(3, 2, 1) was not considered a permutation of (1, 2, 3)")
	}

	if !IsPermutationNoAlloc(F(B(Only(1, 2, 3))), F(B(Only(3, 2, 1)))) {
		t.Error("(1, 2, 3) was not considered a permutation of (3, 2, 1)")
	}

	if !IsPermutationNoAlloc(F(B(Only(1, 2, 3))), F(B(Only(3, 2, 1)))) {
		t.Error("(1, 2, 3) was not considered a permutation of (3, 2, 1)")
	}

	if IsPermutationNoAlloc(F(B(Only(1, 2, 2, 3, 3, 3))), F(B(Only(3, 2, 1)))) {
		t.Error("(1, 2, 2, 3, 3, 3) was considered a permutation of (3, 2, 1)")
	}

	if IsPermutationNoAlloc(F(B(Only(3, 2, 1))), F(B(Only(1, 2, 2, 3, 3, 3)))) {
		t.Error("(3, 2, 1) was considered a permutation of (1, 2, 2, 3, 3, 3)")
	}

	if !IsPermutationNoAlloc(F(B(Only(3, 2, 1, 3, 3, 2))), F(B(Only(1, 2, 2, 3, 3, 3)))) {
		t.Error("(3, 2, 1, 3, 3, 2) was not considered a permutation of (1, 2, 2, 3, 3, 3)")
	}
}
