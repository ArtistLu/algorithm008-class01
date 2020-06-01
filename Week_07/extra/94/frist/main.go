package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	visit = "vist"
	print = "print"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	type s struct {
		opt  string
		node *TreeNode
	}

	stack := []s{s{visit, root}}
	ans := []int{}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur.opt == print {
			ans = append(ans, cur.node.Val)
			continue
		}

		if cur.node.Right != nil {
			stack = append(stack, s{visit, cur.node.Right})
		}
		stack = append(stack, s{print, cur.node})
		if cur.node.Left != nil {
			stack = append(stack, s{visit, cur.node.Left})
		}

	}

	return ans
}
