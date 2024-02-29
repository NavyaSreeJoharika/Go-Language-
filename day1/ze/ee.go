package main

import "fmt"

func main() {
	var arr1 = [5]int{9, 8, 7}
	arr2 := [3]string{"Go", "Stop", "Wait"}
	arr3 := []int{55, 66}
	arr1[4] = 77
	arr1[3] = 89
	var arr4 = [9]string{}
	arr5 := [3]int{}
	arr6 := [8]int{1: 65, 6: 73, 5: 99}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr3[0])
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(len(arr3))
}
