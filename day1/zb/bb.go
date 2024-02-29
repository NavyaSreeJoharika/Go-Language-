package main

import (
	"fmt"
)

func main() {
	var r int = 234387343457632094
	var d int8 = 123
	var s int16 = 2356
	var a int32 = -34534345
	var f int64 = -234552

	var q uint = 234345
	var w uint16 = 234
	var e uint32 = 49877

	fmt.Printf("type =%T , value = %v \n", r, r)
	fmt.Printf("type =%T , value = %v \n", s, s)
	fmt.Printf("type =%T , value = %v \n", d, d)
	fmt.Printf("type =%T , value = %v \n", f, f)
	fmt.Printf("type =%T , value = %v \n", a, a)

	fmt.Printf("type =%T , value = %v \n", q, q)
	fmt.Printf("type =%T , value = %v \n", w, w)
	fmt.Printf("type =%T , value = %v \n", e, e)
}
