package main

import (
	"fmt"
	"testing"
)

type Ty struct {
	a int
}

func (tt Ty) get() int {
	return tt.a
}

func (tt Ty) Set(a int) int {
	tt.a = a
	return tt.a
}
func (tt *Ty) SetPoint(a int) int {
	tt.a = a
	return tt.a
}

func TestMethod(t *testing.T) {
	var aa = Ty{a: 10}
	fmt.Println(aa.get())
	(&aa).Set(20)
	fmt.Println(aa.get())
	(&aa).SetPoint(20)
	fmt.Println(aa.get())
}
