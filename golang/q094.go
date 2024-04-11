/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	output := inorderTraversal(root.Left)
	output = append(output, root.Val)
	return append(output, inorderTraversal(root.Right)...)
}