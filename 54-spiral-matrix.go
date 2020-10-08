// 简化为循环处理最外层


type Point struct {
	X, Y int
}

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	cols := len(matrix[0])

	min := func(x, y int) int {
		r := x
		if r > y {
			r = y
		}
		return r
	}

	walk := func(rows, cols int, c chan Point) {
		limit := (min(rows, cols) + 1) / 2

		for d := 0; d < limit; d++ {
			for x, y := d, d; x < cols-d-1; x++ {
				c <- Point{x, y}
			}

			for x, y := cols-d-1, d; y < rows-d-1; y++ {
				c <- Point{x, y}
			}

			for x, y := cols-d-1, rows-d-1; x > d; x-- {
				c <- Point{x, y}
			}

			for x, y := d, rows-d-1; y > d; y-- {
				c <- Point{x, y}
			}
		}
		
		if rows == cols {
			c <- Point{limit-1, limit-1}
		}
	}

	c := make(chan Point)
	go walk(rows, cols, c)
	res := make([]int, rows*cols)
	for i := 0; i < rows*cols; i++ {
		p := <-c
		v := matrix[p.Y][p.X]
		res[i] = v
	}

	return res
}
