package main

import "fmt"

func init() {
	fmt.Println("init-1 invoked")
}

func init() {
	fmt.Println("init-2 invoked")
}

func main() {
	fmt.Println("go process execution")

}
