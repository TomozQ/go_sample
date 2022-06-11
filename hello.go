package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("input data")
	// 入力された値を半角スペースで区切って配列にする ex)[10 20 30]
	ar := strings.Split(x, " ")
	// fmt.Println(ar)
	t := 0
	// 使わなくても良いとする変数は「 _ 」で表現する -> Goでは宣言された変数は使わなくてはならないため。 _ にすれば使用しなくてよい。
	for _, v := range ar { //_ -> インデックス, v -> 要素,range ar ->呼びだされるたび配列から順番にインデックスと要素を取り出す。全てが取り出されたらforを抜ける

		n, er := strconv.Atoi(v)

		if er != nil {
			goto err
		}

		t += n
	}
	fmt.Println("total: ", t)
	return
err:
	fmt.Println("ERROR!")
}
