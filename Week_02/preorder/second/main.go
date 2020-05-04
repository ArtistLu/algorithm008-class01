package main

func main() {

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

const (
	visit = "visit"
	print = "print"
)

func postorder(root *Node) []int {
	re := []int{}
	if root == nil {
		return re
	}

	type item struct {
		opt  string
		node *Node
	}

	stack := []*item{&item{visit, root}}
	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if s.opt == print {
			re = append(re, s.node.Val)
			continue
		}

		for i := len(s.node.Children) - 1; i >= 0; i-- {
			if s.node.Children[i] != nil {
				stack = append(stack, &item{visit, s.node.Children[i]})
			}
		}
		stack = append(stack, &item{print, s.node})
	}

	return re
}
