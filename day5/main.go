/*
package main

import (

	"fmt"

)

	func main() {
		for i := 0; i < 11; i++ {
			if i == 6 || i == 3 {
				continue
			}
			fmt.Println(i)
		}
	}
*/
/*
package main

import (
	"fmt"
)

func testcount(x int) (y int) {
	if x == 11 {
		fmt.Print(x)
		return
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {
	testcount(1)
}
*/
/*ackage main

import (
	"fmt"
)

func main() {
	var a = 1
	fmt.Printf("%T", a)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	var c bool
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \t %T \n", d, d)
}
*/

package main

import (
	"fmt"
)

/*
	func main() {
		var a, b = 1, "g"
		c, d := 5, "i"
		var g string
		g = "N"
		fmt.Println(a, b)
		fmt.Println(c, d)
		fmt.Println(g)
	}
*/
/*
func myFunc(i int, j string) (result int, txt string) {
	result = i + 1
	txt = "Hi" + " " + j
	return

}
func main() {
	a, _ := myFunc(5, "Guys")
	fmt.Println(a)
}

func meLain() {
	fmt.Println("Hey")
}
*/
/*
func main() {
	//variable naming
	//pascal case
	var AbhiLo = "rt"
	//camel case
	var myFunction = "tree"
	//snake case
	var my_function_done = "Done"

	fmt.Println(AbhiLo)
	fmt.Println(myFunction)
	fmt.Println(my_function_done)
}
*/
const APPLE = "Apple"

func main() {
	const (
		A = 20
		B = 203
	)
	var (
		a = 12
		b = 32
		c = "ADD"
	)
	fmt.Println(a, b, c)
}
