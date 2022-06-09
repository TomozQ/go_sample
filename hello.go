package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print("1から" + x + "の合計は、")
	} else {
		return
	}
	t := 0
	c := 1
	// cがnと等しいか小さい間繰り返す
	for c <= n {
		t += c
		// ほかの言語と違って++や--を式の中で使用することはできない。
		c++
	}
	fmt.Println(t, "です。")
}
