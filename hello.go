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
	for i := 0; i < len(ar); i++ { //len(ar)配列の要素数

		n, er := strconv.Atoi(ar[i])

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
