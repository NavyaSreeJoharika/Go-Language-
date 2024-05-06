package main

import "fmt"

func main() {
	var a = "string "

	for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}
}
