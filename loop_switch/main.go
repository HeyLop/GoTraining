package main

import "fmt"

func switchInit() int {
	fmt.Println("this is switch_init")
	return 2
}
func case1() int {
	fmt.Println("this is case1")
	return 1
}
func case2() int {
	fmt.Println("this is case2")
	return 2
}
func case3() int {
	fmt.Println("this is case3")
	return 3
}

func main() {
	fmt.Println("exec 1------switch")
	switch switchInit() {
	case case1():
		fmt.Println("exec case1()")
	case case2():
		fmt.Println("exec case2()")
	case case3():
		fmt.Println("exec case3()")
	default:
		fmt.Println("exec default message!")
	}

	fmt.Println("exec 2------switch")
	switch switchInit() {
	case case1():
		fmt.Println("exec case1()")
		fallthrough
	case case2():
		fmt.Println("exec case2()")
		fallthrough
	case case3():
		fmt.Println("exec case3()")
		//fallthrough
	default:
		fmt.Println("exec default message!")
	}
}
