package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1,2}
	s := []int{1,2,3}
	children := findContentChildren(g, s)
	fmt.Println(children)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && s[i] >= g[index] {
			result++
			index++  // 满足情况，饼干消耗一个
		}
	}
	return result
}