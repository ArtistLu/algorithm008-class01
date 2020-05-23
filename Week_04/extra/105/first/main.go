package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var p int

func buildTree(preorder []int, inorder []int) *TreeNode {
	p = 0
	return helper(preorder, inorder, 0, len(inorder)-1)
}

func helper(preorder, inorder []int, s, e int) *TreeNode {
	if s > e {
		return nil
	}

	root := &TreeNode{Val: preorder[p]}
	p++

	var i int
	for i = s; i <= e; i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = helper(preorder, inorder, s, i-1)
	root.Right = helper(preorder, inorder, i+1, e)
	return root
}
