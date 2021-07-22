package main

func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := 0
	minLen := len(nums) + 1
	for ; end < len(nums); end++ {
		sum += nums[end] // 外层for循环，每次都会加
		for sum >= target {
			if end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	if minLen > len(nums) {
		return 0
	}
	return minLen
}
