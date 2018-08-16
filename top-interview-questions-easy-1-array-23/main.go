package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	rotate(nums, k)
	fmt.Println(nums, k)
}

func rotate(nums []int, k int) {
	if k < 0 {
		k = -k
	}
	if k != 0 {
		k = k % len(nums)
		k = len(nums) - k
		slice := append(nums[k:], nums[:k]...)
		copy(nums, slice)
	}
}
