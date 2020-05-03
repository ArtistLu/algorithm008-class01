package main

func main() {
	// 前序遍历 preorder = [3,9,20,15,7]
	// 中序遍历 inorder = [9,3,15,20,7]
	// 返回如下的二叉树：

	// 	3
	//  / \
	// 9  20
	// 	/  \
	//  15   7

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
var p int

func buildTree(preorder []int, inorder []int) *TreeNode {
	p = 0
	return helper(preorder, inorder, 0, len(preorder)-1)
}

func helper(preorder, inorder []int, start, end int) *TreeNode {
	if start < end {
		return nil
	}

	root := &TreeNode{Val: preorder[p]}
	p++
	var i int
	for i = start; i <= end; i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = helper(preorder, inorder, start, i-1)
	root.Right = helper(preorder, inorder, i+1, end)
	return root
}
