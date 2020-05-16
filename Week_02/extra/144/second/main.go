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
	ans := []int{}

	var inner func(root *TreeNode)
	inner = func(root *TreeNode) {
		if root == nil {
			return
		}

		ans = append(ans, root.Val)
		inner(root.Left)
		inner(root.Right)
	}

	inner(root)
	return ans
}
