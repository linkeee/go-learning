package main

func maxDepth(root *TreeNode) int {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := getDepth(node.Left), getDepth(node.Right)
		if left > right {
			return 1 + left
		}
		return 1 + right
	}
	depth := getDepth(root)
	return depth
}
