package main

import "fmt"

func main() {
	fmt.Println(1, 2, 3, 4, 5)
	fmt.Println("Numbers to 5")

	fmt.Println("Hello! Greetings Everyone ")

	fmt.Println("hi , ni hao , hola , konichiwa , bonjour , Annyeong Haseyo , namaste , sawasdee krub/ka")
	fmt.Println("feature1 branch ")

	fmt.Println("flow 1")
	fmt.Println()

	fmt.Print(1, 2, 3, 4, 5, 6, 7, 8)
	little()
}

func little() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++

	}

	for i := 0; i < 10; i++ {
		fmt.Println("bruce", i)
	}

}
