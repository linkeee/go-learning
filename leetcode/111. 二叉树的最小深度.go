package main

func minDepth(root *TreeNode) int {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil {
			return 1 + getDepth(node.Right)
		}
		if node.Right == nil {
			return 1 + getDepth(node.Left)
		}
		lDep, rDep := getDepth(node.Left), getDepth(node.Right)
		if lDep < rDep {
			return 1 + lDep
		}
		return 1 + rDep
	}
	depth := getDepth(root)
	return depth
}
