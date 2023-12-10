package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("(\"abcdcaf\"): %v\n", getIndexOfFirstNonRepeatingCharacter("abcdcaf"))
}
func getIndexOfFirstNonRepeatingCharacter(str string) int {
	for i := 0; i < len(str); i++ {
		var isNonRepeating bool = true

		if strings.Contains(str[i+1:], str[i:i+1]) {
			isNonRepeating = false
		}
		// for j := i + 1; j < len(str); j++ {
		// 	if str[i] == str[j] {
		// 		isNonRepeating = false
		// 	}
		// }
		if isNonRepeating {
			return i
		}
	}

	return -1
}
