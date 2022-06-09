package main

import "fmt"

func main() {
	x := 2
	switch x {
	case f(1):
		fmt.Println("* first case")
	case f(2):
		fmt.Println("* second case")
	case f(3):
		fmt.Println("* third case")
	default:
		fmt.Println("* default case")
	}
}

func f(n int) int {
	// Printlnは型の異なる値を複数の引数として指定し、出力することができる。
	fmt.Println("No, ", n)
	return n
}
