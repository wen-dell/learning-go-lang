package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Printf("i está valendo %d\n", i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}
}
