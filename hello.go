package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30}
	b := a[0:2]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(a)
	fmt.Println(b)
	// 出力----
	// [10 20 30] a
	// [10 20] b
	// [10 20 1000] a 値が変更されている
	// [10 20 1000] b
	// [10 20 1000] a 値が追加されていない
	// [10 20 1000 1000] b 値が追加されている
}
