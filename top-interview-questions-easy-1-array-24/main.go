package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	datamap := make(map[int]int)
	for _, v := range nums {
		if _, ok := datamap[v]; ok {
			return true
		}
		datamap[v] = 1
	}
	return false
}
