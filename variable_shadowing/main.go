package main

import "fmt"

var a = 10

func foo(n int) {
	a := 1
	a += n
}
func main() {
	fmt.Println("a=", a)
	foo(5)
	fmt.Println("false after calling foo,a=", a)
}
