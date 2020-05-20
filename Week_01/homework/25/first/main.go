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
func reverseKGroup(head *ListNode, k int) *ListNode {
	i, cur := 0, &ListNode{Next: head}
	for ; i < k && cur != nil; i, cur = i+1, cur.Next {
	}
	if i < k || cur == nil {
		return head
	}
	next := cur.Next
	end := head
	s := &ListNode{}
	for s != cur {
		s, head, head.Next = head, head.Next, s
	}
	end.Next = reverseKGroup(next, k)
	return s
}
