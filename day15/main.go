package main

import (
	"bufio"
	"day15/menu"
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
			menu.Print()
		case "2":
			menu.Add()
		case "3":
			break loop
		default:
			fmt.Println("unknown option")
		}
	}
}
