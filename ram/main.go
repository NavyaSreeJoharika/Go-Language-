package main

import "fmt"

func main() {

	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array

	fmt.Printf("arr1 = %v\n", arr1)
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

}

/*output
arr1 = [9 10 11 12 13 14]
myslice1 = [10 11 12 13]
capacity = 5
myslice1 = [10 11]
length = 2
capacity = 5
myslice1 = [10 11 20 21 22 23]
length = 6
capacity = 10
PS C:\clone golang\Go-Language-\ram> go run main.go
arr1 = [9 10 11 12 13 14]
myslice1 = [10 11 12 13]
length = 4
capacity = 5
myslice1 = [10 11]
length = 2
capacity = 5
myslice1 = [10 11 20 21 22 23]
length = 6
capacity = 10
*/
