package main

import (
	"fmt"
)

func main() {
	a := 2
	// pointer
	b := &a
	a = 10

	// change the value of object using pointer
	*b = 20
	fmt.Println(a, *b)
}
