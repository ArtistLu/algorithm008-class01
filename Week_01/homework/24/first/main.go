package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	re := head.Next
	head.Next = swapPairs(head.Next.Next)
	re.Next = head
	return re
}
