// https://leetcode.com/problems/average-of-levels-in-binary-tree/

package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func treeHeight(root *TreeNode, h int) int {
	if root == nil {
		return h
	}
	return 1 + max(treeHeight(root.Left, h), treeHeight(root.Right, h))
}

func getArrayAverage(arr []int) float64 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	height := treeHeight(root, 0)
	log.Printf("height %d", height)
	res := make([][]int, height)
	level := 0

	returnAverageOfLevels(root, res, level)
	log.Printf("Array  %+v", res)
	average := make([]float64, len(res))
	for i := 0; i < len(res); i++ {
		average[i] = getArrayAverage(res[i])
	}

	log.Printf("average %+v", average)

	return nil
}

func returnAverageOfLevels(root *TreeNode, average [][]int, level int) {
	if root == nil {
		return
	}

	if average[level] == nil {
		average[level] = make([]int, 0)
	}

	average[level] = append(average[level], root.Val)
	level++
	returnAverageOfLevels(root.Left, average, level)
	returnAverageOfLevels(root.Right, average, level)
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	averageOfLevels(root)
}
