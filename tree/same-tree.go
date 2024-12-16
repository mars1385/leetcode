package tree

func sameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	isValidLeft := sameTree(p.Left, q.Left)
	isValidRight := sameTree(p.Right, q.Right)

	if isValidLeft && isValidRight {
		return true
	}
	return false
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return sameTree(p, q)
}
