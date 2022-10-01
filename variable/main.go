package main

import "fmt"

var a, b, c = 1, "a variable string ", true

var (
	d = float32(10.0)
	e = int64(123123)
	f = string("an other variable string ")
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
