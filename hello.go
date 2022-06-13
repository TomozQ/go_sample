package main

import "fmt"

func main() {
	//第一引数にstringのスライス、第二引数に関数を取る。戻り値はstringスライス。 ->第一引数で受けたスライスを第二引数の関数で処理して返却する。
	modify := func(a []string, f func([]string) []string) []string {
		return f(a)
	}
	m := []string{
		"1st",
		"2nd",
		"3rd",
	}
	fmt.Println("initial", m)

	m1 := modify(m, func([]string) []string {
		return append(m, m...)
	})
	fmt.Println("m1", m1)
	m2 := modify(m, func([]string) []string {
		return m[:len(m)-1]
	})
	fmt.Println("m2", m2)
	m3 := modify(m, func([]string) []string {
		return m[1:]
	})
	fmt.Println("m3", m3)
}
