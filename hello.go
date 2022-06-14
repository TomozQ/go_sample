package main

import "fmt"

func main() {
	n := 123
	fmt.Printf("value: %d.\n", n)
	//値渡し
	change1(n) //change1の引数にはnのコピーが渡されている。関数が実行された後は消滅する値。change1内で行った操作は関数外にまったく影響を与えない。
	fmt.Printf("value: %d.\n", n)
	//参照渡し(ポインタ渡し)
	change2(&n)
	fmt.Printf("value: %d.\n", n)
}

func change1(n int) {
	n *= 2
}

//ポインタを渡すことでそのポインタにある値が2倍に書き換えられた。
func change2(n *int) {
	*n *= 2
}

// 出力
// value: 123.
// value: 123.
// value: 246.
