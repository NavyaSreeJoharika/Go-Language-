/*
package main

import (

	"fmt"

)

	func main() {
		var a = 1
		var b = 29
		c := a + b

		fmt.Println(c)

		for x := 0; x <= 5; x++ {
			fmt.Println(x)

		}
	}
*/
/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println(divide(10, 0))
}

func divide(l, r int) int {
	return l / r
}
*/
/*
package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(divide(10, 0))

	result, err := divide1(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func divide(l, r int) int {
	return l / r
}
func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor:must not be zero ")
	}
	return l / r, nil
}
*/
/*
package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println(divide(10, 0))

	// result, err := divide1(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("result:", result)
	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func divide(l, r int) int {
	return l / r
}
// if divide function is  directly used we will get error in the code
func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor:must not be zero ")
	}
	return l / r, nil
}
func divide2(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
*/
/*
package main

import (
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		//return 0, errors.New("cannot divide by zero")
		return 0, fmt.Errorf("cannot divide by zero %f", b)
	}
	return a / b, nil
}

func main() {
	//result, err := Divide(10, 2)
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
}
*/
package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d:%s", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		//return 0, errors.New("cannot divide by zero")
		return 0, CustomError{Code: 1, Message: "cannot divide by zero"}
	}
	return a / b, nil
}

func main() {
	//result, err := Divide(10, 2)
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
}
