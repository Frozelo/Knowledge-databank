// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// https://leetcode.com/problems/product-of-array-except-self/description/

func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    prod := 1

    for i := 0; i < n; i++ {
        res[i] = prod
        prod*= nums[i]
    }

    prod = 1
    for i := n-1; i >=0; i-- {
        res[i] *= prod
        prod*= nums[i]
    }

    return res
}
