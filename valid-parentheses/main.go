// https://leetcode.com/problems/valid-parentheses/

package main

import "log"

func main() {
	log.Printf("Is valid %t", isValid("{]"))
}

func isValid(s string) bool {

	if len(s) == 0 {
		log.Println("1")
		return true
	}

	if len(s)%2 != 0 {
		log.Println("2")
		return false
	}

	stacks := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "]":
			log.Println("should go here")
			if len(stacks) == 0 || stacks[len(stacks)-1] != "[" {
				return false
			}
			stacks = stacks[:len(stacks)-1]
		case ")":
			if len(stacks) == 0 || stacks[len(stacks)-1] != "(" {
				return false
			}
			stacks = stacks[:len(stacks)-1]
		case "}":
			if len(stacks) == 0 || stacks[len(stacks)-1] != "{" {
				return false
			}
			stacks = stacks[:len(stacks)-1]
		default:
			stacks = append(stacks, string(s[i]))
		}
	}

	log.Printf("Stacks %+v", stacks)

	return len(stacks) == 0
}
