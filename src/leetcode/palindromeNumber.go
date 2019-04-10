package leetcode

import "math"

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:
Input: 121
Output: true

Example 2:
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	} else {
		pnumber := 0
		var onumber int = x
		for x != 0 {
			if pnumber > math.MaxInt32/10 {
				return false
			} else {
				pnumber = pnumber*10 + x%10
				x = x / 10
			}
		}
		if pnumber == onumber {
			return true
		} else {
			return false
		}
	}
}
