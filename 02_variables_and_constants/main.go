package main

import "fmt"

func main() {

	// 1) const without type
	const name = "chloe"

	// 2)const with type
	const surname string = "kim"

	// 3) variable without type
	var year = 2020

	// 3) variable with type
	var month int = 12
	// -> month := 12 (same)

	fmt.Println(year, month)

}
