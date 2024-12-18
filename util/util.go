package until

type Comparable interface {
	int | bool
}

func MatrixMatch[C Comparable](a [][]C, b [][]C) bool {
	if len(a) != len(b) {
		return false
	}
	for i, row := range a {
		if len(row) != len(a[i]) {
			return false
		}
		for j, val := range row {
			if val != b[i][j] {
				return false
			}
		}
	}
	return true
}

// Compare two sequences. Number of elements and element-wise
// equality gives a match.
// Use slices.Equal() instead, it should do the same.
func SequenceMatch[C Comparable](a []C, b []C) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
