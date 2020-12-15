package main

import "fmt"

func addAll(numbers ...int) int {
	total := 0
	// for loop w/ range
	for _, number := range numbers {
		total += number
	}
	return total
}

func add(numbers ...int) {
	// not gonna use this loop format
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i)
	}
}

func main() {
	total := addAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)
}
