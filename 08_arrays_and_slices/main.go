package main

import "fmt"

func main() {
	// array
	names := [5]string{"chloe", "camila", "bella"}
	names[3] = "charlotte"
	names[4] = "bree"
	fmt.Println(names)

	// slice
	states := []string{"Florida", "Kentucky"}
	states = append(states, "California")
	fmt.Println(states)
}
