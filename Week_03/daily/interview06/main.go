package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//数组左侧入
func reversePrint1(head *ListNode) []int {
	re := []int{}

	for head != nil {
		re = append([]int{head.Val}, re...)
		head = head.Next
	}

	return re
}

//栈
func reversePrint2(head *ListNode) []int {
	stack := []int{}
	print := []int{}
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	for len(stack) > 0 {
		print = append(print, stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]
	}

	return print
}

//递归
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return append(reversePrint(head.Next), head.Val)
}
