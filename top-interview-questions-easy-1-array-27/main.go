package main

import "fmt"

func main() {
	nums := []int{9, 9}
	fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
	index := len(digits) - 1

	for index >= 0 {
		value, CY := singlePlusOne(digits[index])
		digits[index] = value
		if !CY {
			break
		}
		index--
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func singlePlusOne(num int) (int, bool) {
	if num == 9 {
		return 0, true
	}
	return num + 1, false
}
