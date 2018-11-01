package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	var tmpStr string

	for i := 0; i < len(s); i++ {
		var tmpLen int

		if strings.IndexByte(s[i+1:len(s)], s[i]) > -1 {
			tmpLen = strings.IndexByte(s[i+1:len(s)], s[i])
			tmpStr = s[i : i+tmpLen+1]
		} else {
			tmpLen = len(s)
			tmpStr = s[i:tmpLen]
		}

		fmt.Println(tmpStr)
	}

	return maxLen
}

func main() {
	// lengthOfLongestSubstring("abcbcabcbb")
	lengthOfLongestSubstring("pwwkew")
}
