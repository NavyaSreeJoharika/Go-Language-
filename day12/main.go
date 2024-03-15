/*
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Navya",
		Age:  13,
	}
	fmt.Printf("p =%+v\n", p)
}
*/
/*
//method
package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method with a receiver
// of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}
	res1 := author{
		name:      "Lee",
		branch:    "ECE",
		particles: 473,
		salary:    2334787,
	}

	// Calling the method
	res.show()
	res1.show()
}
*/
/*
package main

import "fmt"

func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, GeeksforGeeks")
}

func main() {

	mul(23, 45)
	defer mul(23, 56)
	show()
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
*/
/*
package main

import "fmt"

type author struct {
	name   string
	branch string
}

func (a *author) show_1(abranch string) {
	(*a).branch = abranch
}

func (a author) show_2() {

	a.name = "Gourav"
	fmt.Println("Author's name(Before) : ", a.name)
}

func main() {

	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before): ", res.branch)
	res.show_1("ECE")
	fmt.Println("Branch Name(After): ", res.branch)

	(&res).show_2()
	fmt.Println("Author's name(After): ", res.name)
}
*/
/*
package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
}

func (a *author) show(abranch string) {
	(*a).branch = abranch
}

func main() {

	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name:", res.name)
	fmt.Println("Branch Name(Before):", res.branch)

	p := &res

	p.show("ECE")
	fmt.Println("Author's name:", res.name)
	fmt.Println("Branch Name(After):", res.branch)
}
*/
/*
package main

import (
	"fmt"
)

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {
	value1 := data(23)
	value2 := data(20)

	res := value1.multiply(value2)
	fmt.Println("Final result: ", res)
}
*/
/*
package main

import (
	"fmt"
)

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
	fmt.Println("Idol's Name:", a.name)
	fmt.Println("Branch Name:", a.branch)
	fmt.Println("Debue year:", a.particles)
	fmt.Println("Salary:", a.salary)
}

func main() {
	res := author{
		name:      "kim",
		branch:    "Pop",
		particles: 2015,
		salary:    10000000000,
	}
	res.show()
}
*/
/*
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) details(i int) int {
	p.name = "new name"
	fmt.Println(p.name, p.age)
	return 0
}

func main() {
	fmt.Println("Methods")
	p := Person{
		name: "James",
		age:  25,
	}
	p.details(20)
	fmt.Printf("p = %v\n", p)
}
*/
/*
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) details() {
	p.name = "new name"
	fmt.Println(p.name, p.age)

}

func main() {
	fmt.Println("Methods")
	p := Person{
		name: "James",
		age:  25,
	}
	p.details()
	fmt.Printf("p = %v\n", p)
}
*/
/*
package main

import (
	"fmt"
)

type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}
func main() {
	var i MyInt = 5
	fmt.Println(i.IsPositive())
}
*/

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) IsAdult() bool {
	return p.Age >= 18
}
func main() {
	person := &Person{Name: "Jack", Age: 25}
	fmt.Println(person.IsAdult())
}
