// leetcode https://leetcode.com/problems/roman-to-integer/

package main

import "log"

var RomanIntegers = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var RomanPlacingBefore = map[byte]byte{
	'V': 'I',
	'X': 'I',
	'L': 'X',
	'C': 'X',
	'D': 'C',
	'M': 'C',
}

func romanToInt(s string) int {

	if len(s) == 0 {
		return 0
	}

	res := 0
	for i := 0; i < len(s); i++ {
		// log.Printf("+ %d (%s)", RomanIntegers[s[i]], string(s[i]))
		if v, ok := RomanPlacingBefore[s[i]]; ok && (i-1) >= 0 && s[i-1] == v {
			// log.Printf("i=%d, v(%s)=%d", i, string(v), RomanIntegers[v])
			// log.Printf("+ %d (%s)", RomanIntegers[s[i]], string(s[i]))
			res -= RomanIntegers[v]
			res -= RomanIntegers[v]
			res += RomanIntegers[s[i]]
			continue
		}
		res += RomanIntegers[s[i]]
	}

	log.Printf("Res %d", res)

	return 0
}

func main() {
	// romanToInt("III")
	// romanToInt("LVIII")
	romanToInt("IV")
	romanToInt("MCMXCIV")
}
