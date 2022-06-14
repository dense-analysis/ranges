package ranges

import "testing"

func TestIsSameLength(t *testing.T) {
	t.Parallel()

	if !IsSameLength(I(Only[int]()), I(Only[float64]())) {
		t.Error("Two empty ranges were not the same length")
	}

	if !IsSameLength(I(Only(1)), I(Only(4.0))) {
		t.Error("(1) was not considered the same length as (4.0)")
	}

	if !IsSameLength(I(Only(4, 5, 6)), I(Only(7.0, 8.0, 9.0))) {
		t.Error("(4, 5, 6) was not considered the same length as (7.0, 8.0, 9.0)")
	}
}
