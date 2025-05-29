// Single Number
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

// Example 1:
// Input: nums = [2,2,1]
// Output: 1

// Example 2:
// Input: nums = [4,1,2,1,2]
// Output: 4

// Example 3:
// Input: nums = [1]
// Output: 1

// https://leetcode.com/problems/single-number/

// O(n), O(n)
func singleNumber(nums []int) int {
    m := make(map[int]int, len(nums))

    for _, num := range nums {
        if v, ok := m[num]; ok {
            m[num] = v+1
        } else {
            m[num] = 1
        }
    }

    for i, cnt := range m {
        if cnt == 1 {
            return i
        }
    }

    return 0
}

// O(n), O(1)
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums{
        result ^= num
    }

    return result
}


