package main

import "fmt"

func main() {
	n := 123
	//&nでnのポインタを取得
	p := &n //ポインタを代入すると自動推論によりポインタ型と認識される。今回nはint型なので*int型というポイント型の変数となる。
	m := 10000
	p2 := &m
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
}

// 出力
// p value:123, address:0xc000012088
// p2 value:10000, address:0xc0000120a0
// p value:10000, address:0xc0000120a0
// p2 value:123, address:0xc000012088
