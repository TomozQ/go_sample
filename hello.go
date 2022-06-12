package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)

	a = push(a, 1000)
	fmt.Println(a)

	a = pop(a)
	fmt.Println(a)

	a = unshift(a, 10000)
	fmt.Println(a)

	a = shift(a)
	fmt.Println(a)

	a = insert(a, 1000, 2)
	fmt.Println(a)

	a = remove(a, 2)
	fmt.Println(a)
}

// 最後に追加
func push(a []int, v int) []int {
	return append(a, v)
}

// 最後を削除
func pop(a []int) []int {
	return a[:len(a)-1]
}

// 最初に追加
func unshift(a []int, v int) []int {
	// ... -> スライスを展開する =>a[0], a[1], a[2]
	return append([]int{v}, a...) //vが入ったスライスを用意し、第2引数にaを展開したものを指定する。 => [v, a[0], a[1], a[2]]
}

// 最初を削除
func shift(a []int) []int {
	return a[1:] //インデックス番号1から後のみを返却
}

// 指定の位置に追加
func insert(a []int, v int, p int) []int {
	//a = [10, 20, 30]
	//	aの最後に0を追加
	a = append(a, 0)
	//a = [10, 20, 30, 0]
	//v -> 1000, p -> 2
	a = append(a[:p+1], a[p:len(a)-1]...) //a[:3] -> [10, 20, 30], a[2:3]... -> 30
	//a = [10, 20, 30, 30]
	a[p] = v
	// a[p] -> a[2] -> 30 = 1000
	// a = [10, 20, 1000, 30]
	return a
}

// 指定の位置を削除
func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...) //a[:p] ->[10, 20], a[3:]... -> 30
}

// 出力
// [10 20 30]			# 初期状態
// [10 20 30 1000]		# 最後に追加
// [10 20 30]			# 最後を削除
// [10000 10 20 30]		# 最初に追加
// [10 20 30]			# 最初を削除
// [10 20 1000 30]		# インデックス2に追加
// [10 20 30]			# インデックス2を削除
