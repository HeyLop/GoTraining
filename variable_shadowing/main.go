package main

import "fmt"
import "errors"

var aa = 10

func foo(n int) {
	aa := 1
	aa += n
}
func bar() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
		fmt.Println(a, b, c)
	}
}

var a int = 2020

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a {
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)
	return err
}

type new int

func getYear() (new, error) {
	var b int16 = 2021
	return new(b), nil
}

func main() {

	fmt.Println("a=", a)
	foo(5)
	fmt.Println("false after calling foo,a=", a)
	fmt.Println("calling function bar ")
	bar()

	err := checkYear()
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
