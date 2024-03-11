/*
package main

import (

	"fmt"

)

	func main() {
		a := map[int]string{1: "hi", 2: "Hello", 3: "Welcome", 4: "Good Day"}
		for k, v := range a {
			fmt.Printf("%v :%v \n", k, v)

		}
		b := make(map[int]string)
		fmt.Println(b)
		fmt.Println(b == nil)

		a[1] = "hey" //update elements
		a[8] = "God" //add elements

		delete(a, 2)
		fmt.Println(a)
		val, ok := a[4]
		fmt.Println(val, ok)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	var b = 20
	x := 0xFF
	y := 0x9C
	fmt.Printf("Address of a variable: %X\n", b)
	fmt.Printf("Address of a variable: %X\n", &b)

	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)
}
*/
/*
package main

import (
	"fmt"
)

func sss(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return sss(x + 1)
}
func main() {
	fmt.Println(sss(11))
	//sss(1)
}
*/
/*
package main

import (
	"fmt"
)

type gee struct {
	number      int
	name        string
	phoneNumber int
}

func main() {
	var peas1 gee
	var peas2 gee

	peas1.number = 1
	peas1.name = "Navya"
	peas1.phoneNumber = 12948234238

	peas2.number = 2
	peas2.name = "Ram"
	peas2.phoneNumber = 24345837427

	fmt.Println(peas1.number)
	fmt.Println(peas1.name)
	fmt.Println(peas1.phoneNumber)

	fmt.Println(peas2.number)
	fmt.Println(peas2.name)
	fmt.Println(peas2.phoneNumber)

	printgee(peas1)
	printgee(peas2)

}
func printgee(peas gee) {
	fmt.Println("Numb :", peas.number)
	fmt.Println("Name :", peas.name)
	fmt.Println("phone number :", peas.phoneNumber)
}
*/
/*
package main

import (
	"fmt"
)

func frac(x int) int {
	if x == 11 {
		return 0
	} else {
		return frac(x * 11)
	}
}
func main() {
	fmt.Println(frac(1))
}
*/
/*
package main

import (
	"fmt"
)

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

func main() {
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:     "123 Main St",
			City:       "Anytown",
			State:      "CA",
			PostalCode: "12345",
		},
	}

	fmt.Println(p.FirstName, p.LastName)
	fmt.Println("Age:", p.Age)
	fmt.Println("Address:")
	fmt.Println("Street:", p.Address.Street)
	fmt.Println("City:", p.Address.City)
	fmt.Println("State:", p.Address.State)
	fmt.Println("Postal Code:", p.Address.PostalCode)
}
*/
/*
package main

import (
	"fmt"
)

type gojo struct {
	height int
	color  string
}

type nanami struct {
	person  string
	details gojo
}

func main() {
	result := nanami{
		person:  "strong",
		details: gojo{7, "blue"},
	}
	fmt.Println(result.details.height)
	fmt.Println(result.details.color)
	fmt.Println(result.person)
	fmt.Println(result)
}
*/
/*
package main

import "fmt"

type Salaried interface {
	getSalary() int
}

type Salary struct {
	basic     int
	insurence int
	allowance int
}

func (s Salary) getSalary() int {
	return s.basic + s.allowance + s.insurence
}

type Employee struct {
	firstname, lastname string
	Salaried
}

func main() {
	ross := Employee{
		firstname: "Ross",
		lastname:  "Gandi",
		Salaried:  Salary{123, 2356467, 464609345},
	}
	fmt.Println("Ross Details ", ross)
	fmt.Println("Ross Details ", ross.firstname)
	fmt.Println("Ross Details ", ross.lastname)
	fmt.Println("Ross Details ", ross.Salaried.getSalary())

}
*/
package main

import (
	"fmt"
)

func main() {
	var map_1 map[string]int

	fmt.Println(map_1)

	var map_2 = make(map[int]string)

	fmt.Println(map_1 == nil)
	fmt.Println(map_2 == nil)

	var map_3 = map[int]string{
		30: "u",
		32: "w",
		56: "r",
	}
	fmt.Println(map_3)
}
