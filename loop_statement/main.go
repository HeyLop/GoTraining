package main

import (
	"fmt"
	"runtime"
	"time"
)

func readByExtSwitch(ext string) {
	switch ext {
	case "json":
		fmt.Println("this is json")
	case "yml":
		fmt.Println("this is yaml")
	case "ini":
		fmt.Println("this is ini")
	default:
		fmt.Println("un know file type ")

	}

}

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

	s := "中国人"
	for k, v := range s {
		fmt.Printf("the num is %d the letter is %s the byte string is %b \n", k, string(v), v)
	}

	// map，for遍历
	var m1 = map[string]int{
		"happy":  10,
		"flower": 11,
		"leaf":   12,
	}
	for k, v := range m1 {
		fmt.Printf("the map key is %s,the map vaule value is %d \n ", k, v)
	}

	//带label带for循环
	var cl1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum int
loop1:
	for i := 0; i < len(cl1); i++ {
		if cl1[i]%2 == 0 {
			continue loop1
		}
		sum += cl1[i]
	}
	//fmt.Println(sum)
	fmt.Printf("the cl1 sum is %d\n", sum)

	//for循环多变量
	for a, b, c := 0, 1, 2; (a < 10) && (b < 15) && (c < 20); a, b, c = a+1, b+1, c+1 {
		fmt.Println(a, b, c)
	}

	//break
	var sc1 = []int{5, 9, 7, 10, 8, 3}
	for i := 0; i < len(sc1); i++ {
		if sc1[i]%2 == 0 {
			fmt.Printf("the slice sc1 frist even number is %d\n", sc1[i])
			break
		}
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

	var ext = "json"
	readByExtSwitch(ext)
	readByExtSwitch("txt")

}
