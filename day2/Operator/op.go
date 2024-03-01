package main

import (
	"fmt"
)

func main() {
	/*
		var a = 12
		var b = 12 + 7
		c := a + b
		d := c + 123
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	*/
	/*
		var a = 7 - 11
		var b = 125 / 2
		var c = 7 + 11
		var d = 7 * 11
		var e = 7 % 11
		f := a
		g := b
		a++
		b--
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println(g)
		fmt.Println(a + c)
		fmt.Println(a / 9)
	*/
	/*
		var a = 7
		var b = a
		var c = a
		var d = a
		var e = a
		f := a
		g := a
		b += 3
		c -= 8
		d *= 2
		e /= 12
		f %= 2
		g >>= 3
		k := f << 3
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(e)
		fmt.Println(f)
		fmt.Println(g)
		fmt.Println(k)
	*/
	/*
		var x = 5
		y := x < 4 && x < 10
		fmt.Println(x < 4 && x < 10)
		fmt.Println(x < 6 && x < 10)
		fmt.Println("y is", y)

		fmt.Println(x < 5 && x < 10)
		fmt.Println(!(x < 5 && x < 10))
		fmt.Println(!(x < 5 && x < 10))
		fmt.Println(x < 5 || x < 10)

		fmt.Println(x < 5 || x < 4)
	*/
	/*
		var x = 9

		fmt.Printf("x = %b\n", x)

		fmt.Printf("x << 2 is %b\n", x<<2)

		var y = 8

		fmt.Printf("x = %b\n", x)
		fmt.Printf("y = %b\n", y)

		fmt.Printf("x & y is %b\n", x&y)
		fmt.Printf("x & y is %b\n", x|y)
		fmt.Printf("x & y is %b\n", x^y)
		fmt.Printf("x & y is %b\n", x<<y)
		fmt.Printf("x & y is %b\n", x>>y)
	*/

	x := 16
	y := 9
	fmt.Println(x > y)
	fmt.Println(x != y)
	z := 45
	s := false
	fmt.Println((x > y) && (y > z))
	fmt.Println((x == y) || s)

}
