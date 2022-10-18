package main

import "fmt"

func foo() {
	fmt.Println("calling function foo")
	bar()
	fmt.Println("exit function foo")
}

func bar() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("the function bar have a panic event,and the panic error:", e)
		}
	}()
	fmt.Println("calling function bar")
	panic("panic in function bar")
	zoo()
	fmt.Println("exit function bar")
}

func zoo() {
	fmt.Println("calling function zoo")
	bar()
	fmt.Println("exit function zoo")
}

func main() {
	fmt.Println("calling function main")
	foo()
	fmt.Println("exit function main")
}
