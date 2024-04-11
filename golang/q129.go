/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var output int
	recur(root, 0, &output)
	return output
}

func recur(root *TreeNode, cur int, output *int) {
	if root == nil {
		return
	}

	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*output += cur
	}
	recur(root.Left, cur, output)
	recur(root.Right, cur, output)
}