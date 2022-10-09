package main

import (
	"fmt"
	"runtime"
	"time"
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

	//for循环多变量
	for a, b, c := 0, 1, 2; (a < 10) && (b < 15) && (c < 20); a, b, c = a+1, b+1, c+1 {
		fmt.Println(a, b, c)
	}

	fmt.Println("for-01")
	var m = []string{"a", "b", "c", "d"}
	for k, v := range m {
		fmt.Printf("the key is %d,the value is %s.\n ", k, v)
	}

	fmt.Println("for-02")
	for k, v := range m {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Printf("the key is %d,the value is %s.\n ", k, v)
		}()
	}
	time.Sleep(time.Second * 5)

	fmt.Println("for-3")
	for k, v := range m {
		go func(k int, v string) {
			time.Sleep(time.Second * 1)
			fmt.Printf("the key is %d,the value is %s.\n ", k, v)
		}(k, v)
	}
	time.Sleep(time.Second * 5)

}
