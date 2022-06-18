package main

import (
	"fmt"
	"strconv"
	"time"
)

func prmsg(n int, s string){
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string){
	const nm string = "first"
	for i := 0; i < 10; i++{
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string){
	for i := 0; i < 10; i++{
		prmsg(n, "second:["+ <-c +"]")
	}
}

func main() {
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println()
}

// 出力
// first0
// first1
// second:[first0]
// first2
// second:[first1]
// first3
// second:[first2]
// first4
// second:[first3]
// first5
// second:[first4]
// first6
// second:[first5]
// first7
// second:[first6]
// first8
// second:[first7]
// first9
// second:[first8]
// second:[first9]