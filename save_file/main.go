package main

import (
	"bytes"
	"fmt"
	"io"
)

func Save(w io.Writer, date []byte) error {
	return nil
}

func main() {
	fmt.Println(bytes.ToUpper([]byte("hello")))
}
