package main

type TreeNode struct {
	Left, Right *TreeNode
	Val int
}

func countNodes递归(root *TreeNode) int {
	var getNum func(node *TreeNode) int
	getNum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + getNum(node.Left) + getNum(node.Right)
	}
	return getNum(root)
}

func countNodes层次(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	result := 0
	for len(queue) > 0 {
		size := len(queue)
		result += size
		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
	}
	return result
}
