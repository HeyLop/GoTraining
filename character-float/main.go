package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 123.456789
	fmt.Println(a)
	bits := math.Float64bits(a)
	fmt.Println(bits)

	var s string = "hello string character"
	fmt.Println(s)

	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0
	fmt.Println(math.Float32bits(f1))
	fmt.Println(math.Float32bits(f2))

	//go支持原生字符串，所见即所得
	var gopher string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(gopher)

}
