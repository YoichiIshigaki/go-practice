package main

import "fmt"

func main() {
	var f float64 = 10
	var n int = int(f)

	println(n)

	var sum float64
	sum = 5 + 6 + 3
	avg := sum / 3
	println(avg)
	if avg > 4.5 {
		println("good")
	} else {
		println("bad")

	}

	var a, b, c bool
	println(a, b, c)
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}

	var ns []int

	var m map[string]int

	println(ns)
	println(m)

	p := struct {
		name string
		age  int
	}{name: "Gopher", age: 10}
	// フィールドにアクセスする例
	p.age++
	println(p.name, p.age)

	// ゼロ値で初期化
	var ns1 [5]int
	// 配列リテラルで初期化
	var ns2 = [5]int{10, 20, 30, 40, 50}
	// 要素数を値から推論
	ns3 := [...]int{10, 20, 30, 40, 50}
	// 5番目が50、10番目が100で他が0の要素数11の配列
	ns4 := [...]int{5: 50, 10: 100}
	fmt.Println(ns1, ns2, ns3, ns4)

	arr := [...]int{10, 20, 30, 40, 50}
	// 要素にアクセス
	println(arr[3]) // 添字は変数でもよい
	// 長さ
	println(len(ns))
	// スライス演算
	fmt.Println(arr[1:3])

	aa := []int{10, 20}
	// [10 20] 2
	fmt.Println("aa = ", aa, cap(aa))

	bb := append(aa, 30) // (1)
	aa[0] = 100          // (2)
	// [10 20 30] 4
	fmt.Println(bb, cap(bb))

	cc := append(bb, 40) // (3)
	bb[1] = 200          // (4)
	// [10 200 30 40] 4
	fmt.Println(cc, cap(cc))
}
