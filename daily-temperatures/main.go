// https://leetcode.com/problems/daily-temperatures/

package main

import (
	"log"
)

func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {
		return []int{}
	}

	if len(temperatures) == 1 {
		return []int{0}
	}

	indexesStack := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(indexesStack) > 0 && temperatures[i] > temperatures[indexesStack[len(indexesStack)-1]] {
			indexOfLastItem := indexesStack[len(indexesStack)-1] // index of last item
			ans[indexOfLastItem] = i - indexOfLastItem
			indexesStack = indexesStack[:len(indexesStack)-1]
		}
		indexesStack = append(indexesStack, i)
	}

	return ans
}

func main() {
	t := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	// t1 := []int{1, 2, 3}
	log.Printf("%+v", dailyTemperatures(t))
	// dailyTemperatures(t1)
}

func naive(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				ans[i] = j - i
				break
			}
		}
	}
	log.Printf("Anws %+v", ans)
	return ans
}
