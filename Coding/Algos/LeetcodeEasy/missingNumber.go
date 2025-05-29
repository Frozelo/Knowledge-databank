// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

// Input: nums = [3,0,1]
// Output: 2
// Explanation:
// n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

// Approach 1 With sorting O(n*log n)

func missingNumber(nums []int) int {
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if nums[i] != i {
            return i
        }
    }

    return len(nums)
}

// Approach 2 O(n)
func missingNumber(nums []int) int {
    n:=len(nums)
    sum:=(1+n)*n/2

    for _, num:=range nums{
        sum-=num
    }

    return sum
}
