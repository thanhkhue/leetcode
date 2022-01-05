// https://leetcode.com/problems/integer-to-roman/

package main

import "log"

var Romains = []string{"I", "V", "X", "L", "C", "D", "M"}
var Integers = []int{1, 5, 10, 50, 100, 500, 1000}
var less = []int{0, 1, 1, 10, 10, 100, 100}
var IntegerToRomains = map[int]string{
	100: "C",
	10:  "X",
	1:   "I",
}

func intToRoman(num int) string {
	if num == 0 {
		return ""
	}

	orig := num
	res := ""
	for i := len(Integers) - 1; i >= 0; i-- {
		for num > 0 && num >= Integers[i] {
			res += Romains[i]
			num -= Integers[i]
			continue
		}
		for (num > 0) && num >= Integers[i]-less[i] && num < Integers[i] {
			res += IntegerToRomains[less[i]] + Romains[i]
			num = num - (Integers[i] - less[i])
			continue
		}
	}

	log.Printf("int %d -> res %s", orig, res)
	return res
}

func main() {
	nums := []int{3, 58, 1994}
	for i := 0; i < len(nums); i++ {
		intToRoman(nums[i])
	}
}
