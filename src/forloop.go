package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		if i == 3 {
			// break
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n ------- doubleLoop ------- \n")
	doubleLoop()
	fmt.Println("\n ------- arrayLoop ------- \n")
	arrayLoop()
	fmt.Println("\n ------- infinityLoop ------- \n")
	infinityLoop()
}

func doubleLoop() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
	}
}

func arrayLoop() {
	arr := [...]int{2, 4, 6, 8, 10}

	sum := 0
	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
		sum += arr[i]
	}
	fmt.Println("sum = ", sum)
}

func infinityLoop() {
	i := 0
	for {
		fmt.Println(i)
		if i == 4 {
			break
		}
		i++
	}
}
