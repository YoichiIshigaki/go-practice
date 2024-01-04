package main

import (
	"fmt"
	"slices"
	// "golang.org/x/exp/slices"
)

func main() {
	ns := []int{10, 20, 30, 40, 50}

	// 削除: [10 40 50]
	ns = slices.Delete(ns, 1, 3)
	fmt.Println(ns)

	// 挿入: [10 60 70 40 50]
	ns = slices.Insert(ns, 1, 60, 70)
	fmt.Println(ns)

	// 要素があるか: true
	ok := slices.Contains(ns, 70)
	fmt.Println(ok)

	// ソート: [10 40 50 60 70]
	slices.Sort(ns)
	fmt.Println(ns)

	a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}

	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])
}
