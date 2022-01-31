package main

import "log"

func searchInsert(nums []int, target int) int {
	idx := -1
	mid := -1
	start, end := 0, len(nums)-1
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			log.Printf("Result %d", mid)
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	log.Printf("Mid %d", mid)

	if nums[mid] > target {
		for i := mid; i >= 0; i-- {
			if nums[i] <= target {
				log.Printf("1/Res of i %d", i+1)
				return i + 1
			}
		}
	}

	if nums[mid] < target {
		for i := mid; i <= len(nums)-1; i++ {
			if nums[i] <= target {
				log.Printf("2/Res of i %d", i+1)
				return i + 1
			}
		}
	}

	log.Println("Reach here")
	if idx < 0 {
		return 0
	}

	return idx
}

func main() {
	arr := []int{1, 3, 5, 6}
	target := 0
	searchInsert(arr, target)
}
