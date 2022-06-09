package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	// n, err := strconv.Atoi(x)

	// ifに入るとまずn, err := strconv.Atoi(x)が実行される。errがnilの場合には偶数奇数の判定に入る。
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("整数ではありません。")
	}
	// if err != nil {
	// 	fmt.Println("ERROR!!")
	// 	return
	// }
	// fmt.Print(x + "は、")
	// if n%2 == 0 {
	// 	fmt.Println("偶数です")
	// } else {
	// 	fmt.Println("奇数です")
	// }
}
