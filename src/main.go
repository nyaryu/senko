package main

import (
	"fmt"
	"leetcode"
)

func main() {
	var nums = []int{5, 7, 11, 20}
	target := 16
	var result []int
	result = leetcode.TwoSum(nums, target)
	fmt.Println(result)
}
