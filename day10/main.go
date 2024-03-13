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
	five   = 5
	six    = "six"
	_
	four = iota
)

func main() {
	fmt.Println(PI)
	fmt.Println(first, second)
	fmt.Println(third, four)
	//need not to use all the declared constants
}
*/
/*
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

	var pr *int = new(int)
	fmt.Println(pr)
	fmt.Println(*pr)

	*pr = 10
	fmt.Println(*pr)
	fmt.Println(pr)

	pr = nil
	fmt.Println(pr)

}
*/
/*
package main

import (
	"fmt"
)

const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4

	// admin will have all of the permissions
	admin = read | write | remove
)

func main() {
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)
}
*/
/*
package main

import (
	"fmt"
)

const (
	first  = 1
	second = iota
	third  = iota + 5
	four
)

func main() {
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(four)
}

*/
/*
package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	arr[0] = 0
	arr[1] = 4
	arr[2] = 34

	fmt.Println(arr)

	arr2 := [3]int{33, 32, 1}
	fmt.Println(arr2)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	fmt.Println(arr, slice)
	arr[1] = 42
	slice[2] = 33
	fmt.Println(arr, slice)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(slice)
	slice = append(slice, 23, 4, 23, 5, 46)
	fmt.Println(slice)
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

}
*/
/*
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"fu": 42}
	fmt.Println(m)
	fmt.Println(m["fu"])

	m["fu"] = 24
	fmt.Println(m)
	delete(m, "fu")
	fmt.Println(m)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var s user
	fmt.Println(s)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var s user
	s.ID = 1
	s.FirstName = "Navya"
	s.LastName = "Jo"
	fmt.Println(s)
	fmt.Println(s.FirstName)

	u2 := user{ID: 1, FirstName: "Navya", LastName: "Jo"}
	fmt.Println(u2)
	u3 := user{ID: 1,
		FirstName: "Madhu",
		LastName:  "Jo",
	}
	fmt.Println(u3)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server ...")
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
*/ /*
package main

import (
	"fmt"
)

func main() {

	port := 3000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

func startWebServer(port int) bool {
	fmt.Println("Starting server ...")
	fmt.Println("Server started on port", port)
	return true
}
*/
/*
package main

import (
	"fmt"
)

func main() {

	port := 3000
	port, err := startWebServer(port)
	fmt.Println(port, err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server ...")
	fmt.Println("Server started on port", port)
	//return nil
	//return errors.New("something went wrong")
	return port, nil
}
*/

package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "GO",
		LastName:  "Gooooo",
	}
	fmt.Println(u)
}
