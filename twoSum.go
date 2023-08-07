package main

import "fmt"

func main() {
	// 1. TWO SUM. Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	array1 := []int{1, 2, 4, 3, 7}

	fmt.Println(twoSum(array1, 9))
}
func twoSum(nums []int, target int) []int {
	new := []int{}
	for i := 0; i < len(nums); i++ {
		for k := len(nums) - 1; k > -1; k-- {
			if i != k {
				if nums[i]+nums[k] == target {

					new = append(new, i)
				}
			}
		}
	}
	return new
}
