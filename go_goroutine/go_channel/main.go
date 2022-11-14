package main

import "fmt"

func main() {

	ch := make(chan int, 6)
	fmt.Println("chan:", ch)

}
