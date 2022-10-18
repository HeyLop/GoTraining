package main

import (
	"errors"
	"fmt"
)

func doSomething(s string) (string, error) {
	if s == "ok" {
		return s, nil
	}
	return s, errors.New("the error character")
}

var ErrSentinel = errors.New("the ErrorSentinel Error")

type myError struct {
	e string
}

func (e *myError) Error() string {
	return e.e
}

func main() {

	//wrap error,error as
	var err00 = &myError{"MyError Demo"}
	err11 := fmt.Errorf("wrap error00 : %w", err00)
	err22 := fmt.Errorf("wrap error11 : %w", err11)
	var e *myError
	if errors.As(err22, &e) {
		fmt.Println("the error is on the chain of err2")
		fmt.Println(err22 == e)
		fmt.Println(err00 == e)
		return
	}
	fmt.Println("the error is not on the chain of err2")

	//warp error，error is
	s, err := doSomething("error")
	if err != nil {
		fmt.Println("calling function doSomething hava a err,the character is ", s)
		fmt.Println("the error is ", err)
	}

	err1 := fmt.Errorf("Warp error ErrSentinel %w ", ErrSentinel)
	err2 := fmt.Errorf("warp error err1 %w", err1)
	fmt.Println(err2 == ErrSentinel)
	if errors.Is(err2, ErrSentinel) {
		fmt.Println("the err2 is ErrSentinel ")
		return
	}
	fmt.Println("the err2 not is ErrSentinel ")

}

//todo error 分支模式
