package main

import "fmt"

// 94. 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

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

	fmt.Println(postorderTraversal(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	print = "print"
	visit = "visit"
)

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}

	type bullet struct {
		opt  string
		node *TreeNode
	}

	stack := []bullet{bullet{visit, root}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if p.opt == print {
			ans = append(ans, p.node.Val)
			continue
		}

		stack = append(stack, bullet{print, p.node})

		if p.node.Right != nil {
			stack = append(stack, bullet{visit, p.node.Right})
		}

		if p.node.Left != nil {
			stack = append(stack, bullet{visit, p.node.Left})
		}
	}

	return ans
}
