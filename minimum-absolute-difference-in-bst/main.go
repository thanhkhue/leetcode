// https://leetcode.com/problems/minimum-absolute-difference-in-bst/

package main

import (
	"log"
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
				236
		104				701
	null	227
*/

func main() {
	// [236,104,701,null,227,null,911]
	root := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val: 104,
			Right: &TreeNode{
				Val: 227,
			},
		},
		Right: &TreeNode{
			Val: 701,
			Right: &TreeNode{
				Val: 911,
			},
		},
	}
	getMinimumDifference(root)
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	arr := make([]int, 0)
	minDiff(root, &arr)
	log.Printf("Table %+v", arr)
	sort.Ints(arr)

	m := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if absInt(arr[i], arr[i-1]) < m {
			m = absInt(arr[i], arr[i-1])
		}
	}

	log.Printf("min abs %d", m)
	return m
}

func minDiff(root *TreeNode, table *[]int) {
	if root == nil {
		return
	}

	*table = append(*table, root.Val)

	if root.Left != nil {
		minDiff(root.Left, table)
	}
	if root.Right != nil {
		minDiff(root.Right, table)
	}

}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func absInt(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
