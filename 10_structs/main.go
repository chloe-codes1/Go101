package main

import "fmt"

// 1. Describe the shape of structure
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// 2. Give value to structure
	chloeFav := []string{"guacamole", "omelet", "biscuits and gravy"}
	chloe := person{"chloe", 26, chloeFav}

	camilaFav := []string{"sirloin", "grilled mushroom"}
	camila := person{name: "camila", age: 26, favFood: camilaFav}

	fmt.Println(chloe)
	fmt.Println(camila.name, camila.age)

	for _, value := range camilaFav {
		fmt.Print(value, " ")
	}
}
