package main

import "fmt"

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func canIDrive(age int) bool {
	// create a variable right incide of the "if"
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(19))
	fmt.Println(canIDrive(20))
}
