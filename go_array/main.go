package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func foo(arr [5]int) {
	fmt.Println(arr)
	fmt.Printf("数组长度 %d\n", len(arr))
	fmt.Printf("数组大小 %d\n", unsafe.Sizeof(arr))
}

func main() {
	var ar1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(ar1)
	fmt.Printf("数组长度 %d\n", len(ar1))
	fmt.Printf("数组大小 %d\n", unsafe.Sizeof(ar1))

	var ar2 = [...]int{
		6, 7, 8, 9, 10,
	}
	foo(ar2)

	var arr3 = [100]int{
		99: 39,
	}
	fmt.Println(arr3)

	var arr1 = [2][3][4]int{}
	fmt.Println(arr1)

	var sc1 = []int{1, 2, 3}
	fmt.Println("slice sc1  value is ", sc1)
	fmt.Println("slice sc1  length is ", len(sc1))
	sc2 := append(sc1, 7)
	fmt.Println("slice sc2  value is ", sc2)
	fmt.Println("slice sc2  length is ", len(sc2))

	sc3 := make([]int, 6, 10)
	fmt.Println(sc3)
	fmt.Println(reflect.TypeOf(sc3))

	arr4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sc4 := arr4[3:8:8]
	sc5 := arr4[3:8:9]
	fmt.Println(sc4)
	fmt.Println(len(sc4))
	fmt.Println(unsafe.Sizeof(sc4), cap(sc4))
	fmt.Println(sc5)
	fmt.Println(len(sc5))
	fmt.Println(unsafe.Sizeof(sc5), cap(sc5))

	var m1 = make(map[string]string)
	m1["name"] = "hello map"
	fmt.Println(m1)
	//查找key是否存在
	_, ok := m1["name"]
	if !ok {
		fmt.Println("name not in map m1")
	}
	fmt.Println("name in map m1")

	//map的遍历
	fmt.Println("{")
	for k, v := range m1 {
		fmt.Println("key:", k, "---values:", v)

	}
	fmt.Println("}")

}
