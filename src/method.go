package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

// func (s Student) avg() {
// 	fmt.Println(s.name, (s.math+s.english)/2)
// }

func (s Student) avg(math, english float64) (avg float64) {
	avg = (math + english) / 2
	return
}

func main() {

	student := Student{math: 70, name: "sato", english: 60}
	avg := student.avg(student.math, student.english)
	fmt.Println(avg)
	// fmt.Println(student)

}
