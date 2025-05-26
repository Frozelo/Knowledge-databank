// Given an array of strings strs, group the annagrams together. You can return the answer in any order.

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// https://leetcode.com/problems/group-anagrams/description/


// O(n*log(n) *k )
func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, str := range strs {
        runes := []rune(str)

        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })


        sortedStr := string(runes)

        m[sortedStr] = append(m[sortedStr], str)
    }

    res := make([][]string, 0, len(m))


    for _, words := range m {
        res = append(res, words)
    }

    return res
}
