// https://leetcode.com/problems/rotate-image/

package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	/*
		1	2	3
		4	5	6
		7	8	9
	*/

	// step 1: Transpose the array
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// ====>
	/*
		1	4	7
		2	5	8
		3	6	9
	*/

	// step 2: swap column first vs last
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}

	// ====>
	/*
		7	4	1
		8	5	2
		9	6	3
	*/

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

}

func main() {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
}
