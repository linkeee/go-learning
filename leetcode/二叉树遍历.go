package main

type TreeNode struct {
	Left, Right *TreeNode
	Val int
}

func Pre1(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func Pre2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	stack = append(stack, root) // 根结点先入栈
	for len(stack) > 0 { // 最终得把栈里的全部遍历完
		node := stack[len(stack)-1]; stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Right, node.Left)
	}
	return res
}

func Mid1(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

func Mid2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur) // 将各个根结点压栈，为了之后弹出并遍历右节点
			cur = cur.Left
		} else { // 没有左节点了，该出栈了
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func Post1(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func Post2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var pre *TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {  // 压入所有左节点
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1] // 没有左节点了，出栈
		stack = stack[:len(stack)-1] // 出栈之后，还要先判断右节点，如果有且未遍历过，则还要再次入栈
		if cur.Right == nil || cur.Right == pre { // 右节点为空或已遍历过，则可以遍历当前根结点
			res = append(res, cur.Val)
			pre = cur
			cur = nil // 不置为空的话，还会重复遍历左节点
		} else {
			stack = append(stack, cur)
			cur = cur.Right
		}
	}
	return res
}

func 层次遍历(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		each := []int{}
		for i := 0; i < l; i++ {
			node := queue[0]; queue = queue[1:]
			if node != nil {
				each = append(each, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		if len(each) > 0 {
			res = append(res, each)
		}
	}
	return res
}
