/*
package main

import (
	"fmt"
)

func main() {
	x := "disk"
	fmt.Println(x)
	fmt.Printf("%17s \n", x)
	fmt.Printf("%16s \n", x)
	fmt.Printf("%15s \n", x)
	fmt.Printf("%14s \n", x)
	fmt.Printf("%13s \n", x)
	fmt.Printf("%12s \n", x)
	fmt.Printf("%11s \n", x)
	fmt.Printf("%10s \n", x)
	fmt.Printf("%9s \n", x)
	fmt.Printf("%8s \n", x)
	fmt.Printf("%7s \n", x)
	fmt.Printf("%6s \n", x)
	fmt.Printf("%q", x)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	day := 6

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4,11:
		fmt.Println("Thursday")
	case 5,10:
		fmt.Println("Friday")
	case 6,9:
		fmt.Println("Saturday")
	case 7,8:
		fmt.Println("Sunday")
	default:
		fmt.Println("don't worry")
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 5, 6, 8, 7, 3, 4, 23, 2, 31, 10, 2, 38, 43, 4, 5}
	fmt.Println(arr)

	Slice := arr[1 : len(arr)-4]
	mySlice := make([]int, len(Slice), 100)
	copy(mySlice, Slice)

	fmt.Println(Slice)
	fmt.Println(mySlice)
	fmt.Println(cap(mySlice))
	fmt.Println(len(mySlice))

}
*/
/*
package main

import (
	"fmt"
)

func main() {

	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < 5; i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(i, fruits[j])
		}
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 1; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j], "looking good")
		}
	}
}
*/
/*
package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func main() {
	myMessage()
}
*/
/*
package main

import (
	"fmt"
)

type good struct {
	name  string
	age   int
	india bool
}

func main() {
	var pass1 good
	var pass2 good

	pass1.name = "Charl"
	pass1.age = 31
	pass1.india = true

	pass2.india = true
	pass2.name = "Kris"
	pass2.age = 19

	fmt.Println("Name -", pass1.name)
	fmt.Println("Age -", pass1.age)
	fmt.Println("Indian -", pass1.india)

	printgood(pass2)
	printgood(pass1)
}
func printgood(pass good) {
	fmt.Println("Name: ", pass.name)
	fmt.Println("Age: ", pass.age)
	fmt.Println("Indian: ", pass.india)

}
*/
/*
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = map[int]int{1: 11, 2: 22, 3: 33, 4: 44}
	b := map[int]string{1: "apple", 2: "mango", 3: "orange"}

	fmt.Printf("a %v\n", a)
	fmt.Printf("b %v\n", b)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	// a is no longer empty
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[int]int)
	a[1] = 31
	a[2] = 32
	a[3] = 33

	b := make(map[string]int)
	b["car"] = 5
	b["bike"] = 10
	b["cycle"] = 30
	fmt.Printf("a %v\n", a)
	fmt.Printf("b %v\n", b)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[int]int)
	var b map[string]int

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf(a["brand"])
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	b := make(map[string]string)
	b["car"] = "5"
	b["bike"] = "10"
	b["cycle"] = "30"

	fmt.Printf(b["bike"])
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["bike"] = "Royal"
	a["model"] = "A-20"
	a["year"] = "1964"

	fmt.Println(a)

	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element

	fmt.Println(a)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["bike"] = "Royal"
	a["model"] = "A-20"
	a["year"] = "1964"

	fmt.Println(a)

	delete(a, "year")

	fmt.Println(a)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Royal", "model": "M-30", "year": "1964", "day": ""}

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1) //Royal true
	fmt.Println(val2, ok2) //""False
	fmt.Println(val3, ok3) //"" True
	fmt.Println(ok4)       //true
}
*/

package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Royal", "model": "M-30", "year": "1970"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1990"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}
