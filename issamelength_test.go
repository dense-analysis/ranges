package ranges

import "testing"

func TestIsSameLength(t *testing.T) {
	t.Parallel()

	if !IsSameLength(Only[int](), Only[float64]()) {
		t.Error("Two empty ranges were not the same length")
	}

	if !IsSameLength(Only(1), Only(4.0)) {
		t.Error("(1) was not considered the same length as (4.0)")
	}

	if !IsSameLength(Only(4, 5, 6), Only(7.0, 8.0, 9.0)) {
		t.Error("(4, 5, 6) was not considered the same length as (7.0, 8.0, 9.0)")
	}
}
