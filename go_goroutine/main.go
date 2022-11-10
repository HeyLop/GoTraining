package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("this is a goroutine process")

	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 3)

	//fmt.Println()
	time.Sleep(time.Second * 1)
}
