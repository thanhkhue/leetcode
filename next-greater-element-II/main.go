/*
https://leetcode.com/problems/next-greater-element-ii/
*/
package main

type ValueIndex struct {
	Val   int
	Index int
}

func nextGreaterElements(nums []int) []int {
	stacks := make([]ValueIndex, 0)

	origLength := len(nums)
	nums = append(nums, nums...)
	nextGreaterElements := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nextGreaterElements[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		for len(stacks) > 0 && nums[i] > stacks[len(stacks)-1].Val {
			if nextGreaterElements[stacks[len(stacks)-1].Index] == -1 {
				nextGreaterElements[stacks[len(stacks)-1].Index] = nums[i]
			}
			stacks = stacks[:len(stacks)-1]
		}
		stacks = append(stacks, ValueIndex{
			Val:   nums[i],
			Index: i,
		})
	}

	return nextGreaterElements[:origLength]
}

func main() {
	nums := []int{5, 4, 3, 2, 1} // result: [-1,5,5,5,5]
	nextGreaterElements(nums)
	nums1 := []int{1, 2, 1}
	nextGreaterElements(nums1)
}
