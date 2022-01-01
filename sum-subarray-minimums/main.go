/*
https://leetcode.com/problems/sum-of-subarray-minimums/
*/

package main

import (
	"fmt"
	"log"
	"math"
)

/*
// previous_less[i] = j means A[j] is the previous less element of A[i].
// previous_less[i] = -1 means there is no previous less element of A[i].
vector<int> previous_less(A.size(), -1);
for(int i = 0; i < A.size(); i++){
  while(!in_stk.empty() && A[in_stk.top()] > A[i]){
    in_stk.pop();
  }
  previous_less[i] = in_stk.empty()? -1: in_stk.top();
  in_stk.push(i);
}
*/
const MOD = int(1e9 + 7)

func sumSubarrayMins(A []int) int {
	N := len(A)
	dp := make([]int, N, N)
	stack := make([]int, 0)
	var res = 0
	for i := 0; i != N; i++ {
		for len(stack) != 0 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1] // pop
		}
		if len(stack) != 0 {
			prev := stack[len(stack)-1]
			dp[i] = (A[i]*(i-prev) + dp[prev]) % MOD
		} else {
			dp[i] = (A[i] * (i + 1)) % MOD
		}
		stack = append(stack, i) // push
		res = (res + dp[i]) % MOD
	}
	return res
}

type ValIndex struct {
	Val      int
	Distance int
}

func nextLess(arr []int) map[int]int {
	nextLess := make(map[int]int)
	incStack := make([]ValIndex, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		for len(incStack) > 0 && incStack[len(incStack)-1].Val > arr[i] {
			incStack = incStack[:len(incStack)-1]
		}
		if len(incStack) == 0 {
			nextLess[i] = -1
		} else {
			// nextLess[i] = incStack[len(incStack)-1].Index
		}
		incStack = append(incStack, ValIndex{arr[i], i})
	}
	return nextLess
}

func leftDistancesOfNumberToMin(arr []int) []int {
	leftStack := make([]ValIndex, 0)
	left := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		distance := 1
		if len(leftStack) > 0 && leftStack[len(leftStack)-1].Val > arr[i] {
			distance += leftStack[len(leftStack)-1].Distance
			leftStack = leftStack[:len(leftStack)-1] // pop
		}
		leftStack = append(leftStack, ValIndex{
			Val:      arr[i],
			Distance: distance,
		}) // push
		left[i] = distance
	}
	return left
}

func rightDistancesOfNumberToMin(arr []int) []int {
	rightStack := make([]ValIndex, 0)
	right := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		distance := 1
		if len(rightStack) > 0 && rightStack[len(rightStack)-1].Val > arr[i] {
			distance += rightStack[len(rightStack)-1].Distance
			rightStack = rightStack[:len(rightStack)-1] // pop
		}
		rightStack = append(rightStack, ValIndex{
			Val:      arr[i],
			Distance: distance,
		})
		right[i] = distance
	}
	return right
}

func main() {
	arr := []int{3, 1, 2, 4}

	left := leftDistancesOfNumberToMin(arr)
	right := rightDistancesOfNumberToMin(arr)
	sum := 0
	mod := int(1e9 + 7)
	for i := 0; i < len(arr); i++ {
		sum += arr[i] * left[i] * right[i] % mod
	}
	log.Printf("Result %d", sum)
}

func sumSubarrayMinss(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sumMin := math.MaxInt
	markIndexes := make(map[string]bool)
	for i := 0; i <= len(arr); i++ {
		for j := 0; j <= len(arr); j++ {
			if i == j {
				continue
			}
			if i < j && !markIndexes[fmt.Sprintf("%d-%d", i, j)] {
				sumMin += minArray(arr[i:j])
				markIndexes[fmt.Sprintf("%d-%d", i, j)] = true
			} else if !markIndexes[fmt.Sprintf("%d-%d", j, i)] {
				markIndexes[fmt.Sprintf("%d-%d", j, i)] = true
				sumMin += minArray(arr[j:i])

			}
		}
	}

	return sumMin
}

func minArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func naiveSolution(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sumMin := math.MaxInt
	markIndexes := make(map[string]bool)
	for i := 0; i <= len(arr); i++ {
		for j := 0; j <= len(arr); j++ {
			if i == j {
				continue
			}
			if i < j && !markIndexes[fmt.Sprintf("%d-%d", i, j)] {
				sumMin += minArray(arr[i:j])
				markIndexes[fmt.Sprintf("%d-%d", i, j)] = true
			} else if !markIndexes[fmt.Sprintf("%d-%d", j, i)] {
				markIndexes[fmt.Sprintf("%d-%d", j, i)] = true
				sumMin += minArray(arr[j:i])

			}
		}
	}

	return sumMin
}
