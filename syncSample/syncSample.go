package syncSample

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintUseSync() {
	wg.Add(1)
	go func() {
		wg.Done()
		fmt.Println("syncSample1")
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		fmt.Println("syncSample2")
	}()

	wg.Wait()
	fmt.Println("syncSample finished!!")
}
