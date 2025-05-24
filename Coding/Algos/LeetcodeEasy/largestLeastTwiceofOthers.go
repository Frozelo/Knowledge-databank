// You are given an integer array nums where the largest integer is unique.
// Determine whether the largest element in the array is at least twice as much as every other number in the array. 
// If it is, return the index of the largest element, or return -1 otherwise.


// O(n)
func dominantIndex(nums []int) int {
    if len(nums) < 0 {
        return -1
    }

    maxIdx, maxVal := 0, nums[0]
    secondMax := -1 << 31

    for i, num := range nums {
        if num > maxVal {
            secondMax = maxVal
            maxVal = num
            maxIdx = i
        } else if num > secondMax && num != maxVal {
            secondMax = num
        }
    }

    if maxVal >= secondMax*2{
        return maxIdx
    }

    return -1

    


}
