package main

import "fmt"

func main() {
	nums := []int{3,2,1,0,4}
	ret := canJump(nums)
	fmt.Println(ret)
}

func canJump(nums []int) bool {
	max_i := 0
	for i, v := range nums {
		if i > max_i {
			return false
		}
		if i + v > max_i {
			max_i = i + v
		}
	}
	return true
}
