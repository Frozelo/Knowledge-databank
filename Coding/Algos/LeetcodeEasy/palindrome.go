// Given an integer x, return true if x is a palindrome , and false otherwise.
// example 121 is palindrome, 1 2 1 -> -121 is not a palindrome


package main

import "fmt"

// approach 1
func isPalindrome(x int) bool {
    
    str := strconv.Itoa(x)

    for i := 0; i < len(str) / 2; i++ {
        if(str[i] != str[len(str) - i - 1]) {
            return false
        }
    }

    return true
}

// approach 2
func main() {
	n := 121
	var res int
	for a := n; a > 0; {
		prem := a % 10
		res = res*10 + prem
		a = a / 10
	}

	fmt.Println(res, n)
	if res != n {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}



