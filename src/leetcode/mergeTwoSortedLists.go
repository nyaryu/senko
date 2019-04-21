package leetcode

/*
No.21
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	temp := &result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp.Next = l2
			break
		}
		if l2 == nil {
			temp.Next = l1
			break
		}
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
