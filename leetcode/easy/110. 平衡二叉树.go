package main

type TreeNode struct {
	Left, Right *TreeNode
	Val int
}

// 求深度可以从上到下去查 所以需要前序遍历（中左右），而高度只能从下到上去查，所以只能后序遍历（左右中）
func isBalanced(root *TreeNode) bool {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lDep, rDep := getDepth(node.Left), getDepth(node.Right)
		if lDep == -1 || rDep == -1 || lDep - rDep > 1 || rDep - lDep > 1 {
			return -1
		}
		if lDep > rDep {
			return 1 + lDep
		}
		return 1 + rDep
	}
	if getDepth(root) == -1 {
		return false
	}
	return true
}
