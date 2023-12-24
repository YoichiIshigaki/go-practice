package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 2.基本構文
	// https://docs.google.com/presentation/d/1CIMDenDLZ7NPNgzmfbCNH_W3dYjaTEBdUYfUuXXuMHk/edit#slide=id.g129fad4b4e0_1_47
	x := 100
	x1 := 100_000_000
	println(x)
	println(x1)
	println("Hello World")

	// 定数は型を持たないので無限の精度になる
	const n = 10000000000000000000 / 10000000000000000000
	var m = n  // mはint型
	println(m) // 1

	// const (
	// 	a = 1 + 2
	// 	b
	// 	c
	// )
	// fmt.Println(a, b, c)

	const (
		a, b = iota, iota
		c    = iota
		d, e = iota, iota
	)
	fmt.Println(a, b, c, d, e)

	const (
		StatusOK = iota // 0
		StatusNG        // 1（StatusNG = iotaと同じ意味）
	)
	fmt.Println(StatusOK, StatusNG)

	const (
		FlagA = 1 << iota // 1
		FlagB             // 2 （FlagB = 1 << iotaと同じ意味）
		FlagC             // 4
		FlagD             // 8
	)

	fmt.Println(FlagA, FlagB, FlagC, FlagD)

	nn := 100 + 200
	// 変数を使った演算
	mm := nn + 100
	// 文字列の足し算
	msg := "hoge" + "fuga"

	fmt.Println(nn, mm, msg)

	// 制御構文
	if x == 1 {
		println("xは1")
	} else if x == 2 {
		println("xは2")
	} else {
		println("xは1でも2でもない")
	}

	aa := 1
	if aa == 0 {
		fmt.Println(aa)
	} else {
		fmt.Println("0ではない")
	}

	// 代入文を書く
	// if aaa := f(); a > 0 {
	// 	fmt.Println(aaa)
	// } else {
	// 	fmt.Println(2 * aaa)
	// }

	aaaa := 0
	switch aaaa {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	for i := 0; i <= 100; i = i + 1 {
		fmt.Println(i)
	}
	// 継続条件のみ
	// for j <= 10 {
	// 	fmt.Println(j)
	// }

	// var i int
	// LOOP: // for文にラベルをつける
	// for {
	// 	fmt.Println(i)
	// 	switch {
	// 	case i%2 == 1:
	// 		break
	// 	}
	// 	i++
	// }

	var i int
	for {
		fmt.Println(i)
		if i%2 == 1 {
			break
		}
		i++
	}
	// 現在時刻を数値で取得する
	t := time.Now().UnixNano()
	// 乱数のたねを設定
	rand.Seed(t)
	// xは0-10の間の値になる
	s := rand.Intn(10)
	println("結果は", s)
}
