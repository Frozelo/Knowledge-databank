package main

import (
	"fmt"
)

func testSlices2() {
	a := []byte{'a', 'b', 'c'} // len a = 3; cap = 3
	b := append(a[1:2], 'd')   // a[1:2] = len 1; cap = 2 b = [b, d]; a = [a, b, d]
	b[0] = 'z'                 // a = [a z d]; b = [z d]

	fmt.Printf("%s\n", a) // что отобразится после вызова?
	// a = [a z d]
}

func main() {
	testSlices2()
}
