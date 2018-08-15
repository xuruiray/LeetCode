package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	fmt.Println(len(nums))
}

func removeDuplicates(nums []int) int {
	count := 0
	datamap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, isDuplicate := datamap[nums[i]]; isDuplicate {
			if i == len(nums)-1 {
				nums = nums[:i]
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
			i = i - 1
		} else {
			datamap[nums[i]] = nums[i]
			count++
		}
	}
	return count
}
