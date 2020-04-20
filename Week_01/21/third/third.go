package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	l1 := generateList(list)
	l2 := generateList([]int{1, 3, 4})
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
//递归
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

func mergeTwoListsx(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	re := &ListNode{}
	cur := re
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

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
