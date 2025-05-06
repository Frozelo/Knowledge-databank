```go
package main

import (
	"fmt"
)

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"

	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

func main() {
	testSlices1()
}
```

Вывод - ([a, q, c])
