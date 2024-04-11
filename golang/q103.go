/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var stackLeft, stackRight []*TreeNode
	stackRight = append(stackRight, root)
	right := true
	nodesRem := 1
	nextLevel := 0

	var output [][]int
	var level []int
	for nodesRem > 0 || nextLevel > 0 {

		if nodesRem == 0 {
			right = !right
			nodesRem = nextLevel
			nextLevel = 0
			output = append(output, level)
			level = make([]int, 0)
		}
		if right {

			last := len(stackRight) - 1
			node := stackRight[last]
			stackRight = stackRight[:last]

			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel++
				stackLeft = append(stackLeft, node.Left)
			}
			if node.Right != nil {
				nextLevel++
				stackLeft = append(stackLeft, node.Right)
			}

		} else {

			last := len(stackLeft) - 1
			node := stackLeft[last]
			stackLeft = stackLeft[:last]

			level = append(level, node.Val)
			if node.Right != nil {
				nextLevel++
				stackRight = append(stackRight, node.Right)
			}
			if node.Left != nil {
				nextLevel++
				stackRight = append(stackRight, node.Left)
			}

		}

		nodesRem--
	}

	output = append(output, level)
	return output
}