package main

import "fmt"

// import "Week_02/binarytree"
func main() {
	a := &TreeNode{10, nil, nil}
	b := &TreeNode{7, nil, nil}
	c := &TreeNode{12, nil, nil}
	d := &TreeNode{4, nil, nil}
	e := &TreeNode{8, nil, nil}
	f := &TreeNode{11, nil, nil}
	g := &TreeNode{14, nil, nil}
	// h := &TreeNode{8, nil, nil}
	// i := &TreeNode{9, nil, nil}
	// j := &TreeNode{10, nil, nil}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	fmt.Println(levelOrder(a))
}

// 二叉树层序遍历
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue1, queue2 := []*TreeNode{root}, []*TreeNode{}
	a, ans := []int{}, [][]int{}
	for len(queue1) > 0 || len(queue2) > 0 {
		q := queue1[0]
		queue1 = queue1[1:]

		a = append(a, q.Val)

		if q.Left != nil {
			queue2 = append(queue2, q.Left)
		}
		if q.Right != nil {
			queue2 = append(queue2, q.Right)
		}

		if len(queue1) == 0 {
			queue1, queue2 = queue2, queue1
			ans, a = append(ans, a), []int{}
		}
	}

	return ans
}
