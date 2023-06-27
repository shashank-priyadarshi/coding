package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"abaa", "abc", "abca"}))
}

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
//
// words: an array of strings to search through
// commonChar: string containing the longest common prefix found
// Returns the longest common prefix as a string
func longestCommonPrefix(words []string) (commonChar string) {
	if len(words[0]) == 0 {
		return
	}
	firstChar := words[0][0]
	for key, word := range words {
		if len(word) == 0 {
			return ""
		}
		if word[0] != firstChar {
			return ""
		}
		if key == 0 {
			commonChar = word
			continue
		}
		tempCommonChar := commonChars(words[key-1], word)
		if len(tempCommonChar) == 0 {
			return
		}
		if len(tempCommonChar) < len(commonChar) {
			commonChar = tempCommonChar
		}
	}
	return
}

// commonChars returns the common prefix between two strings.
//
// word1: the first string to compare.
// word2: the second string to compare.
// commonChar: the common prefix between word1 and word2.
//             An empty string is returned if either word1 or word2 is empty.
//             If word1 and word2 are the same, word1 is returned.
//             If no common prefix is found, an empty string is returned.
func commonChars(word1, word2 string) (commonChar string) {
	if word1 == word2 {
		return word1
	}
	len1, len2, length := len(word1), len(word2), 0
	if len1 == 0 || len2 == 0 {
		return
	}
	if len1 <= len2 {
		length = len1
	} else {
		length = len2
	}
	for i := 0; i < length; i++ {
		if word1[i] != word2[i] {
			return word1[:i]
		}
	}
	return word1[:length]
}
