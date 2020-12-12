package closure

import "fmt"

func PrintClosure(a int) func() {
	i := 0

	return func() {
		i += 1
		fmt.Println(i + a)
	}
}
