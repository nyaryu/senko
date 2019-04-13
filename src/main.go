package main

import (
	"fmt"
	"leetcode"
)

func main() {
	var x = []string{"c", "c", "c"}
	var res = leetcode.LongestCommonPrefix(x)
	fmt.Println(res)
}
