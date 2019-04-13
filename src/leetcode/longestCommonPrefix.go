package leetcode

/*
No.14
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:
All given inputs are in lowercase letters a-z.
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	} else {
		for i := 1; i < len(strs[0])+1; i++ {
			for j := 1; j < len(strs); j++ {
				if len(strs[j]) == 0 {
					return ""
				} else if strs[0][:1] != strs[j][:1] {
					return ""
				} else if i > len(strs[j]) || strs[0][:i] != strs[j][:i] {
					return strs[0][:i-1]
				}
			}
		}
	}
	return strs[0]
}
