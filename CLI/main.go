package main

import (
	"fmt"
	"os"
)

func mian() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage : mycli[command]")
		return
	}
	command := args[1]
	switch command {
	case "hello":
		fmt.Println("Hello")
	case "greet":
		if len(args) < 3 {
			fmt.Println("Usage : mycli greet [name]")
			return
		}
		name := args[2]
		fmt.Printf("Greeting, %s!\n", name)

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

}
