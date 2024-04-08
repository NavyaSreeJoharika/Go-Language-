/*
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}

//go run . -level DEBUG
// go run main.go
// go run . -level INFO
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "Name to greet")

	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)
}

//go run main.go == Hello, World!
//go run . -name Navya == Hello, Navya!
*/
/*
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "Name to greet")
	age := flag.Int("age", 30, "Age of the person")
	flag.Parse()
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}

//go run main.go  === Hello, World! You are 30 years old.
//go run . -name Navya -age 22 ==== Hello, Navya! You are 22 years old.

*/
/*
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands
	greetCommand := flag.NewFlagSet("greet", flag.ExitOnError)
	greetName := greetCommand.String("name", "World", "Name to greet")

	if len(os.Args) < 2 {
		fmt.Println("Usage: greet [command]")
		return
	}

	switch os.Args[1] {
	case "greet":
		greetCommand.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *greetName)
	default:
		fmt.Println("Unknown command")
	}
}
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "Name to greet")
	flag.Parse()

	if name == "" {
		fmt.Println("Please provide a name using the -name flag.")
		return
	}

	fmt.Printf("Hello, %s!\n", name)
}

//go run . -name Pramodh === Hello, Pramodh!
