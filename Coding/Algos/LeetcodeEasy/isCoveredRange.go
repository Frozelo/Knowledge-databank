// You are given a 2D integer array ranges and two integers left and right. Each ranges[i] = [starti, endi] represents an inclusive interval between starti and endi.
// Return true if each integer in the inclusive range [left, right] is covered by at least one interval in ranges. Return false otherwise.
// An integer x is covered by an interval ranges[i] = [starti, endi] if starti <= x <= endi.

// https://leetcode.com/problems/check-if-all-the-integers-in-a-range-are-covered

// non optimized approach O(n*m)
func isCovered(ranges [][]int, left int, right int) bool {
    for left <= right {
        covered := false
        for _, r := range ranges {
            if r[0] <= left && left <= r[1] {
                covered = true
                break
            }
        }

        if !covered {
            return false
        }

        left++
    }
    return true
}


// optimized approach - cover ints from 1 to 51 O(n+m)
func isCovered(ranges [][]int, left int, right int) bool {
    set := make([]bool, 52)

    for _, r := range ranges {
        for i := r[0]; i <= r[1]; i++ {
            set[i] = true
        } 
    }

    for i := left; i <= right; i++ {
        if !set[i] {
            return false
        }
    }

    return true
}
