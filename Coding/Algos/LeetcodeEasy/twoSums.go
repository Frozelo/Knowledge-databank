// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// https://leetcode.com/problems/two-sum/description/


package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			fmt.Println(target - num)
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{4, 5}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
