// Testing
package main

import (
	"fmt"
	array "girado73/goLib/baselib/array"
)

func main() {
	s := make([]string, 3)

	s[0] = "Hello"
	s[1] = "World"
	s[2] = "Go"

	fmt.Println(array.ToString(s, ", "))
}
