package main

import "fmt"

func main() {
	root := &TreeNode{}

	re := maxPathSum(root)
	fmt.Println(re)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxPathSum(root *TreeNode) int {
	helper(root)
	return ans
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	right := helper(root.Right)
	ans = max(ans, left+right+root.Val)
	return max(left, right) + root.Val
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
