package main

import (
	"fmt"
)

func main() {
	var a string = "Jo"
	var b = "Hi"
	var c string
	d := "hey"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("Type : %T , Value : %v \n", a, a)
	fmt.Printf("Type : %T , Value : %v \n", b, b)
	fmt.Printf("Type : %T , Value : %v \n", c, c)
	fmt.Printf("Typr : %T , Value : %v \n\n", d, d)

	fmt.Printf("Type : %T , Value : %#v \n", a, a)
	fmt.Printf("Type : %T , Value : %#v \n", b, b)
	fmt.Printf("Type : %T , Value : %#v \n", c, c)
	fmt.Printf("Typr : %T , Value : %#v \n", d, d)

}
