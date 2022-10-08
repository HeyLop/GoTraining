package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		fmt.Println("this os is linux")
	}

	//a, b := true, false
	//if (a && b) != true {
	//	println("(a && b) != true")
	//	return
	//}
	//println("a && (b != true) == false")

	for i := 0; i < 10; i++ {

		if i == 5 {
			fmt.Println("error num ", i)
			continue
		}
		fmt.Println(i)
	}

}
