/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var output [][]int
	if root == nil {
		return output
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		queue_len := len(queue)
		var vals []int
		for queue_len > 0 {

			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			vals = append(vals, node.Val)

			queue_len--
		}

		output = append([][]int{vals}, output...)
	}

	return output
}