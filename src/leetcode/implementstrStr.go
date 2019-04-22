package leetcode

/*
No.28
Implement strStr().
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

func ImplementstrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	i := 0
	for i < len(haystack) {
		if haystack[i] == needle[0] {
			j := 1
			for ; j < len(needle); j++ {
				if i+j >= len(haystack) || haystack[i+j] != needle[j] {
					break
				}
			}

			if j == len(needle) {
				return i
			}
		}
		i++
	}
	return -1
}
