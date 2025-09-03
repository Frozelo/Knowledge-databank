package main

import "fmt"

func main() {
	x := []int{}      // init x slice; len = 0, cap = 0
	x = append(x, 0)  // len x = 1; cap = 1 [0]
	x = append(x, 1)  // len x = 2; cap = 2 [0, 1]
	x = append(x, 2)  // len x = 3; cap = 4 [0, 1, 2]
	y := append(x, 3) // len y = 4; cap = 4 y = [0, 1, 2, 3]
	z := append(x, 4) // len z = 4; cap = 4 y = z = [0, 1, 2, 4]
	fmt.Println(y, z) // что отобразится после вызова? [0, 1, 2, 4], [0, 1, 2, 4]
}
