package leetcode

/*
No.35
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You may assume no duplicates in the array.

Example 1:
Input: [1,3,5,6], 5
Output: 2

Example 2:
Input: [1,3,5,6], 2
Output: 1

Example 3:
Input: [1,3,5,6], 7
Output: 4

Example 4:
Input: [1,3,5,6], 0
Output: 0
*/
func SearchInsertPosition(nums []int, target int) int {
	begin := 0
	end := len(nums)
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			begin = mid + 1
		} else if nums[mid] > target {
			end = mid
		}
	}
	return begin
}

func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else {
			if nums[i] > target {
				return 0
			} else if nums[i] < target {
				if i == len(nums)-1 || nums[i+1] > target {
					return i + 1
				}
			}

		}
	}
	return 0
}
