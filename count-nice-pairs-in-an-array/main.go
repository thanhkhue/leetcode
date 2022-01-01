package main

import "log"

// https://leetcode.com/problems/count-nice-pairs-in-an-array/

func main() {
	// log.Printf("Reverse of %d is %d", 9283549, reverse(9283549))
	arr := []int{13, 10, 35, 24, 76}
	log.Printf("Number of nice pairs is %d", countNicePairs(arr))
}

func countNicePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	markedSum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rev := reverse(nums[i])
		v := markedSum[nums[i]-rev]
		res = (res + v) % 1000000007
		markedSum[nums[i]-rev] = v + 1
	}

	return res
}

func reverse(num int) int {
	r := 0
	for num > 0 {
		r = r*10 + (num % 10)
		num = num / 10
	}
	return r
}
