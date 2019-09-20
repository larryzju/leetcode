func rotate(matrix [][]int)  {
	n := len(matrix)
	conv := func(i, j int) (int, int) {
		return n-j-1, i
	}
    	
	for r := 0; r < n-1; r++ {
		for c := r; c < n - r - 1; c++ {
			x0, y0 := c, r
			x1, y1 := conv(x0, y0)
			x2, y2 := conv(x1, y1)
			x3, y3 := conv(x2, y2)
			matrix[x0][y0], matrix[x1][y1], matrix[x2][y2], matrix[x3][y3] = matrix[x1][y1], matrix[x2][y2], matrix[x3][y3], matrix[x0][y0]
		}
	}
}
