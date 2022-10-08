package main

import "fmt"

type book struct {
	name string
	page int
}

func main() {
	var b book
	b.name = "Go Training"
	b.page = 100
	fmt.Println(b)

	var b1 book
	fmt.Println(b1)
}
