package main

import "log"

func main() {
	// arr := []int{2, 2, 1, 5, 1, 5, 5, 2, 3, 1, 1, 5, 3, 2, 3, 3, 3, 1, 3, 3, 4, 3, 1, 3, 1, 4, 5, 5, 2, 2, 1, 3, 5, 2, 2, 4, 3, 2, 5, 3, 1, 1, 3, 3, 2, 5, 2, 1, 2, 4, 3, 4, 4, 3, 2, 4, 4, 1, 3, 3, 3, 5, 5, 5, 4, 1, 1, 2, 3, 3, 2, 5, 3, 4, 5, 3, 1, 2, 5, 4, 5, 2, 3, 3, 1, 5, 2, 4, 2, 4, 4, 3, 1, 3}
	arr := []int{1, 1, 2, 3, 2, 2}
	log.Printf("Number of items %d", len(arr))
	log.Printf("Number of pairs %d", numIdenticalPairs(arr))
}

func numIdenticalPairs(nums []int) int {
	res := 0
	indexes := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		res += indexes[nums[i]]
		indexes[nums[i]]++
	}
	return res
}

func numIdenticalPairsV3(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	indexes := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		indexes[nums[i]]++
	}

	res := 0
	for i := 0; i < len(indexes); i++ {
		res += combination(indexes[i])
	}

	return res
}

func combination(i int) int {
	sum := 0
	for j := 0; j < i; j++ {
		sum += j
	}
	return sum
}

func numIdenticalPairsV2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	indexed := make([]int, 100)
	for i := 0; i < len(nums); i++ {
		indexed[nums[i]]++
	}

	numPairs := 0
	for i := 0; i < len(indexed); i++ {
		if indexed[i] < 2 {
			log.Printf("index %d encounter %d time(s)", i, indexed[i])
			continue
		}
		numPairs += indexed[i] - 1
	}

	return numPairs
}

func combinationPairs(i int) int {
	return factorial(i) / (factorial(2) * factorial(i-2))

}

func factorial(num int) int {
	total := 1
	for i := 1; i <= num; i++ {
		total = total * i
	}
	return total
}
