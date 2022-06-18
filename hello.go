package main

import (
	"fmt"
	"time"
)

// select{
// case 文:
// 	...実行する処理...
// case 文:
// 	...実行する処理...
// case 文:
// 	...実行する処理...
// default:
// 	...全てに当てはまらない場合の処理...
// }

// switchと違いcaseには値ではなく操作の文を指定する。基本的には「<-チャンネル」という形でチャンネルから値を取り出す文を記述しておく。（チャンネルから値を取り出せたcaseに進む）

func count (n int, s int, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
}

func main() {
	n1, n2, n3 := 3, 5, 10
	m1, m2, m3 := 100, 75, 50
	c1 := make(chan int)
	go count(n1, m1, c1)
	c2 := make(chan int)
	go count(n2, m2, c2)
	c3 := make(chan int)
	go count(n3, m3, c3)

	for i := 0; i < n1 + n2 + n3; i++ {
		select{
		case re := <- c1:
			fmt.Println("*  first", re)
		
		case re := <- c2:
			fmt.Println("** second", re)
		
		case re := <- c3:
			fmt.Println("***third", re)
		}
	}
	fmt.Println("*** finish. ***")
}

// 出力
// ***third 1
// ** second 1
// *  first 1
// ***third 2
// ** second 2
// ***third 3
// *  first 2
// ***third 4
// ** second 3
// ***third 5
// *  first 3
// ** second 4
// ***third 6
// ***third 7
// ** second 5
// ***third 8
// ***third 9
// ***third 10
// *** finish. ***