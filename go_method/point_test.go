package main

import (
	"fmt"
	"testing"
)

type Tyt struct {
	i int
}

func (ty Tyt) M1() {
	ty.i = 10
}

func (ty *Tyt) M2() {
	ty.i = 11
}

func TestPointMethod(t *testing.T) {

	var a1 Tyt
	fmt.Println(a1.i)
	a1.M1()
	fmt.Println(a1.i)
	a1.M2()
	fmt.Println(a1.i)

	var a2 = &Tyt{}
	fmt.Println(a2.i)
	a2.M1()
	fmt.Println(a2.i)
	a2.M2()
	fmt.Println(a2.i)

}
