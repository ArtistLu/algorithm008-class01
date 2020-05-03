package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if root == p || root == q {
// 		return root
// 	}

// 	l := lowestCommonAncestor(root.Left, p, q)
// 	r := lowestCommonAncestor(root.Right, p, q)

// 	if l != nil && r != nil {
// 		return root
// 	}

// 	if l == nil {
// 		return r
// 	}

// 	return l
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	type node struct {
		parent *TreeNode
		val    *TreeNode
	}

	queue := []*node{&node{val: root, parent: root}}

	re := &TreeNode{}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.val == nil {
			continue
		}

		if n.val == p || n.val == q {
			re = n.parent
			break
		}

		queue = append(queue, &node{n.val, n.val.Left})
		queue = append(queue, &node{n.val, n.val.Right})
	}

	if re == nil {
		return nil
	}
	if (re.Left == p && re.Right == q) || (re.Left == q && re.Right == p) {
		return re
	}

	if re.Left == p || re.Left == q {
		return re.Left
	} else {
		return re.Right
	}
}
