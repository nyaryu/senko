package leetcode

/*
No.1
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func TwoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for a := 0; a < len(nums)-1; a++ {
		var first = nums[a]
		for b := a + 1; b < len(nums); b++ {
			var second = nums[b]
			if first+second == target {
				res[0] = a
				res[1] = b
			}
		}
	}
	return res
}
