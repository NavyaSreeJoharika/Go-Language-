/*
package main

import (
	"fmt"
	"os"
)

func main() {
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

*/
/*

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [your name]")
		os.Exit(1)
	}

	name := os.Args[1]
	fmt.Printf("Hello, %s! Welcome to the CLI!\n", name)

	fmt.Println("Program name:", os.Args[0])

}

*/
/*
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Basic Calculator CLI!")
	fmt.Println("Available operations: add, subtract, multiply, divide")

	var operation string
	fmt.Print("Enter operation (add/subtract/multiply/divide): ")
	fmt.Scanln(&operation)

	var num1, num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result := calculate(operation, num1, num2)
	fmt.Println("Result:", result)
}

func calculate(operation string, num1, num2 float64) float64 {
	switch operation {
	case "add":
		return num1 + num2
	case "subtract":
		return num1 - num2
	case "multiply":
		return num1 * num2
	case "divide":
		if num2 != 0 {
			return num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
	default:
		fmt.Println("Error: Invalid operation")
		os.Exit(1)
	}
	return 0
}
*/

/*
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A simple CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the CLI application!")
		},
	}

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greet a user",
		Args:  cobra.ExactArgs(3), // Expects exactly three arguments
		Run: func(cmd *cobra.Command, args []string) {
			firstName := args[0]
			lastName := args[1]
			age := args[2]
			fmt.Printf("Hello, %s %s! You are %s years old.\n", firstName, lastName, age)
		},
	}

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


Output:
PS C:\Go Language\clone golang\Go-Language-\CLI> go run main.go greet Pramodh Gompa 19
Hello, Pramodh Gompa! You are 19 years old.

*/

package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Initialize a new AWS session using your AWS credentials
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("your-aws-region"),
	}))

	// Create an S3 service client
	svc := s3.New(sess)

	// List all buckets
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error listing buckets:", err)
		return
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(bucket.Name), aws.TimeValue(bucket.CreationDate))
	}
}
