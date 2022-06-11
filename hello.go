package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	t := 0
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err //errラベルの処理にジャンプする
	}
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
	return

err: //ラベル
	fmt.Println("ERROR!")
}
