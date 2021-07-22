package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var indexOfMaxInSli func([]int) int
	indexOfMaxInSli = func(nums []int) (index int) {
		max := nums[0]
		for i := range nums {
			if nums[i] > max {
				index = i
				max = nums[i]
			}
		}
		return
	}
	maxIndex := indexOfMaxInSli(nums)
	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}
