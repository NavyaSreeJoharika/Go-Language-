/*
package main

import (

	"fmt"

)

	func main() {
		a := "Hi"
		var hee = &a
		gee := &a

		fmt.Println(gee)
		fmt.Println(*hee)
	}
*/
/*
package main

import (
	"fmt"
)

type orange struct {
	one   int
	two   int
	three int
}

func (a orange) apple() {
	fmt.Println("one:", a.one)
	fmt.Println("two:", a.two)
	fmt.Println("three:", a.three)
}

func main() {
	res := orange{
		one:   1,
		two:   2,
		three: 3,
	}
	res.apple()

}
*/
/*
package main

import "fmt"

type data int

func (d1 data) red(d2 data) data {
	return d1 + d2
}

func main() {
	value1 := data(2)
	value2 := data(3)

	res := value1.red(value2)

	fmt.Println("final redult:", res)

}
*/
/*
package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
}

// Method with a receiver of author type
func (a *author) show(abranch string) {
	(*a).branch = abranch
}

// Main function
func main() {

	// Initializing the values of the author structure
	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(Before): ", res.branch)

	// Creating a pointer
	p := &res

	// Calling the show method
	p.show("ECE")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(After): ", res.branch)
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type pro struct {
		Name        string
		Age         int
		Description string
		secret      string
	}
	x := pro{
		Name:   "Boo",
		Age:    12,
		secret: "Shhhhh",
	}
	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := map[string]string{
		"fu": "bai",
	}
	data, _ := json.Marshal(x)
	fmt.Println(string(data))
}
*/
/*
package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {
	human1 := Human{"Akhila", 23, "Delhi"}

	human_enc, err := json.Marshal(human1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "Roy", Age: 32, Address: "Pune"},
		{Name: "Paru", Age: 21, Address: "Munbai"},
		{Name: "Shiva", Age: 20, Address: "Bangalore"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println(string(human2_enc))

}
*/

package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3}

	for k, v := range a {
		fmt.Printf("%v :%v,", k, v)
	}

	var s = make(map[string]string)

	s["Name"] = "Jeon"
	s["Year"] = "1997"

	fmt.Printf("%v", s)

	var r = map[int]string{1: "Apple", 2: "orange", 3: "graps"}

	fmt.Printf("\n %v\n", r)
}
