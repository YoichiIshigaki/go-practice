package main

import "fmt"

func main() {
	x := 10
	y := 12

	if age := x + y; age >= 20 {
		fmt.Println(age)
		fmt.Println("adult")
	} else if age == 0 {

		fmt.Println("baby")
	} else {
		fmt.Println("child")

	}
	// fmt.Println(age)
}
