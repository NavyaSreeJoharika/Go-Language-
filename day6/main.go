/*
package main
import (

	"fmt"

)

	func main() {
		var i, j = "Hello", "World"
		fmt.Print(i, "n", j)
	}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var i = -15
	var k = "Good"
	var j = "Hi"
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
	fmt.Print(j, " ", k, "\n")
	fmt.Printf("%v%#v", k, j)
	fmt.Print(i)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	arr1kjasdfI := []int{1, 2, 4, 3, 4, 3}
	fmt.Println(arr1kjasdfI)

	arr_2 := [5]int{2, 5}
	fmt.Println(arr_2)

	srr := [4]string{"a", "b", "c"}
	fmt.Println(srr)

	var arr1 = []int{1, 43, 2, 3, 4, 9}
	fmt.Println(arr1)

	var ppp = [1]string{"1"}

	fmt.Println(ppp)

	fmt.Printf("%v  - % v \n", arr1kjasdfI, srr)

	bb1 := []int{1, 3}
	fmt.Println(bb1)
	bb1[0] = 34
	fmt.Println(bb1)
	cc1 := []int{1: 10, 4: 12}
	fmt.Println(cc1)
	//x := cc1[1]
	var x = cc1[1]
	fmt.Println(x)
	//slice1 := arr1[1:9]
	slice1 := arr1[3:]
	//slice1 := arr1[3:6]
	//slice1 := arr1[3:7]
	//slice1 := arr1[:6]
	fmt.Println(slice1)

	mySlice := make([]int, 5, 24)
	fmt.Println(mySlice)

	//mySlice[20] = 123
	mySlice[2] = 123
	fmt.Println(mySlice)

	my_Slice2 := make([]int, 0, 5)
	fmt.Println(my_Slice2)

	mySlice4 := append(my_Slice2, 24, 634, 23)
	fmt.Println(mySlice4)
	fmt.Println(len(mySlice4))
	fmt.Println(cap(mySlice4))

	mySlice8 := append(my_Slice2, 2, 64, 3, 6, 34, 234, 23, 234, 4, 3)
	fmt.Println(mySlice8)
	fmt.Println(len(mySlice8))
	fmt.Println(cap(mySlice8))

	number := mySlice8(len(mySlice8)-5:len(mySlice8)-10)
	numberc := make([]int,len(number))
	nnnn :=(numberc ,number)

	mySlice6 := append(mySlice, mySlice4...)
	fmt.Println(mySlice6)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))


}
*/
/*
package main

import (
	"fmt"
)

func main() {
	arr1 :=[]int{1,4,6,3,2,7,9,0,74,7}
	fmt.Println(len(arr1))

	number := arr1[ :(len(arr1))-4];
	numberc := make([]int,len(number))
	copy (numberc ,number)

	fmt.Println(numberc)
	//fmt.Println(len(numberc))
	//fmt.Println(cap(numberc))

}
*/
package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++

	}

	for i := 0; i < 5; i++ {
		fmt.Println("bruce", i)
	}
}
