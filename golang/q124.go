/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	return max(recur(root))
}

func recur(node *TreeNode) (int, int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	}

	if node.Left == nil {
		e2, n2 := recur(node.Right)
		return max(e2, 0) + node.Val, max(n2, node.Val, e2)
	}

	if node.Right == nil {
		e1, n1 := recur(node.Left)
		return max(e1, 0) + node.Val, max(n1, node.Val, e1)
	}

	e1, n1 := recur(node.Left)
	e2, n2 := recur(node.Right)

	edge := max(e1, e2, 0) + node.Val
	not_edge := max(n1, n2, e1+node.Val+e2, node.Val, e1, e2)

	return edge, not_edge
}