package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

func recRes(x []int) ([]int, error) {
	if reflect.TypeOf(x).Elem().Kind() == reflect.Int {

	}
	return nil, errors.New("this character is not slice")
}

//传入任意个形参
func appendElem(s1 []int, elem ...int) []int {
	if len(elem) == 0 {
		fmt.Println(" no elem to append")
		return s1
	}
	s1 = append(s1, elem...)
	return s1
}

//匿名函数
var (
	myFPrintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a)
	}
)

func setup(task string) func() {
	fmt.Println("do some setup stuff for ", task)
	return func() {
		fmt.Println("do some teardown stuff for ", task)
	}
}

//乘法计算函数
func times(x, y int) (z int) {
	z = x * y
	return
}

func parseTimes(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reflect.TypeOf(x))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(appendElem(s1))
	fmt.Println(appendElem(s1, 4, 5, 6))

	//calling 匿名函数myFPrintf
	fmt.Println(reflect.TypeOf(any(myFPrintf)))
	myFPrintf(os.Stdout, "%s\n", "hello go")

	//calling function setup
	teardown := setup("demo-task")
	defer teardown()
	fmt.Println("do something business stuff ")

	//calling function times
	fmt.Println(times(2, 4))
	fmt.Println(times(2, 5))
	fmt.Println(times(2, 6))
	//calling function parseTimes
	twoTimes := parseTimes(2)
	fmt.Println(twoTimes(4))
	fmt.Println(twoTimes(5))
	fmt.Println(twoTimes(6))

	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", "happy", 7)
	fmt.Println(n)
	fmt.Println(err)
}
