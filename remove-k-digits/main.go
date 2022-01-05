// https://leetcode.com/problems/remove-k-digits/

package main

import (
	"log"
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) == 0 || num == "0" || len(num) == k {
		return "0"
	}

	stacks := make([]string, 0)
	for i := 0; i < len(num); i++ {
		for len(stacks) > 0 && string(num[i]) < stacks[len(stacks)-1] && k > 0 {
			stacks = stacks[:len(stacks)-1]
			k--
		}
		stacks = append(stacks, string(num[i]))
	}

	for k > 0 {
		stacks = stacks[:len(stacks)-1]
		k--
	}

	if len(stacks) == 0 || len(stacks) == 1 && stacks[0] == "0" {
		return "0"
	}

	res := strings.Join(stacks, "")
	log.Printf("res %s", res)
	for strings.HasPrefix(res, "0") {
		res = strings.TrimPrefix(res, "0")
	}

	if res == "" {
		return "0"
	}

	log.Printf("k %d", k)
	log.Printf("res %s", res)
	log.Printf("Stacks %+v", stacks)
	return res
}

func main() {
	str := "100"
	removeKdigits(str, 1)
}
