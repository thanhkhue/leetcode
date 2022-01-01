package main

// https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/submissions/

func interchangeableRectangles(rectangles [][]int) int64 {
	ratio := make(map[float64]int)
	res := 0
	for i := 0; i < len(rectangles); i++ {
		res += ratio[float64(rectangles[i][0])/float64(rectangles[i][1])]
		ratio[float64(rectangles[i][0])/float64(rectangles[i][1])]++
	}

	return int64(res)

}
