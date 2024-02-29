package main

import (
	"fmt"
)

func main() {
	arr1 := [4]string{"kali", "billi", "choti", "hai"}
	myslice := arr1[1:3]

	fmt.Println(arr1)
	fmt.Println(myslice)
	fmt.Printf("myslice = %v\n ", myslice)
	fmt.Printf("lenght = %d \n", len(myslice))
	fmt.Printf("Capcity = %d \n", cap(myslice))

}
