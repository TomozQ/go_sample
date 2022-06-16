package main

import "fmt"

//interface
// type 名前 interface {
// 	メソッドA
// 	メソッドB
// 	...必要なだけ用意...
// }

//Data is interface
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

//Mydata is struct
type Mydata struct {
	Name string
	Data []int
}

//Initial is init method
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

//PrintData is println all data
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

//Check is method
func (md *Mydata) Chack() {
	fmt.Printf("Check! [%s]", md.Name)
}

func main() {
	var ob Mydata = Mydata{}
	// new関数は特定の型の値を作成するものではない。代入する変数の型に合わせて値は扱われる。
	// var ob Data = new(Mydata) //この場合はData型の変数が定義されているのでここにnewで代入した値はData型となる。
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
	ob.Chack()
}

//インターフェースは直接値を作成することはできない → Data()やData{}はできない

// いずれもobに代入されているのはMydataでCheck関数を使用できそうだが、「実際に何が入っているか」よりも「その型は何なのか」ということによって用意されているメソッドを認識している。
//「インターフェイス型として変数を用意する場合は、その値はインターフェイスに用意された機能しか使えない。」
// ⇓
// 下記は実行できる
// func main() {
// 	var ob Mydata = Mydata{} //←この部分
// 	ob.Initial("Sachiko", []int{55, 66, 77})
// 	ob.PrintData()
// 	ob.Chack()
// }

// 下記は実行できない
// func main() {
// 	var ob Data = new(Mydata) //←この部分
// 	ob.Initial("Sachiko", []int{55, 66, 77})
// 	ob.PrintData()
// 	ob.Chack()
// }
