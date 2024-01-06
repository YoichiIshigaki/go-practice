package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

func main() {
	var student1 Student
	student1.name = "sato"
	student1.math = 60
	student1.english = 60

	fmt.Println(student1)

	student2 := Student{"sato", 60, 70}
	fmt.Println(student2)

	student3 := Student{math: 70, name: "sato", english: 60}
	fmt.Println(student3)

}
