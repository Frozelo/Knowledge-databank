// A string is good if there are no repeated characters.
// Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.
// Note that if there are multiple occurrences of the same substring, every occurrence should be counted.
// A substring is a contiguous sequence of characters in a string.

func countGoodSubstrings(s string) int {
    var count int

	for i := 0; i < len(s)-2; i++ {

		m := make(map[byte]struct{})
		substring := s[i : i+3]
		var isUnique = true

		for j := 0; j < len(substring); j++ {
			if _, ok := m[substring[j]]; ok {
				isUnique = false
				break
			} else {
				m[substring[j]] = struct{}{}
			}
		}

		if isUnique == true {
			count++
		}
	}

	return count
}
