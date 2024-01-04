package main

import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting)
}

func main() {
	sayHello("Good Morning")
	sayHello("Good Afternoon")
	sayHello("Good Night")
	calTriple(3)
	// 戻り値が二つ
	result01, result02 := cal(6, 2)
	fmt.Println(result01, result02)

	// 無名関数
	hello := func(greeting string) {
		fmt.Println(greeting)
	}

	hello("Hello World!!")

	func(name string) {
		fmt.Println("Hello " + name + " !!")
	}("John")

}

func calTriple(x int) {
	fmt.Println(x * 3)
}

func cal(x, y int) (int, int) {
	a := (x / y)
	b := (x * y)

	return a, b
}
