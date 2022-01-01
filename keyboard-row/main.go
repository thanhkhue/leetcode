/*
https://leetcode.com/problems/keyboard-row/
*/

package main

import (
	"log"
)

/*
first row: qwertyuiop
second row: asdfghjkl
third row: zxcvbnm
*/
var wordToRow = map[byte]int{}

func init() {
	words := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for index, word := range words {
		for i := 0; i < len(word); i++ {
			wordToRow[word[i]] = index + 1
			wordToRow[word[i]-32] = index + 1
		}
	}
}

func findWords(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	results := make([]string, 0)
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if isLetterInSameRow(word) {
			results = append(results, word)
		}
	}

	return results
}

func isLetterInSameRow(word string) bool {
	row := wordToRow[word[0]]
	if row < 1 || row > 3 {
		return false
	}

	for i := 1; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			if row != wordToRow[word[i]+32] {
				return false
			}
			continue
		} else if word[i] >= 'a' && word[i] <= 'z' {
			if row != wordToRow[word[i]] {
				return false
			}
			continue
		}
		return false
	}
	return true
}

func main() {

	words := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for index, word := range words {
		for i := 0; i < len(word); i++ {
			wordToRow[word[i]] = index + 1
			wordToRow[word[i]-32] = index + 1
		}
	}

	// log.Printf("a=%d, A+32=%d, Z=%d, a-A=%d", int('a'), int('A'+32), int('Z'), int('a'-'A'))
	words1 := []string{"Hello", "Alaska", "Dad", "Peace"}
	log.Printf("Results %+v", findWords(words1))

	words2 := []string{"omk"}
	log.Printf("Results %+v", findWords(words2))

	words3 := []string{"adsdf", "sfd"}
	log.Printf("Results %+v", findWords(words3))
}
