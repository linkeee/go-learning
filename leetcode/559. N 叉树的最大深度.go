package main

type Node struct {
	Val int
	Children []*Node
}

func maxDepth1(root *Node) int {
	var getDepth func(node *Node) int
	getDepth = func(node *Node) int {
		if node == nil {
			return 0
		}
		childDepth := 0
		for _, child := range node.Children {
			temp := getDepth(child)
			if temp > childDepth {
				childDepth = temp
			}
		}
		return 1 + childDepth
	}
	depth := getDepth(root)
	return depth
}
