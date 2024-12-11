package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "math"

func goNext(root *TreeNode, result *int) int {

	if root == nil {
		return 0
	}

	left := goNext(root.Left, result)
	right := goNext(root.Right, result)

	tempValue := int(math.Max(float64(right), float64(left))) + 1
	if *result < left+right {
		*result = left + right
	}

	return tempValue

}
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}
	result := 0
	goNext(root, &result)
	return result
}
