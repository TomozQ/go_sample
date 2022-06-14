package main

import "fmt"

func main() {
	n := 123
	//&nでnのポインタを取得
	p := &n //ポインタを代入すると自動推論によりポインタ型と認識される。今回nはint型なので*int型というポイント型の変数となる。
	fmt.Println("number: ", n)
	fmt.Println("pointer: ", p)
	//pは自動推論によりポインタ型とされているので*pとして値を取得する。
	fmt.Println("value: ", *p)
}

// 出力
// number:  123
// pointer:  0xc000012088
// value:  123
