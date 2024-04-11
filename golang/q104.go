
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	i := 0
	return recur(preorder, inorder, &i, 0, len(preorder))
}

func recur(preorder []int, inorder []int, i *int, start int, end int) *TreeNode {
	if start >= end {
		return nil
	}

	val := preorder[*i]
	j := index(inorder, val, start, end)
	*i++

	node := TreeNode{
		val,
		recur(preorder, inorder, i, start, j),
		recur(preorder, inorder, i, j+1, end)}

	return &node
}

func index(lst []int, val, start, end int) int {
	for i := start; i < end; i++ {
		if val == lst[i] {
			return i
		}
	}
	return -1
}