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

// Naked return
func lenAndLower(name string) (length int, lowercase string) {
	// defer
	defer fmt.Println("Done!")

	length = len(name)
	lowercase = strings.ToLower(name)
	return
}

func main() {
	fmt.Println(multiply(2, 4))
	// ignored value인  underscore(_) 사용하기
	totalLength, _ := lenAndUpper("chloe")
	fmt.Println(totalLength)
	repeatMe("chloe", "camila", "bella", "chalotte")
	length, name := lenAndLower("CHLOE")
	fmt.Println(length, name)
}
