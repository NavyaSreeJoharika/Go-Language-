/*
package main

import (

	"io"
	"net/http"
	"os"

)

	func main() {
		http.HandleFunc("/", Handler)
		http.ListenAndServe(":3000", nil)
	}

	func Handler(W http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./menu1.txt")
		io.Copy(W, f)
	}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type menuItem struct {
		name   string
		prices map[string]float64
	}
	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.10, "triple": 2.50}},
	}

	in := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1)Print Menu")
		fmt.Println("2)Add item")
		fmt.Println("3) Exit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {

		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf(" %10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		case "3":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}
*/
/*
package main

import "fmt"

func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v \n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v \n", dividend, divisor, divide(dividend, divisor))

}

func divide(dividend, divisor int) int {

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}
*/
package main

import (
	"bufio"
	"day14/menu"
	"fmt"
	"os"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1)Print Menu")
		fmt.Println("2)Add item")
		fmt.Println("3) Exit")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu()
		case "2":
			menu.AddItem()
		case "3":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}
