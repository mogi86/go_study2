package main

import (
	"fmt"
	"github.com/mogi86/go_study2/closure"
	"github.com/mogi86/go_study2/construction"
	"github.com/mogi86/go_study2/slice"
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
}
