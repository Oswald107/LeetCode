/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	return recur3(1, n)
}

func recur3(start int, end int) []*TreeNode {
	if start == end {
		return []*TreeNode{&TreeNode{start, nil, nil}}
	}

	if start > end {
		return []*TreeNode{nil}
	}

	var output []*TreeNode
	i := start
	for i < end+1 {
		fmt.Println(i)
		left := recur3(start, i-1)
		right := recur3(i+1, end)
		for _, l := range left {
			for _, r := range right {
				output = append(output, &TreeNode{i, l, r})
			}
		}

		i++
	}

	return output
}
