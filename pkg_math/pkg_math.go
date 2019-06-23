package pkg_math

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintMath() {
	// seed is absolutely necessary!
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}