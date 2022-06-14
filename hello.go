package main

import "fmt"

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initial(&ar)
	fmt.Println(ar)
}

func initial(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0 //(*ar) -> arのポインタにあるスライスを取得できる。※*ar[i]だとar[i]にあるポインタを取得しようとしていると判断される。
	}
}

// 出力
// [10 20 30]
// [0 0 0]
