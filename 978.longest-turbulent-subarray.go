func maxTurbulenceSize(A []int) int {
	max := 1
	for i, j := 0, 2; j < len(A); j++ {
		diff := A[j-1] - A[j-2]
		if !(diff > 0 && A[j] < A[j-1]) && !(diff < 0 && A[j] > A[j-1]) {
			i = j - 1
		}
		
		if diff != 0 && j-i+1 > max {
			max = j - i + 1
		}
	}
	return max
}
