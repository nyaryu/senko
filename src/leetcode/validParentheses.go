package leetcode

/*
No.20
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false

Example 4:
Input: "([)]"
Output: false

Example 5:
Input: "{[]}"
Output: true
*/

func ValidParentheses(s string) bool {
	maxsize := len(s)
	if len(s) == 0 {
		return true
	} else if len(s)%2 == 1 {
		return false
	} else if len(s) >= 2 {
		if s[0:1] == ")" || s[0:1] == "]" || s[0:1] == "}" {
			return false
		} else {
			var list = initSql(maxsize)
			for i := 1; i < maxsize; i++ {
				if s[i-1:i] == "(" || s[i-1:i] == "[" || s[i-1:i] == "{" {
					list.data[list.length] = s[i-1 : i]
					list.length++
				}
				if list.length > 0 {
					if (list.data[list.length-1] == "(" && s[i:i+1] == ")") || (list.data[list.length-1] == "[" && s[i:i+1] == "]") || (list.data[list.length-1] == "{" && s[i:i+1] == "}") {
						list.length--
					}
				}
			}
			if list.length == 0 {
				return true
			}
		}
	}
	return false
}

type SqList struct {
	maxsize int
	length  int
	data    []string
}

func initSql(maxsize int) *SqList {
	return &SqList{maxsize: maxsize, data: make([]string, maxsize)}
}
