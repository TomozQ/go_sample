package main

import "fmt"

var mydata struct {
	Name string
	Data []int
}

func main() {
	//構造体
	// struct{
	// 	...変数宣言...
	// }

	// 変数を指定した定義の構造体の値として宣言できる。
	// var 変数 struct{......}
	// 利用する際には 構造体.変数
	mydata.Name = "Taro"
	mydata.Data = []int{10, 20, 30}
	fmt.Println(mydata)
}
