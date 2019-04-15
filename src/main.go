package main

import (
	"fmt"
	"leetcode"
)

func main() {
	x := "(])"
	var res = leetcode.ValidParentheses(x)
	fmt.Println(res)
	y := []int{3, 2, 2, 3}
	val := 3
	var res1 = leetcode.RemoveElement(y, val)
	fmt.Println(res1)
}
