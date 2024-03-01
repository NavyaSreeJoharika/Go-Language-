package main

import (
	"fmt"
)

func main() {
	/*adm := 43
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
	fmt.Printf("43 in binary = %b \n", adm)
	fmt.Printf("42 in hexadecimal = %x \n", adm)
	fmt.Printf("a = %v, b = %v, c = %v, d = %v, e = %v, f = %v \n", a, b, c, d, e, f)
	fmt.Printf("a in binary =%b \nb in binary =%b \n", a, b)
	*/
	a := "Hi        "
	b := "Well"
	var c int
	fmt.Println(a, b)
	fmt.Print(a, b, "\n")
	fmt.Printf("%#v\n%v\n%T\n%T\n", a, b, a, b)
	fmt.Printf("%T", c)

}
