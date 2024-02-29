package main

import (
	"fmt"
)

func main() {
	myslice := [5]int{1, 2, 0, 4, 5}
	myslice1 := []int{}
	myslice2 := []string{}
	myslice3 := [9]string{"Apples", "are", "tasty"}

	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice3)
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

}
