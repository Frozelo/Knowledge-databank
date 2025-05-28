// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, 
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
/// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// https://leetcode.com/problems/valid-palindrome/description

// O(n)
func isPalindrome(s string) bool {
    s = strings.ToLower(s)

    left := 0
    right := len(s) - 1

    for left < right {
        for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])){
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])){
			right--
		}
        
        if s[left] != s[right] {
            return false
        }

        left++
        right--
    } 

    return true
}

func isPalindrome(s string) bool {
	// Шаг 1: Привести к нижнему регистру
	s = strings.ToLower(s)

	// Шаг 2: Удалить неалфавитно-цифровые символы
	var filteredRunes []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filteredRunes = append(filteredRunes, r)
		}
	}

	// Шаг 3: Проверить, является ли строка одинаковой с обеих сторон
	n := len(filteredRunes)
	for i := 0; i < n/2; i++ {
		if filteredRunes[i] != filteredRunes[n-i-1] {
			return false
		}
	}
	return true
}

