package main

import "fmt"

type book struct {
	title string
	price int
}

func (b book) name() string {
	return b.title
}

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return a
}

func main() {
	var b = book{title: "go program training", price: 32}
	fmt.Println(b.name(), b.price)

	b1 := &book{title: "java program training", price: 31}
	fmt.Println(b1.name(), b1.price)

	var a = T{a: 10}
	a.Set(20)
	fmt.Println(a.Get())

	var t T
	fmt.Println(t.Get())
	t.Set(10)
	(&t).Set(1)
	fmt.Println(t.Get())
}
