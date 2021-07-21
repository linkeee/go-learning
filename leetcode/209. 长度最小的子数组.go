package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := len(nums) + 1 // 不加一无法分辨出是不是整个数组满足情况还是都不满足情况
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target { // 满足这个条件，就要先计算长度
			if right - left + 1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if minLen == len(nums) + 1 {
		return 0
	}
	return minLen
}

func main() {
	nums := []int{2,3,1,2,4,3}
	target := 7
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
