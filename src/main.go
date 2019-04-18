package main

import (
	"fmt"
	"leetcode"
)

func main() {
	y := []int{0, 1, 2, 3}
	target := 4
	var res1 = leetcode.SearchInsert(y, target)
	fmt.Println(res1)
}
