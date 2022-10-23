package main

import (
	"fmt"
	"reflect"
	"testing"
)

func DumpMethod(i interface{}) {
	fType := reflect.TypeOf(i)
	if fType == nil {
		fmt.Println("there is no dynamic type!")
		return
	}
	fNum := fType.NumMethod()
	if fNum == 0 {
		fmt.Printf("the %s's methed is empty!", fType)
		return
	}
	fmt.Printf("the %s's methed is %d --- \n", fType, fNum)
	for i := 0; i < fNum; i++ {
		fmt.Println(fType.Method(i).Name)
	}
	fmt.Println()

}

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (t T) M3() {
}

func Test_Inter(t *testing.T) {
	fmt.Println("testing Test_inter")
	var t1 T
	var t2 *T
	DumpMethod(t1)
	DumpMethod(t2)
}
