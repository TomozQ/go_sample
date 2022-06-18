package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
)

// SrData is structure
type SrData struct {
	msg string
	mux sync.Mutex
}

func main() {
	sd := SrData{msg: "Start"}
	prmsg := func(nm string, n int){
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	main := func(n int){
		const nm string = "*main"
		sd.mux.Lock() //☆
		for i := 0; i < 5; i++{
			sd.msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
		sd.mux.Unlock() //☆
	}

	hello := func(n int){
		const nm string = "hello"
		sd.mux.Lock() //☆
		for i := 0; i < 5; i++ {
			sd.msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock() //☆
	}

	go main(100)
	go hello(50)
	time.Sleep(5 * time.Second)
}

// 出力
// hello Start h0
// hello Start h0 h1
// hello Start h0 h1 h2
// hello Start h0 h1 h2 h3
// hello Start h0 h1 h2 h3 h4
// *main Start h0 h1 h2 h3 h4 m0
// *main Start h0 h1 h2 h3 h4 m0 m1
// *main Start h0 h1 h2 h3 h4 m0 m1 m2
// *main Start h0 h1 h2 h3 h4 m0 m1 m2 m3
// *main Start h0 h1 h2 h3 h4 m0 m1 m2 m3 m4

// sd.mux.Lock() -> 他スレッドがsdにアクセスできなくなる。他スレッドはこの処理がアンロックされるまで待ち続けることになる。
// スレッドをロックするとアンロックされるまでの間、他スレッドはpending状態となる。

// 別スレッドで実行している処理を全てロックしてしまうと結局閉校処理ではなく「複数のスレッドが順に実行される」という逐次処理になってしまう。
// せっかくの並行処理の利点が失われてしまうから、ロックは「この処理だけは外部からアクセスされては困る」という必要最低限の範囲に絞って行うようにする。