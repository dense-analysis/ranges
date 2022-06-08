package ranges

// Max returns the maximum value among provided arguments.
func Max[T Ordered](v1, v2 T, values ...T) T {
	maxVal := v1

	if v2 > v1 {
		maxVal = v2
	}

	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

// Min returns the minimum value among provided arguments.
func Min[T Ordered](v1, v2 T, values ...T) T {
	minVal := v1

	if v2 < v1 {
		minVal = v2
	}

	for _, v := range values {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}
