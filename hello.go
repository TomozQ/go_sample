package main

import "fmt"

func main() {
	// 配列の定義
	a := [5]int{10, 20, 30, 40, 50}
	// 配列からスライスを生成
	b := a[0:3]
	// ※スライスの定義
	// 変数 := []型{値1, 値2 ...}
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 200
	fmt.Println(a)
	fmt.Println(b)
	// aとbは同じ値を参照している
}
