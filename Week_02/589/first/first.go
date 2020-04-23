package main

import "fmt"

func main() {
	root := &Node{
		Val:      1,
		Children: []*Node{},
	}

	second := []int{3, 2, 4}
	for _, v := range second {
		root.Children = append(root.Children, &Node{
			Val:      v,
			Children: []*Node{},
		})
	}

	third := []int{5, 6}
	for _, v := range third {
		root.Children[0].Children = append(root.Children[0].Children, &Node{
			Val:      v,
			Children: []*Node{},
		})
	}

	fmt.Println(preorder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	re := []int{}
	if root == nil {
		return re
	}

	re = append(re, root.Val)

	for _, c := range root.Children {
		re = append(re, preorder(c)...)
	}

	return re
}
