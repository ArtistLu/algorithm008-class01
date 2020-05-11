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

	a, ans := []int{}, [][]int{}
	queue1, queue2 := []*TreeNode{root}, []*TreeNode{}
	for len(queue1) > 0 && len(queue2) > 0 {
		q := queue1[0]
		queue1 = queue1[1:]
		a = append(a, q.Val)
		if q.Left != nil {
			queue2 = append(queue2, q.Left)
		}
		if q.Right != nil {
			queue2 = append(queue2, q.Right)
		}

		if len(queue1) == 0 {
			ans = append(ans, append([]int{}, a...))
			a, queue1, queue2 = []int{}, queue2, queue1
		}
	}

	return ans
}
