/*
package main

import (

	"bytes"
	"fmt"
	"strings"

)

	type printer interface {
		Print() string
	}

	type user struct {
		username string
		id       int
	}

	func (u user) Print() string {
		return fmt.Sprintf("%v [%v]\n", u.username, u.id)
	}

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	func (mi menuItem) Print() string {
		var b bytes.Buffer
		b.WriteString(mi.name + "\n")
		b.WriteString(strings.Repeat("-", 10) + "\n")

		for size, cost := range mi.prices {
			fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
		}

		return b.String()
	}

	func main() {
		var p printer
		p = user{username: "adent", id: 42}

		fmt.Println(p.Print())

		p = menuItem{name: "Coffee",
			prices: map[string]float64{"small": 1.20,
				"medium": 1.50,
				"large":  1.94,
			},
		}
		fmt.Println(p.Print())
		u, ok := p.(user)
		fmt.Println(u, ok)
		mi, ok := p.(menuItem)
		fmt.Println(mi, ok)

		switch v := p.(type) {
		case user:
			fmt.Println("Found a user!", v)
		case menuItem:
			fmt.Println("\n Found a menuItem!", v)
		default:
			fmt.Println("i am not sure what this is ...")
		}
	}
*/
/*
package main

import "fmt"

func main() {
	testScores := []float64{
		3.45,
		23,
		545,
		3.45,
	}
	c := clone(testScores)
	fmt.Println(&testScores[0], &c[0], c)

}

func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
*/
/*
package main

import "fmt"

func main() {
	testScores := map[string]float64{
		"Harry":   87.23,
		"Hellana": 293,
		"Ron":     12.42,
		"Neville": 24,
	}
	c := clone(testScores)
	fmt.Println(c)

}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
*/
/*
package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}

	s1 := add(a1)
	fmt.Printf("Sum of %v:%v\n", a1, s1)
}
func add(s []int) int {
	var result int
	for _, v := range s {
		result += v
	}
	return result
}
*/
/*
package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{1.12, 3.12}

	s1 := add(a1)
	s2 := add(a2)

	fmt.Printf("Sum of %v:%v\n", a1, s1)
	fmt.Printf("Sum of %v:%v\n", a2, s2)

}

type addable interface {
	int | float64
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
*/

package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{1.12, 3.12}
	a3 := []string{"Foo", "Bee", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)
	fmt.Printf("Sum of %v:%v\n", a1, s1)
	fmt.Printf("Sum of %v:%v\n", a2, s2)
	fmt.Printf("Sum of %v:%v\n", a3, s3)
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
