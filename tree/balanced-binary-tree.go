package tree

import "math"

func validateTree(root *TreeNode) (bool, float64) {

	if root == nil {
		return true, 0
	}

	isValidLeft, depthLeft := validateTree(root.Left)
	isValidRight, depthRight := validateTree(root.Right)

	if isValidLeft && isValidRight && math.Abs(float64(depthLeft)-float64(depthRight)) <= 1 {
		return true, math.Max(float64(depthLeft), float64(depthRight)) + 1
	}
	return false, 0
}
func isBalanced(root *TreeNode) bool {
	result, _ := validateTree(root)

	return result
}
