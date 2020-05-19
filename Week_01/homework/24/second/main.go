package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	re := &ListNode{Next: head}
	for cur := re; cur.Next != nil && cur.Next.Next != nil; {
		s, e := cur.Next, cur.Next.Next
		cur.Next = e
		s.Next = e.Next
		e.Next = s
		cur = s
	}

	return re.Next
}
