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
	go total(100, c)
	// c(チャンネル)から値を取り出して出力
	fmt.Println("total: ", <-c)
}

// 出力
// total:  5050
// チャンネルから値を取得する場合、その値が送られてくるまで処理を待つ
// 今回の場合total関数でチャンネルに値が追加されるまではmain関数のfmt.Printlnは実行されない。