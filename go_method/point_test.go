package main

import (
	"fmt"
	"reflect"
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

func dumpMethodSet(i interface{}) {
	dynType := reflect.TypeOf(i)
	if dynType == nil {
		fmt.Println("there is no dynamic type")
	}
	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method is empty \n", dynType)
		return
	}
	fmt.Printf("%s's method set: \n", dynType)
	for i := 0; i < n; i++ {
		fmt.Println("-", dynType.Method(i).Name)
	}
}

func TestPointMethod(t *testing.T) {

	var a1 Tyt
	fmt.Println(a1.i)
	a1.M1()
	fmt.Println(a1.i)
	a1.M2()
	fmt.Println(a1.i)

	dumpMethodSet(a1)

	var a2 = &Tyt{}
	fmt.Println(a2.i)
	a2.M1()
	fmt.Println(a2.i)
	a2.M2()
	fmt.Println(a2.i)
	dumpMethodSet(a2)

}
