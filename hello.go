package main

import "fmt"

func main() {
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}
func modify(d string) func() []string {
	m := []string{
		"1st",
		"2nd",
	}
	return func() []string {
		return append(m, d)
	}
}

// 出力
// [1st 2nd *新しい値*]
// [1st 2nd +new data+]
// m1とm2にはmodifyの戻り値である関数が代入される
// その後m1,m2が実行された際には、modifyが実行された際のm,dの値が保持されている。
// ^
// |
// クロージャ―
