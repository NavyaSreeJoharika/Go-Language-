package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	username string
	age      int
	verbose  bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "A brief description of your application",
		Long:  "A longer description that spans multiple lines and likely contains examples and usage of using your application.",
		Run: func(cmd *cobra.Command, args []string) {
			// This will be executed when the command is called
			// if verbose {
			// 	fmt.Println("Verbose mode enabled")
			// }
			fmt.Printf("Hello, %s! You are %d years old.\n", username, age)

		},
	}

	// Define flags
	rootCmd.Flags().StringVarP(&username, "name", "n", "", "Your name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "Your age")
	// rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	// Execute command
	rootCmd.Execute()
}
