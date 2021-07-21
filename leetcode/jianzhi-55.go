package main

import "fmt"

func main() {
	var input int
	for {
		if _, err := fmt.Scanf("%d", &input); err != nil {
			return
		}
		res := fibonacci3(input)
		fmt.Println(res)
	}
}

// 递归
func fibonacci1(input int) int {
	if input < 2 {
		return input
	}
	return fibonacci1(input-1) + fibonacci1(input-2)
}

func fibonacci2(input int) int {
	if input < 2 {
		return input
	}
	dp := []int{0, 1}
	for i := 2; i <= input; i++ {
		dp = append(dp, dp[i-1] + dp[i-2])
	}
	return dp[input]
}

func fibonacci3(input int) int {
	if input < 2 {
		return input
	}
	dp := [3]int{0, 1}
	for i := 2; i <= input; i++ {
		dp[2] = dp[0] + dp[1]
		dp[0], dp[1] = dp[1], dp[2]
	}
	return dp[2]
}

func fibonacci4(input int) int {
	if input < 2 {
		return input
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= input; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
