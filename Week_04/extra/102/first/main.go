package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}

	type item struct {
		level int
		node  *TreeNode
	}
	queue := []*item{&item{0, root}}

	a, l := []int{}, 0
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.level != l {
			ans = append(ans, a)
			a = []int{}
		}
		l = q.level
		a = append(a, q.node.Val)

		if q.node.Left != nil {
			queue = append(queue, &item{q.level + 1, q.node.Left})
		}

		if q.node.Right != nil {
			queue = append(queue, &item{q.level + 1, q.node.Right})
		}
	}

	ans = append(ans, a)

	return ans
}
