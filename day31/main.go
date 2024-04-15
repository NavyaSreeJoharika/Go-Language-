/*
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}
*/

/*
Output ====
PS C:\golang\day31> go run main.go
Welcome to the playground!
The time is 2024-04-15 09:47:55.1988774 +0530 IST m=+0.002156601
*/

/*
package main

import "fmt"

func main() {

	var w string
	w = "HI"

	fmt.Println(w)

	var b int = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.234
	fmt.Println(d)
	var e int = int(d)

	fmt.Println(e)
}
*/

/*
output ===
PS C:\golang\day31> go run main.go
HI
99
true
3.234
3
*/
/*
package main

import "fmt"

func main() {
	a, b := 5, 2

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(float32(a) / float32(b))
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a > b)

}

*/
/* output ===
PS C:\golang\day31> go run main.go
7
3
10
2
1
2.5
false
false
true
*/
/*

package main

import "fmt"

func main() {
	const a = 42
	var i int = 42
	const b float32 = 32
	var f32 float32 = b
	var f64 float64 = float64(b)

	fmt.Println(f32, f64)
	fmt.Println(i)

	const c = iota
	fmt.Println(c)

	const (
		d = 2 * 5
		e // = 2 * 5
		f = iota
		g
		h = 10 * iota
	)
	fmt.Println(d, e, f, g, h)
}
*/
/* output ===
32 32
42
0
10 10 2 3 40
*/

package main

import "fmt"

func main() {
	s := "Hello, World"

	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hello, Everyone"

	fmt.Println(s)
	p = new(string)

	fmt.Println(p, *p)
}

/* output ===
0xc00004e250
Hello, World
Hello, Everyone
0xc00004e280
*/
