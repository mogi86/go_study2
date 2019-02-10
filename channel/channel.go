package channel

import "fmt"

func Receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println("channel:", i)
	}
}

func PrintNumber1(ch chan bool) {
	for i := 0; i <= 10; i++ {
		fmt.Println("並行1:", i)
	}
	ch <- true
}

func PrintNumber2(ch chan bool) {
	for i := 0; i <= 10; i++ {
		fmt.Println("並行2:", i)
	}
	ch <- true
}

func UseClose(ch chan int) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- i
	}
}