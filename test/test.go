package main

import "fmt"

func main() {
	test1()
}

func test1() int {
	res := [][]int {
		{1,2,3,11},
		{4,5,6,12},
		{7,8,9,13},
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			fmt.Print(res[i][j], &res[i][j], "\t")
		}
		fmt.Println()
	}
	return 1
}

