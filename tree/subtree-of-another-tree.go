package tree

func subTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	isValidLeft := subTree(p.Left, q.Left)
	isValidRight := subTree(p.Right, q.Right)
	if isValidLeft && isValidRight {
		return true
	}
	return false
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil {
		return false
	}
	return subTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
