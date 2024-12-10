package tree

func goNext(root *TreeNode) int {

	resL := 1
	resR := 1
	if root.Left != nil {
		resL += goNext(root.Left)
	}

	if root.Right != nil {
		resR += goNext(root.Right)
	}

	if resL > resR {
		return resL
	}
	return resR
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return goNext(root)
}
