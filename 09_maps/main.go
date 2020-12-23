package main

import "fmt"

func main() {
	// map
	chloe := map[string]string{"name": "chloe", "age": "26"}
	fmt.Println(chloe)

	// iterate over a map
	for key, value := range chloe {
		fmt.Println(key, value)
	}
}
