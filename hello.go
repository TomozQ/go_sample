package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	// 要素5つのスライスを初期化
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)
}

// 出力
// どちらにも&がついていることからtaroはポインタだとわかる
// &{ []}
// &{Taro [0 0 0 0 0]}
