package main

import (
	"fmt"
	"leetcode"
)

func main() {
	x := "(])"
	var res = leetcode.ValidParentheses(x)
	fmt.Println(res)
}
