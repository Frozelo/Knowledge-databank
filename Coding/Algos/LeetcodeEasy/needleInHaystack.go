package main

import "fmt"

func main() {
	var haystack string = "sadbutsad"
	var needle string = "sad"

	for i := 0; i < len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			fmt.Println(i)
		} else {
			fmt.Println("-1")
		}
	}

}
