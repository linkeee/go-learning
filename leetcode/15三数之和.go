package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1,0,1,0}
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	if l < 3 || nums[0] > 0 {
		return nil
	}
	var res [][]int

	for i := 0; i < l-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, l-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}
