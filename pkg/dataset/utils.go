package dataset

// MakeRange creates a slice of integers from zero to max-1.
func MakeRange(max int) []int {
	if max < 0 {
		return []int{}
	}

	r := make([]int, max)
	for i := range r {
		r[i] = i
	}

	return r
}
