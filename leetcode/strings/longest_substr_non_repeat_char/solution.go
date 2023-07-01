package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("ciridmtxsgaryv"))
}

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
//
// s: The input string.
// longestSubstr: The length of the longest substring without repeating characters.
// Returns the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) (longestSubstr int) {
	length := len(s)
	if length <= 1 {
		return length
	}
	longestSubstr = 1
	substrMap := make(map[string]map[string]int)
	for _, charRune := range s {
		charStr := string(charRune)
		for k, v := range substrMap {
			delete(substrMap, k)
			if _, ok := v[charStr]; !ok {
				k += charStr
				v[charStr] = 1
				substrMap[k] = v
				length = len(k)
				if length > longestSubstr {
					longestSubstr = length
				}
			}
		}
		substrMap[charStr] = make(map[string]int)
		substrMap[charStr][charStr] = 1
	}
	return
}
