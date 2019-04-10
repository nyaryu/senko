package main

import (
	"fmt"
	"leetcode"
)

func main() {
	var nums = []int{5, 7, 11, 20}
	target := 16
	var res1 []int
	res1 = leetcode.TwoSum(nums, target)
	fmt.Println(res1)
	var x = 1534236469
	var res2 int
	res2 = leetcode.ReverseInteger(x)
	fmt.Println(res2)
	var y = 11234543211
	var res3 bool
	res3 = leetcode.PalindromeNumber(y)
	fmt.Println(res3)
}
