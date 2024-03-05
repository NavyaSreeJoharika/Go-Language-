/*
package main

import "fmt"

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello 😃")
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	//var name, age = "Lee", 23
	var name = "chan"
	age := 25
	day := true
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%v\t %v\n", name, age)
	fmt.Printf("%#v \t %#v \n", name, age)
	fmt.Printf("%T \t %T \n", name, age)
	fmt.Printf("%% \n")
	fmt.Printf("%v %% \t %v%%\n", name, age)
	fmt.Printf(" %t \n", day)
	fmt.Printf(" %b \n", age)
	fmt.Printf(" %c \n", age)
	fmt.Printf(" %d \n", age)
	fmt.Printf(" %o \n", age)
	fmt.Printf(" %O \n", age)
	fmt.Printf(" %q \n", age)
	fmt.Printf(" %x \n", age)
	fmt.Printf(" %X \n", age)
	fmt.Printf(" %U \n", age)
	fmt.Printf(" %s \n", name)
	fmt.Printf(" %q \n", name)
	fmt.Printf(" %x \n", name)
	fmt.Printf(" %X \n", name)
	fmt.Printf(" %q \n", age)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println(`	The Cap function returns one of the following,
	depending on the type of v : Array: If v is an array,
	then the Cap function returns the number of elements in v .
	Slice: If v is a slice ,
	the Cap function returns the maximum length v can reach upon being resliced.
	If v is nil , 0 is returned.`)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.747
	fmt.Printf("%v of type %T \n", m, m)

	k := float64(y)
	fmt.Printf("%v of type %T \n", k, k)
}*/

/*

package main

import (
	"fmt"
)

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.747
	fmt.Printf("%v of type %T \n", m, m)

	k := float64(m)
	fmt.Printf("%v of type %T \n", k, k)
}
*/
/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))
	fmt.Println("My Favorite number is ", rand.Intn(2000))

}
*/
/*
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("now you have %g problem .\n", math.Sqrt(100))
	fmt.Println(math.Pi)

}
*/
/*
package main

import (
	"fmt"
)


func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(23, 43))
	sayHello()
}
func sayHello() {
	fmt.Println("say Hello")
}


func add(a int, b int) int {
	return a + b
}
func main() {
	a := add(100, 200)
	fmt.Println(a)
	fmt.Println(add(23, 65))
}
*/
/*
package main

import (
	"fmt"
)

func sap(a string, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println(sap("Morning", "Good"))

	a, b := sap("World", "Hello")
	fmt.Println(a, b)
}
*/
/*
package main

import (
	"fmt"
)

func slice(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := slice(23)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(slice(234))
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

	myslice2 := arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	myslice1 = append(myslice1, 21, 22, 23) // Change length by appending items
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

var c, python, java, true bool

const b = 43

func main() {
	var i int
	fmt.Println(i, c, python, java, true, b)

}
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time now", time.Now())
	arr1 := [4]int{4, 5, 3, 2}
	arr2 := []string{"a", "b", "c", "d", "e", "f"}
	slice := arr1[2:3]
	slice = arr1[1:4]
	slice1 := arr2[1:5]

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
*/
/*
package main

import (
	"fmt"
)


func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[len(numbers)-10:]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	day := 4
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("NoGO")
	}

	var car = "bang"
	for i := 0; i < len(car); i++ {
		fmt.Println(i, string(car[i]))
		}

func add(x int, y int) int {
	return x + y

}

func main() {
	fmt.Println(add(23, 43))
	sayHello()
}
func sayHello() {
	fmt.Println("say Hello")
}
*/

package main

import (
	"fmt"
)

func test_count(x int) (y int) {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	//y = test_count(x + 1)
	return test_count(x + 1)
}
func main() {
	test_count(1)
}

/*
package main

import (
	"fmt"
)

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(factorial_recursion(4))
}
*/
