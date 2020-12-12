package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type A struct {
	Name        string
	Price       int
	Description string
}

func main() {
	a1 := A{
		Name:        "Test",
		Price:       100,
		Description: "Foo",
	}

	a2 := A{
		Name:        "Test",
		Price:       200,
		Description: "Foo2",
	}

	fmt.Println(cmp.Equal(a1, a2))
	fmt.Println(cmp.Diff(a1, a2))
}
