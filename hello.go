package main

import (
	"fmt"
	"time"
)

func total(c chan int){
	n := <-c  // チャンネルは、値がまだ用意されていない場合には、送られてくるまで処理を待つ。 -> 値が得られないというエラーになることはない。
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
}

func main() {
	c := make(chan int)
	go total(c)
	c <- 100 // チャンネルは複数のスレッド間で値をやり取りするためのものなので、双方で準備ができていないとならない => Goルーチンを実行した後でないとチャンネルは使えない。
	time.Sleep(100 * time.Millisecond)
}
