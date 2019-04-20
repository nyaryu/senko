package leetcode

import (
	"strconv"
	"strings"
)

/*
No.38
he count-and-say sequence is the sequence of integers with the first five terms as following:
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.
Note: Each term of the sequence of integers will be represented as a string.


Example 1:
Input: 1
Output: "1"

Example 2:
Input: 4
Output: "1211"
*/
func CountAndSay(n int) string {
	result := "1"
	i := 1
	if n == 1 {
		return "1"
	}
	for i != n {
		tempArray := strings.Split(result, "")
		result = sequence(tempArray)
		i++
	}
	return result
}

func sequence(tempArray []string) string {
	value := tempArray[0]
	count := 0
	n := len(tempArray)
	current := 0
	result := ""
	for current != n {
		if tempArray[current] == value {
			count++
		} else {
			result += strconv.Itoa(count)
			result += value
			value = tempArray[current]
			count = 1
		}
		current++
	}
	result += strconv.Itoa(count)
	result += value
	return result
}
