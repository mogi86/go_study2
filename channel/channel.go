package channel

import "fmt"

func Receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("channel:", i)
	}
}