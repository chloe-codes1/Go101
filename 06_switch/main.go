package main

import "fmt"

func canIDrink(age int) bool {
	switch age {
	case 18:
		return false
	case 19:
		return true
	}
	return false
}

// Avoid if else with switch
func canIDrive(age int) bool {
	switch {
	case age < 19:
		return false
	case age == 19:
		return true
	case age > 19:
		return true
	}
	return false
}

// Variable expression in switch
func canIVote(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 19:
		return false
	case koreanAge == 19:
		return true
	case koreanAge > 19:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(19))
	fmt.Println(canIDrive(18))
	fmt.Println(canIVote(29))
}
