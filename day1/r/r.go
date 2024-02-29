package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hi", "Welcome"
	fmt.Print(i, "\t")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n")
	fmt.Print(j)

	fmt.Print("\n", i, j, " ")

	fmt.Print(" ", i, " ", j, "\n")

	fmt.Println(i, j)
}
