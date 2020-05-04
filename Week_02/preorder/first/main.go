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
	fmt.Println(preorderTraversal(a))
}

// 二叉树前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	print = "print"
	visit = "visit"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	re := []int{}

	type node struct {
		opt  string
		node *TreeNode
	}

	stack := []*node{&node{visit, root}}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if s.opt == print {
			re = append(re, s.node.Val)
			continue
		}
		if s.node.Right != nil {
			stack = append(stack, &node{visit, s.node.Right})
		}
		if s.node.Left != nil {
			stack = append(stack, &node{visit, s.node.Left})
		}
		stack = append(stack, &node{print, s.node})
	}

	return re
}
