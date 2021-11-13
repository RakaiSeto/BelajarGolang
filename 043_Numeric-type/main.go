package main

import (
	"fmt"
	"runtime"
)

var a int8 = 127  //int 8 maksudnya bisa menampung sampe 2^8
var b int16 = 255  //int 16 maksudnya bisa menampung sampe 2^16, dst
var c int32 = 511
var d int64 = 1023
var e float32 = 12.345
var f float64 = 12.3456

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}