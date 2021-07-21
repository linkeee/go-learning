package main

import "fmt"

func main() {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return
	}
	println(upstairs(n))
}

func upstairs(n int) int {
	if n < 3 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
