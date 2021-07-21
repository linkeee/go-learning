package main

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var path []string // 每次加单条路径，遇到叶子结点就组成一条路径，然后回溯
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {   // 因为要路径，所以要根结点开始，叶子结点终止，就是前序遍历
		path = append(path, strconv.Itoa(node.Val))  // 前序遍历，先根结点
		// 然后左节点、右节点
		// 但是由于不确定node是不是叶子，所以要先处理node是叶子的情况
		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(path, "->"))
			return
		}
		// 真正开始处理左右节点
		if node.Left != nil {
			traversal(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			traversal(node.Right)
			path = path[:len(path)-1]
		}
	}
	traversal(root)
	return result
}