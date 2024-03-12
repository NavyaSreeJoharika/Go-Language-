/*
package main

import (

	"fmt"

)

	func main() {
		var data = map[int]string{1: "A", 2: "B"}
		fmt.Println(data)

		var data11 = make(map[int]map[int]string)
		data["Date_2"] = make(map[string]map[string]string)
		data["Date_2"]["Sistem_A"] = make(map[string]string)
		data11[1] = "brand"
		data11[2] = "char"
		data11[3] = "table"

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"k1": 1, "k2": 2}
	fmt.Println(m)
	m["k3"] = 3
	m["k4"] = 4
	fmt.Println("Map :", m)
	var h = make(map[int]string)

	fmt.Println(h == nil)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

		n2 := map[string]int{"foo": 1, "bar": 2}
		if maps.Equal(n, n2) {
			fmt.Println("n == n2")
		}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]map[string]int)
	vegmap := map[string]int{"carrot": 6, "ginger": 2}
	a["veg"] = vegmap
	fruitmap := map[string]int{"apple": 1, "orange": 10}
	a["Fruits"] = fruitmap

	fmt.Println(a)
	for key, Value := range a {
		fmt.Println("Category:", key)
		fmt.Println("Category Details:", a[key])
		for k, value := range Value {
			fmt.Printf("%v\t %v \n", k, value)
		}

	}

}
*/
/*
package main

import "fmt"

type Salary struct {
	basic     int
	insurence int
	allowance int
}

type Employee struct {
	firstname, lastname string
	Salary
}

func main() {
	ross := Employee{
		firstname: "Ross",
		lastname:  "Gandi",
		Salary:    Salary{123, 2356467, 464609345},
	}
	fmt.Println("Ross Details ", ross)
	fmt.Println("Ross Details First Name", ross.firstname)
	fmt.Println("Ross Details Last Name", ross.lastname)
	fmt.Println("Ross Details allowance ", ross.Salary.allowance)
	fmt.Println("Ross Details insurence ", ross.Salary.insurence)
	fmt.Println("Ross Details basic", ross.Salary.basic)

}
*/
/*
package main

import "fmt"

func main() {
	fmt.Print("Hi")
}
*/
/*
package main

import (

	"fmt"

)

	func main() {
		fmt.Print("Hello")
	}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	*firstName = "Auther"

	fmt.Println(firstName) //gives memory location
	fmt.Println(*firstName)

	secondName := "Shuuu"
	fmt.Println(secondName)

	ptr := &secondName
	fmt.Println(ptr, *ptr)

	secondName = "Tricia"
	fmt.Println(ptr, *ptr)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	const PI = 3.145
	fmt.Println(PI)

	const c = 3
	fmt.Println(c + 3)

	fmt.Println(c + 1.2)

	const d int = 4
	fmt.Println(d + 2)
	fmt.Println(float32(d) + 1.2)
}
*/
/*
package main

import (
	"fmt"
)

const PI = 3.141
const (
	first  = 1
	second = "second"
	third  = iota
	four   = iota
)

func main() {
	fmt.Println(PI)
	fmt.Println(first, second)
	fmt.Println(third, four)

}
*/
package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var ptr *int = &x

	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
