package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

//	md(Mydata)の部分をレシーバーと呼ぶ
func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}
func main() {
	taro := Mydata{"Hanako", []int{98, 76, 54, 32, 10}}
	taro.PrintData()
}

// 出力
// *** Mydata ***
// Name:  Hanako
// Data:  [98 76 54 32 10]
// *** end ***
