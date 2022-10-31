package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	errors.New("bad things happened"),
}

func bad() bool {
	return false
}
func returnErrors() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}
func main() {
	err := returnErrors()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}
