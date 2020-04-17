package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 7}
	l1 := generateList(list)
	l2 := generateList([]int{1, 4, 5, 7, 8})
	dumpList(l1)
	dumpList(l2)

	l3 := mergeTwoLists(l1, l2)
	dumpList(l3)
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	re := &ListNode{}

	if l1.Val < l2.Val {
		re.Next = l1
		l1 = l1.Next
	} else {
		re.Next = l2
		l2 = l2.Next
	}

	re.Next.Next = mergeTwoLists(l1, l2)

	return re.Next
}

func generateList(list []int) *ListNode {
	var head *ListNode
	if len(list) == 0 {
		return head
	}
	head = &ListNode{
		Val: list[0],
	}
	cur := head
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{
			Val: list[i],
		}
		cur = cur.Next
	}

	return head
}

func dumpList(list *ListNode) {
	for list != nil {
		fmt.Print("->", list.Val)
		list = list.Next
	}
	fmt.Println()
}
