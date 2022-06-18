package main

import(
	"fmt"
)

func total(n int, c chan int){
	t := 0
	for i := 1; i <= n; i++{
		t += i
	} 
	// チャンネルに t を格納する
	c <- t
}

func main() {
	// int型を格納するチャンネルを生成
	c := make(chan int)
	// total関数を並列処理で実行
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	
	x, y, z := <-c, <-c, <-c // <-cを呼びだすごとにチャンネルから値が取り出されている。3回実行することで全ての値を取り出せる。
	fmt.Println(x, y, z)
}

// 出力
// 55 500500 5050
// 実行したスレッドのかかる時間によりチャンネルに値が入る順番は変わる。
// Goルーチンを呼びだした順に値が追加されるわけでは無い。