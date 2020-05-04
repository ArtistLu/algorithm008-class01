package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	fmt.Println(preOrderNR(a))
	fmt.Println(preOrder(a))
	fmt.Println(inOrder(a))
	// f.Left = h
	// f.Right = i
	// i.Right = j
	// obj := Constructor()
	// s := obj.serialize(a)
	// root := obj.deserialize(s)
	// printTree(root)
	// fmt.Println(s)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
func parseNode(node *TreeNode) string {
	if node == nil {
		return "null"
	}

	return strconv.Itoa(node.Val)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var inner func(root *TreeNode) string
	inner = func(root *TreeNode) string {
		if root == nil {
			return parseNode(root)
		}

		str := parseNode(root)
		str += "," + inner(root.Left)
		str += "," + inner(root.Right)
		return strings.Trim(str, ",")
	}

	return "[" + inner(root) + "]"
}

func paseString(s string) (int, bool) {
	x, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return x, true
}

//Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Trim(data, "[]")
	vals := strings.Split(s, ",")

	type queue struct {
		val  string
		next *queue
	}
	head := &queue{}
	cur := head
	for _, v := range vals {
		cur.next = &queue{val: v}
		cur = cur.next
	}

	q := head.next

	var inner func() *TreeNode
	inner = func() *TreeNode {
		if q == nil {
			return nil
		}

		val, ok := paseString(q.val)
		q = q.next

		if !ok {
			return nil
		}

		root := &TreeNode{Val: val}
		root.Left = inner()
		root.Right = inner()
		return root
	}

	return inner()
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

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func preOrderNR(root *TreeNode) []int {
	re := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return re
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		re = append(re, s.Val)

		if s.Right != nil {
			stack = append(stack, s.Right)
		}
		if s.Left != nil {
			stack = append(stack, s.Left)
		}
	}

	return re
}

func preOrder(root *TreeNode) []int {
	re := []int{}
	if root == nil {
		return re
	}

	re = append(re, root.Val)

	l := preOrder(root.Left)
	r := preOrder(root.Right)
	return append(append(re, l...), r...)
}

func inOrder(root *TreeNode) []int {
	re := []int{}
	if root == nil {
		return re
	}

	l := inOrder(root.Left)
	r := inOrder(root.Right)
	return append(append(l, root.Val), r...)
}

func inOrderNR(root *TreeNode) []int {
	re := []int{}
	if root == nil {
		return re
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {

	}
}
