package main

import (
	"fmt"
	"github.com/mogi86/go_study2/channel"
	"github.com/mogi86/go_study2/closure"
	"github.com/mogi86/go_study2/construction"
	"github.com/mogi86/go_study2/pkg_math"
	"github.com/mogi86/go_study2/slice"
	"github.com/mogi86/go_study2/syncSample"
)

//standard package: https://golang.org/pkg/

func main() {
	fmt.Println("Hello, 世界")

	//slice
	list := slice.GetSlice()
	for _, v := range list {
		fmt.Println(v)
	}

	//closure
	f := closure.PrintClosure(100)
	f()
	f()
	f()

	//Structs
	c := construction.Vertex1{1, 9}
	c.SetX(1000)
	fmt.Println(c.Calc())

	//channel
	//バッファ指定なしのチャネルには一度に1つのものしか入れられない
	//次のものを入れるには最初に入れたものを取り出してからじゃないといけない
	ch := make(chan int)
	go channel.Receiver(ch)
	for i := 0; i <= 100; i++ {
		ch <- i
	}

	end1, end2 := make(chan bool), make(chan bool)
	go channel.PrintNumber1(end1)
	go channel.PrintNumber2(end2)
	//チャネルに書き込まれるまで待機する
	<-end1
	<-end2

	//sync
	syncSample.PrintUseSync()

	//use close
	ch2 := make(chan int)
	go channel.UseClose(ch2)
	for v := range ch2 {
		fmt.Println(v)
	}

	//それぞれチャネルの呼び出される割合確認
	cnt3, cnt4 := 0, 0
	ch3, ch4 := make(chan int), make(chan int)
	close(ch3)
	close(ch4)
	for j := 0; j < 1000; j++ {
		select {
		case <-ch3:
			cnt3++
		case <-ch4:
			cnt4++
		}
	}
	fmt.Println("cnt3の回数:", cnt3)
	fmt.Println("cnt4の回数:", cnt4)

	//-------------
	//standard pkg
	//-------------
	pkg_math.PrintMath()
}
