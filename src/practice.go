package main

import "fmt"

type Student struct {
	name string
}

func (s Student) calcAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main() {
	a001 := Student{name: "sato"}
	data := []float64{70, 65, 50, 90, 30}
	var avg float64 = a001.calcAvg(data)
	result := a001.judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + " " + result)

}
