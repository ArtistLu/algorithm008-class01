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

	printTree(a)

	fmt.Println(postorder(a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("empty tree!")
	}

	type node struct {
		level int
		node  *TreeNode
	}

	tree := [][]*TreeNode{}
	t := []*TreeNode{}
	queue := []*node{&node{0, root}}
	level := 0
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if level != q.level {
			tree = append(tree, t)
			t = []*TreeNode{}
		}
		if q.node != nil {
			level = q.level
			t = append(t, q.node)
			queue = append(queue, &node{q.level + 1, q.node.Left})
			queue = append(queue, &node{q.level + 1, q.node.Right})
		} else {
			t = append(t, nil)
		}
	}
	max := 0
	for i := 0; i < len(tree); i++ {
		if len(tree[i]) > max {
			max = len(tree[i])
		}
	}

	for i := 0; i < len(tree); i++ {
		for j := 0; j < max-len(tree[i]); j++ {
			fmt.Print(" ")
		}
		for _, n := range tree[i] {
			if n == nil {
				fmt.Print(" ")
			} else {
				fmt.Print(n.Val)
			}

			fmt.Print("  ")
		}
		fmt.Println()
	}
}

const (
	print = "print"
	visit = "visit"
)

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	re := []int{}
	type node struct {
		opt      string
		treeNode *TreeNode
	}

	stack := []*node{&node{visit, root}}

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if s.opt == print {
			re = append(re, s.treeNode.Val)
			continue
		}

		stack = append(stack, &node{print, s.treeNode})
		if s.treeNode.Right != nil {
			stack = append(stack, &node{visit, s.treeNode.Right})
		}
		if s.treeNode.Left != nil {
			stack = append(stack, &node{visit, s.treeNode.Left})
		}
	}

	return re
}
