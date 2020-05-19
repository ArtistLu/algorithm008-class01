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
	fmt.Println(inorderTraversal(a))
}

// 二叉树前序遍历
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	print = "print"
	visit = "visit"
)

func inorderTraversal(root *TreeNode) []int {
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

		if p.node.Right != nil {
			stack = append(stack, bullet{visit, p.node.Right})
		}
		stack = append(stack, bullet{print, p.node})
		if p.node.Left != nil {
			stack = append(stack, bullet{visit, p.node.Left})
		}
	}

	return ans
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	r, c := len(matrix), len(matrix[0])
	l, r := 0, r*c-1

	for l <= r {
		mid := l + (r-l)>>1
		m := matrix[mid/c][mid%c]
		if m == target {
			return true
		} else if m > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
