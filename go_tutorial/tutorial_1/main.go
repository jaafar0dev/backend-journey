package main

import (
	"fmt"
)

// Two sum problem, return indices of the two numbers such that they add up to a specific target.

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	fmt.Println("Hello, World!")
}
