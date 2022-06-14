package main

import "fmt"

func main() {
	n := 123
	//&nでnのポインタを取得
	p := &n //ポインタを代入すると自動推論によりポインタ型と認識される。今回nはint型なので*int型というポイント型の変数となる。
	q := &p
	m := 10000
	p2 := &m
	q2 := &p2
	fmt.Printf("q value:%d, address:%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("q value:%d, address:%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)
}

// 出力
// q value:123, address:0xc000012088
// q2 value:10000, address:0xc0000120a0
// q value:10000, address:0xc0000120a0
// q2 value:123, address:0xc000012088
