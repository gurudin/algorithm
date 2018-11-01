package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	var tmpStr string

	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[i+1:len(s)], s[i]) > -1 {
			tmpStr = s[i : strings.IndexByte(s[i+1:len(s)], s[i])+i+1]
		} else {
			tmpStr = s[i:len(s)]
		}

		for j := 1; j < len(tmpStr); j++ {
			if strings.IndexByte(tmpStr[j+1:len(tmpStr)], tmpStr[j]) > -1 {
				tmpStr = tmpStr[0 : strings.IndexByte(tmpStr[j+1:len(tmpStr)], tmpStr[j])+j+1]
			}
		}

		if maxLen < len(tmpStr) {
			maxLen = len(tmpStr)
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcbcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
