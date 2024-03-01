package main

import "fmt"

func main() {
	/*var see = [4]int{1, 2, 3, 4}
	gee := [2]string{"K", "R"}

	fmt.Println(len(see))
	fmt.Println(cap(gee))
	fmt.Println(see)
	fmt.Println(gee)
	*/
	/*
		myslice := [5]int{1, 2, 3}
		myslice2 := [2]int{7, 8}

		fmt.Println(myslice)
		fmt.Println(len(myslice))
		fmt.Println(cap(myslice))

		fmt.Println(myslice2)
		fmt.Println(len(myslice2))
		fmt.Println(cap(myslice2))
	*/
	/*
		arr1 := [4]int{44, 55, 6, 77}
		myslice := arr1[2:3]

		fmt.Printf("my slice = %v \n", myslice)
		fmt.Printf("length is = %d\n", (len(myslice)))
		fmt.Printf("capacity = %d\n", (cap(myslice)))
	*/
	/*

		myslice := make([]int, 4, 5)
		fmt.Printf("My slice = %v\n", myslice)
		fmt.Printf("Lenght = %d \n", len(myslice))
		fmt.Printf("capacity = %d \n", cap(myslice))

		myslice = append(myslice, 23, 645, 65, 234)
		fmt.Println(myslice)
		fmt.Printf("Lenght = %d \n", len(myslice))
		fmt.Printf("capacity = %d \n", cap(myslice))
	*/
	/*
		P := []int{1, 3, 2, 4}
		fmt.Println(P)
		fmt.Println(P[2])
		P[1] = 80
		fmt.Println(P)
		fmt.Println(P[1])

	*/

	/*
		slice4 := []string{"1", "3", "6", "8"}
		slice2 := []string{"hi", "hi", "hehe"}

		slice3 := append(slice4, slice2...)
		fmt.Println(slice3)
	*/
	/*
	   	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	     myslice1 := arr1[1:5] // Slice array
	     fmt.Printf("myslice1 = %v\n", myslice1)
	     fmt.Printf("length = %d\n", len(myslice1))
	     fmt.Printf("capacity = %d\n", cap(myslice1))

	     myslice1 = arr1[1:3] // Change length by re-slicing the array
	     fmt.Printf("myslice1 = %v\n", myslice1)
	     fmt.Printf("length = %d\n", len(myslice1))
	     fmt.Printf("capacity = %d\n", cap(myslice1))

	     myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	     fmt.Printf("myslice1 = %v\n", myslice1)
	     fmt.Printf("length = %d\n", len(myslice1))
	     fmt.Printf("capacity = %d\n", cap(myslice1))

	*/

	arr1 := [30]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 12, 2, 12, 43, 65, 34, 7, 56, 23, 4, 56, 53, 43, 4, 2, 3, 2, 3, 45}
	abb := arr1[11:20]
	fmt.Printf("My Slice =%v\n", abb)
	fmt.Printf("Length = %d\n", (len(abb)))
	fmt.Printf("Capacity = %d\n", (cap(abb)))
}
