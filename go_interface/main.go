package main

import (
	"fmt"
	"io"
	"strings"
)

type myInt int

func (n *myInt) add(m int) {
	*n = *n + myInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*myInt
	t
	io.Reader
	s string
	n int
}

func main() {
	m := myInt(17)
	r := strings.NewReader("hello word")
	s := S{
		myInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}
	sl := make([]byte, len("hello word"))
	//s.Reader.Read(sl)
	s.Read(sl)
	fmt.Println("sl string value:", string(sl))
	//fmt.Println("the print []byte :", string(io.Reader.Read(sl)))
	//s.myInt.add(5)
	s.add(5)
	fmt.Println("s.myInt value:", *(s.myInt))
}
