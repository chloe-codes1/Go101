package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// Variadic function
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 4))
	// ignored value인  underscore(_) 사용하기
	totalLength, _ := lenAndUpper("chloe")
	fmt.Println(totalLength)
	repeatMe("chloe", "camila", "bella", "chalotte")
}
