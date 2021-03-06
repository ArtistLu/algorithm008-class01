package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.Val == p.Val || root.Val == q.Val {
// 		return root
// 	}

// 	left := lowestCommonAncestor(root.Left, p, q)
// 	right := lowestCommonAncestor(root.Right, p, q)

// 	if left != nil && right != nil {
// 		return root
// 	}

// 	if left == nil {
// 		return right
// 	}

// 	return left
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}

	if l != nil {
		return r
	}

	return r
}
