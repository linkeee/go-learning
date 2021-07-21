package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		// 只要有一个节点为空，就不能再继续往下比了，要返回了
		if left == right && left == nil {
			return true
		} else if left == nil || right == nil {
			return false
		}
		// 现在是两个节点不为空，比较数值后，还要继续往下递归
		if left.Val != right.Val {
			return false
		}
		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		return outside && inside
	}

	lRoot, rRoot := root.Left, root.Right
	return compare(lRoot, rRoot)
}
