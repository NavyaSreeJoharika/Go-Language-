package main

import (
	"fmt"
)

func main() {
	/*
		var x = 7
			if x > 2 {
				fmt.Println("x is greater than 2")
			}
	*/
	/*
		if 20 > 60 {
			fmt.Println("20 is greater than 60")
		}
	*/
	/*
		if 22 > 2 {
			fmt.Println("22 is greater than 2")
		}
	*/
	/*
		v := 33
		z := 20

		if v > z {
			fmt.Println("v is greater than z and v =", v)
		}
	*/
	/*
		if 20 > 60 {
			fmt.Println("20 is greater than 60")
		} else {
			fmt.Println("20 is smaller than 60")
		}
	*/
	/*
		x := 30
		if x >= 10 {
			fmt.Println("x is larger than or equal to 10.")
		} else if x > 20 {
			fmt.Println("x is larger than 20.")
		} else {
			fmt.Println("x is less than 10.")
		}
	*/
	/*
		x := 9
		if x >= 10 {
			fmt.Println("x is larger than or equal to 10.")
		} else if x > 20 {
			fmt.Println("x is larger than 20.")
		} else {
			fmt.Println("x is less than 10.")
		}
	*/
	/*
		x := 8
		if x >= 10 {
			fmt.Println("x is larger than or equal to 10.")
		} else if x > 5 {
			fmt.Println("x is larger than 5.")
		} else {
			fmt.Println("x is less than 10.")
		}
	*/
	/*
		x := 30
		if x >= 10 {
			fmt.Println("x is larger than or equal to 10.")
			if x > 20 {
				fmt.Println("x is larger than 20.")
			}
		} else {
			fmt.Println("x is less than 10.")
		}*/
	/*
		num := 20
		if num >= 10 {
			fmt.Println("Num is more than 10.")
			if num > 15 {
				fmt.Println("Num is also more than 15.")
			}
		} else {
			fmt.Println("Num is less than 10.")

		}
	*/
	/*
		y := 9
		k := 23
		if k > 20 {
			fmt.Println("k is greater than 20 ")
			if y < k {
				fmt.Println("k is greater than y")
				if k < y {
					fmt.Println("y is greater than k")
					if k == 23 {
						fmt.Println("y and k have equal values")
					}
				}
			}
		}
	*/
	z := 10
	b := 20
	if z < b {
		fmt.Println("b is greater than z")
		if b < 15 {
			fmt.Println("b is less than 15")
		} else if b < z {
			fmt.Println("z is greater than b")
		} else {
			fmt.Println("go run ")
		}
	} else {
		fmt.Println("Nee hao")
	}
}
