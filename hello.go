package main

import(
	"fmt"
	"time"
	"strconv"
)

func hello(s string, n int){
	for i := 1; i <= 10; i++{
		fmt.Printf("<%d %s>", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main() {
	msg := "start"
	prmsg := func(nm string, n int){
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	hello := func(n int){
		const nm string = "hello"
		for i := 0; i < 10; i++{
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	main := func(n int){
		const nm string = "*main"
		for i := 0; i < 5; i++{
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}

	go hello(60)
	main(100)
}

//出力
// *main start m0
// hello start m0 h0
// hello start m0 h0 h1
// *main start m0 h0 h1 m1
// hello start m0 h0 h1 m1 h2
// hello start m0 h0 h1 m1 h2 h3
// *main start m0 h0 h1 m1 h2 h3 m2
// hello start m0 h0 h1 m1 h2 h3 m2 h4
// hello start m0 h0 h1 m1 h2 h3 m2 h4 h5
// *main start m0 h0 h1 m1 h2 h3 m2 h4 h5 m3
// hello start m0 h0 h1 m1 h2 h3 m2 h4 h5 m3 h6
// *main start m0 h0 h1 m1 h2 h3 m2 h4 h5 m3 h6 m4
// hello start m0 h0 h1 m1 h2 h3 m2 h4 h5 m3 h6 m4 h7
// hello start m0 h0 h1 m1 h2 h3 m2 h4 h5 m3 h6 m4 h7 h8

// hello関数内のループはh0 ~ h9の10回繰り返されるはずだがh9が出力されていない。
// →10回呼ばれる前にメインスレッドの処理が終了している。
// メインスレッドが終了すると、それ以外のその他のスレッドは処理中であっても全て消えてしまう。