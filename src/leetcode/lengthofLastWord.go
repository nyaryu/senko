package leetcode

import (
	"fmt"
	"strings"
)

/*
No.58
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
If the last word does not exist, return 0.
Note: A word is defined as a character sequence consists of non-space characters only.

Example:
Input: "Hello World"
Output: 5
*/

func LengthofLastWord(s string) int {
	if len(s) == 0 || s == " " {
		return 0
	}
	result := 0
	s = strings.TrimRight(s, " ")
	temp := strings.Split(s, "")
	fmt.Println(temp)
	i := 0
	for ; i < len(s); i++ {
		if temp[i] == " " {
			result = 0
		} else {
			result++

		}
	}
	return result
}
