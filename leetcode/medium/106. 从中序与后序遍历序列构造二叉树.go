package main
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	rootIndexInInorder := 0
	for ; rootIndexInInorder < len(inorder); rootIndexInInorder++ {
		if inorder[rootIndexInInorder] == rootVal {
			break
		}
	}
	inorderL := inorder[:rootIndexInInorder]
	inorderR := inorder[rootIndexInInorder+1:]
	postorderL := postorder[:len(inorderL)]
	postorderR := postorder[len(inorderL):len(postorder)-1]
	root.Left = buildTree(inorderL, postorderL)
	root.Right = buildTree(inorderR, postorderR)
	return root
}