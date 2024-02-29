package main

import (
	"fmt"
)

func main() {
	var x int = 33
	var y string = "Noooo"
	var z float32 = 34.54
	fmt.Printf("%v\n", x)
	fmt.Printf("%#v \n", x)
	fmt.Printf("%T \n", x)
	fmt.Printf("%v%% \n", x)

	fmt.Printf("%v \n", y)
	fmt.Printf("%v%% \n", y)
	fmt.Printf("%T \n", y)
	fmt.Printf("%#v \n", y)

	fmt.Printf("%v \n", z)
	fmt.Printf("%v%% \n", z)
	fmt.Printf("%T \n", z)
	fmt.Printf("%#v \n", z)
}
