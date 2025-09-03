package main

import (
	"fmt"
)

func testSlices1() {
	a := []string{"a", "b", "c"} // len a = 3; cap a = 3
	b := a[1:2]                  // len = 1; cap = 2
	b[0] = "q"                   // a = [a, q, c]; b = [q, c]

	fmt.Printf("%s\n", a) // что отобразится после вызова? [a q c]
}

func main() {
	testSlices1()
}
