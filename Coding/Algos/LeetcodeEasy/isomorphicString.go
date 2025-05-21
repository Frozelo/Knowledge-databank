// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. 
// No two characters may map to the same character, but a character may map to itself.

// egg -> add true, foo -> bar false

package main

import "fmt"

func test(s, t string) bool {
	m := make(map[byte]byte)
	m2 := make(map[byte]byte)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sch, tch := s[i], t[i]

		if v, ok := m[sch]; ok {
			if v != tch {
				return false
			}
		} else {
			m[sch] = tch
		}

		// the second checking to resoleve (foo, bar) and (bar, foo) ordering
		if v, ok := m2[tch]; ok {
			if v != sch {
				return false
			}
		} else {
			m2[tch] = sch
		}

	}

	return true
}

func main() {
	s := "egg"
	t := "add"

	fmt.Println(test(s, t))
}
