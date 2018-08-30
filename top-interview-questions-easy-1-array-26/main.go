package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	datamap := make(map[int]int)
	resultList := make([]int, 0, len(nums1))

	for _, v := range nums1 {
		datamap[v]++
	}

	for _, v2 := range nums2 {
		if v, ok := datamap[v2]; ok && v > 0 {
			resultList = append(resultList, v2)
			datamap[v2]--
		}
	}

	return resultList
}
