package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var s *ListNode

	for head != nil {
		head.Next, s, head = s, head, head.Next
	}

	return s
}
