package main

import (
	"fmt"
	"github.com/mogi86/go_study2/channel"
	"github.com/mogi86/go_study2/closure"
	"github.com/mogi86/go_study2/construction"
	"github.com/mogi86/go_study2/slice"
	"github.com/mogi86/go_study2/syncSample"
)

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
	<-end1
	<-end2

	//sync
	syncSample.PrintUseSync()
}
