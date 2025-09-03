package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3} // len a = 3; cap = 3
	b := a              // len b = 3; cap = 3
	b = append(b, 4)    // new backing array for b; len b = 4; cap = 6 [1, 2, 3, 4]
	c := b              // len c = 4; cap = 6
	b[0] = 0            // b = [0, 2, 3, 4]; c = [0, 2, 3, 4]
	e := append(c, 5)   // len e = 5; cap e = 6; [0, 2, 3, 4, 5]
	b[2] = 7            // b = [0, 2, 7, 4]; c = [0, 2, 7, 4]; e = [0, 2, 7, 4, 5]

	fmt.Println(a, b, c, e) // что отобразится после вызова?
	// a = [1, 2, 3]; b = [0, 2, 7, 4]; c = [0, 2, 7, 4]; e = [0, 2, 7, 4, 5]
}
